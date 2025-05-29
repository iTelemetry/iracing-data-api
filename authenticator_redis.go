package irdata

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type RedisAuthenticator struct {
	Authenticator

	redis         redis.Cmdable
	authenticator Authenticator

	Logger             Logger
	KeyPrefix          string
	LocalCacheDuration time.Duration

	localCache map[string]*AuthenticationResult
}

func NewRedisDefaultAuthenticator(redis redis.Cmdable, irdata IRData) (*RedisAuthenticator, error) {
	authenticator, err := NewDefaultAuthenticator(irdata)
	if err != nil {
		return nil, err
	}

	return NewRedisAuthenticator(redis, authenticator)
}

func NewRedisAuthenticator(redis redis.Cmdable, authenticator Authenticator) (*RedisAuthenticator, error) {
	if redis == nil {
		return nil, errors.New("redis is nil")
	} else if authenticator == nil {
		return nil, errors.New("authenticator is nil")
	}

	return &RedisAuthenticator{
		redis:              redis,
		authenticator:      authenticator,
		Logger:             slog.Default(),
		KeyPrefix:          "iracing:irdata:",
		LocalCacheDuration: time.Minute,
		localCache:         make(map[string]*AuthenticationResult),
	}, nil
}

type AuthenticationResult struct {
	Username         string
	PasswordChecksum string
	URL              *url.URL
	Cookies          []*http.Cookie
	Expiration       time.Time

	lastUpdated time.Time
}

func init() {
	gob.Register(AuthenticationResult{})
}

func (a *RedisAuthenticator) Authenticate(username string, password string, _ bool) (*url.URL, []*http.Cookie, time.Time, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%032x", hasher.Sum(nil))

	key := fmt.Sprintf("%s%s-%s", a.KeyPrefix, username, hashedPassword)
	existing, ok := a.localCache[key]
	if ok && existing.lastUpdated.Before(time.Now().Add(-a.LocalCacheDuration)) {
		return existing.URL, existing.Cookies, existing.Expiration, nil
	} else if ok {
		delete(a.localCache, key)
	}

	existingResp := a.redis.Get(ctx, key)
	if existingResp == nil || existingResp.Err() != nil {
		if existingResp == nil || !errors.Is(existingResp.Err(), redis.Nil) {
			var err error
			if existingResp == nil {
				err = errors.New("no response from redis")
			} else {
				err = fmt.Errorf("unexpected response from redis: %s", existingResp.Err())
			}

			a.Logger.Warn("unable to retrieve cached result from redis", "error", err)
		}

		return a.authenticate(username, password, hashedPassword, key)
	}

	result := &AuthenticationResult{}
	decoder := gob.NewDecoder(bytes.NewBuffer([]byte(existingResp.Val())))
	err := decoder.Decode(result)
	if err != nil {
		a.Logger.Warn("unable to decode cached result from redis", "error", err)
		return a.authenticate(username, password, hashedPassword, key)
	}

	result.lastUpdated = time.Now()
	a.localCache[key] = result
	return result.URL, result.Cookies, result.Expiration, nil
}

func (a *RedisAuthenticator) authenticate(username, password, hashedPassword, key string) (*url.URL, []*http.Cookie, time.Time, error) {
	cookieUrl, cookies, expiration, err := a.authenticator.Authenticate(username, password, true)
	if err != nil {
		return nil, nil, time.Time{}, err
	}

	result := &AuthenticationResult{
		Username:         username,
		PasswordChecksum: hashedPassword,
		URL:              cookieUrl,
		Cookies:          cookies,
		Expiration:       expiration,
		lastUpdated:      time.Now(),
	}

	a.localCache[key] = result

	buf := new(bytes.Buffer)
	err = gob.NewEncoder(buf).Encode(result)
	if err != nil {
		a.Logger.Warn("unable to encode cached result for redis", "error", err)
		return cookieUrl, cookies, expiration, nil
	}

	setResp := a.redis.Set(context.Background(), key, buf.Bytes(), time.Until(result.Expiration))
	if setResp == nil || setResp.Err() != nil {
		a.Logger.Warn("unable to set cached result in redis", "error", err)
	}

	return cookieUrl, cookies, expiration, nil
}
