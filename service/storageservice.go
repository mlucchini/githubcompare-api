package service

import (
	"github.com/mlucchini/github-compare-backend/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type StorageService struct {}

func (self *StorageService) Put(ctx context.Context, entity *model.RepositoryStarEvent) (*datastore.Key, error) {
	key := datastore.NewKey(ctx, repositoryStarEventKind, self.stringId(entity), 0, nil)
	_, err := datastore.Put(ctx, key, entity)
	return key, err
}

func (self *StorageService) PutMulti(ctx context.Context, entities []*model.RepositoryStarEvent) ([]*datastore.Key, error) {
	keys := make([]*datastore.Key, len(entities))
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		keys[i] = datastore.NewKey(ctx, repositoryStarEventKind, self.stringId(entity), 0, nil)
	}
	_, err := datastore.PutMulti(ctx, keys, entities)
	return keys, err
}

func (self *StorageService) stringId(entity *model.RepositoryStarEvent) string {
	return entity.RepositoryName + "," + entity.Date.Format(model.YearMonthDayFormat)
}