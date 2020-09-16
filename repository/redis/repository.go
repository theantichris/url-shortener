package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/theantichris/url-shortener/shortener"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(url string) (*redis.Client, error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)
	if _, err = client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, err
}

// NewRedisRepository creates and returns an new Redirect repository for Redis.
func NewRedisRepository(url string) (shortener.RedirectRepository, error) {
	repo := &redisRepository{}

	client, err := newRedisClient(url)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}

	repo.client = client

	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*shortener.Redirect, error) {
	key := r.generateKey(code)
	fields, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	if len(fields) == 0 {
		return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}

	createdAt, err := strconv.ParseInt(fields["created_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}

	redirect := &shortener.Redirect{
		Code:      fields["code"],
		URL:       fields["url"],
		CreatedAt: createdAt,
	}

	return redirect, nil
}

func (r *redisRepository) Store(redirect *shortener.Redirect) error {
	key := r.generateKey(redirect.Code)
	fields := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}

	if _, err := r.client.HMSet(key, fields).Result(); err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}

	return nil
}
