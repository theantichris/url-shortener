package mongodb

import (
	"context"
	"time"

	"github.com/theantichris/url-shortener/shortener"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const collectionName = "redirects"

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(url string, timeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, err
}

// NewMongoRepository creates and returns a new MongoDB repository.
func NewMongoRepository(url, mongoDB string, timeout int) (shortener.RedirectRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(timeout) * time.Second,
		database: mongoDB,
	}

	client, err := newMongoClient(url, repo.timeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}

	repo.client = client

	return repo, nil
}

func (r *mongoRepository) Find(code string) (*shortener.Redirect, error) {
	const funcPath = "repository.Redirect.Find"

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	redirect := &shortener.Redirect{}
	collection := r.client.Database(r.database).Collection(collectionName)
	filter := bson.M{"code": code}

	if err := collection.FindOne(ctx, filter).Decode(&redirect); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(shortener.ErrRedirectNotFound, funcPath) // TODO: can probably handle this more eloquently
		}

		return nil, errors.Wrap(err, funcPath)
	}

	return redirect, nil
}

func (r *mongoRepository) Store(redirect *shortener.Redirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	collection := r.client.Database(r.database).Collection(collectionName)
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       redirect.Code,
			"url":        redirect.URL,
			"created_at": redirect.CreatedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}

	return nil
}
