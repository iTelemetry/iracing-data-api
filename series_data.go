package irdata

type SeriesAssets map[string]SeriesAsset

type SeriesAsset struct {
	LargeImage string `json:"large_image"`
	Logo       string `json:"logo"`
	SeriesCopy string `json:"series_copy"`
	SeriesID   int    `json:"series_id"`
	SmallImage string `json:"small_image"`
}

type Series []struct {
	AllowedLicenses []AllowedLicense `json:"allowed_licenses"`
	Category        string           `json:"category"`
	CategoryID      int              `json:"category_id"`
	Eligible        bool             `json:"eligible"`
	ForumURL        string           `json:"forum_url,omitempty"`
	MaxStarters     int              `json:"max_starters"`
	MinStarters     int              `json:"min_starters"`
	OvalCautionType int              `json:"oval_caution_type"`
	RoadCautionType int              `json:"road_caution_type"`
	SeriesID        int              `json:"series_id"`
	SeriesName      string           `json:"series_name"`
	SeriesShortName string           `json:"series_short_name"`
	SearchFilters   string           `json:"search_filters,omitempty"`
}

type AllowedLicense struct {
	LicenseGroup    int    `json:"license_group"`
	MinLicenseLevel int    `json:"min_license_level"`
	MaxLicenseLevel int    `json:"max_license_level"`
	GroupName       string `json:"group_name"`
}
