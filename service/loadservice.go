package service

import (
	"github.com/mlucchini/github-compare-backend/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type LoadService struct {
	Context context.Context
}

func (self *LoadService) Put(entity *model.RepositoryStarEvent) (*datastore.Key, error) {
	key := datastore.NewKey(self.Context, repositoryStarEventKind, self.stringId(entity), 0, nil)
	_, err := datastore.Put(self.Context, key, entity)
	return key, err
}

func (self *LoadService) PutMulti(entities []*model.RepositoryStarEvent) ([]*datastore.Key, error) {
	keys := make([]*datastore.Key, len(entities))
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		keys[i] = datastore.NewKey(self.Context, repositoryStarEventKind, self.stringId(entity), 0, nil)
	}
	_, err := datastore.PutMulti(self.Context, keys, entities)
	return keys, err
}

func (self *LoadService) stringId(entity *model.RepositoryStarEvent) string {
	return entity.RepositoryName + "," + entity.Date.Format(model.YearMonthDayFormat)
}