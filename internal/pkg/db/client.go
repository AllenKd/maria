package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"maria/internal/pkg/configs"
	"sync"
)

type client struct {
	*mongo.Client
}

type QueryOption struct {
	Limit  int64
	Skip   int64
	SortBy string
	Desc   bool
}

func (q *QueryOption) ToMongoOption() (option options.FindOptions) {
	//opt := options.Find()
	if q.Limit > 0 {
		option.Limit = &q.Limit
	}
	if q.Skip > 0 {
		option.Skip = &q.Skip
	}
	if q.SortBy != "" {
		option.SetSort(bson.M{
			q.SortBy: parseSortOrder(q.Desc),
		})
	} else {
		option.SetSort(bson.M{
			"_id": parseSortOrder(q.Desc),
		})
	}
	return
}

func parseSortOrder(o bool) int {
	if o {
		return 1
	} else {
		return -1
	}
}

var (
	once     sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(configs.New().Mongo.ConnectionString))
		if err != nil {
			panic(err)
		}
		db := c.Database(configs.New().Mongo.Db)
		instance = &client{
			Client:       c,
		}
		if err := instance.Ping(nil, nil); err != nil {
			log.Panic(err.Error())
			panic(err)
		}
		instance.initIndex()
		log.Debug("mongo client initialized")
	})
	return instance
}

func (c client) initIndex() {
	ctx := context.Background()
	for col, idx := range CollectionIndexes {
		log.Debug("create index for collection: ", col)

		db := c.Database(configs.New().Mongo.Db)
		_, err := db.Collection(col).Indexes().CreateMany(ctx, idx)
		if err != nil {
			log.Error("unable to create index for collection: ", col, ". ", err.Error())
			panic(err)
		}
	}
}
