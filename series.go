package irdata

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

func (d *irdata) Series() DataSeries {
	return &irdataSeries{parent: d}
}

type irdataSeries struct {
	parent *irdata
}

type DataSeries interface {
	Assets(ctx context.Context, opts ...SeriesAssetsOption) (SeriesAssets, error)
	Get(ctx context.Context) (Series, error)
	Seasons(ctx context.Context, opts ...SeriesSeasonsOption) (SeriesSeasons, error)
}

type seriesAssetsOptions struct {
	values       *url.Values
	imageBaseUrl string
}

type SeriesAssetsOption func(*seriesAssetsOptions)

func WithImageBaseUrl(url string) SeriesAssetsOption {
	if url != "" && !strings.HasSuffix(url, "/") {
		url = url + "/"
	}

	return func(v *seriesAssetsOptions) {
		v.imageBaseUrl = url
	}
}

func (c *irdataSeries) Assets(ctx context.Context, opts ...SeriesAssetsOption) (SeriesAssets, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/series/assets", d.membersUrl))
	if err != nil {
		return SeriesAssets{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	o := &seriesAssetsOptions{
		values:       &q,
		imageBaseUrl: "https://images-static.iracing.com/img/logos/series/",
	}

	for _, opt := range opts {
		opt(o)
	}

	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output SeriesAssets
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return SeriesAssets{}, err
	}

	for key, series := range output {
		if series.LargeImage != "" && o.imageBaseUrl != "" {
			series.LargeImage = o.imageBaseUrl + series.LargeImage
		}

		if series.SmallImage != "" && o.imageBaseUrl != "" {
			series.SmallImage = o.imageBaseUrl + series.SmallImage
		}

		if series.Logo != "" && o.imageBaseUrl != "" {
			series.Logo = o.imageBaseUrl + series.Logo
		}

		output[key] = series
	}

	return output, nil
}

func (c *irdataSeries) Get(ctx context.Context) (Series, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/series/get", d.membersUrl))
	if err != nil {
		return Series{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output Series
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return Series{}, err
	}

	return output, nil
}

type seriesSeasonsOptions struct {
	values *url.Values
}

type SeriesSeasonsOption func(*seriesSeasonsOptions)

func WithIncludeSeries(include bool) SeriesSeasonsOption {
	return func(v *seriesSeasonsOptions) {
		v.values.Set("include_series", fmt.Sprintf("%t", include))
	}
}

func (c *irdataSeries) Seasons(ctx context.Context, opts ...SeriesSeasonsOption) (SeriesSeasons, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/series/seasons", d.membersUrl))
	if err != nil {
		return SeriesSeasons{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	o := &seriesSeasonsOptions{
		values: &q,
	}

	for _, opt := range opts {
		opt(o)
	}

	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output SeriesSeasons
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return SeriesSeasons{}, err
	}

	return output, nil
}
