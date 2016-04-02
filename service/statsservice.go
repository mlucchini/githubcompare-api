package service

import (
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"github.com/mlucchini/githubcompare/model"
)

type StatsService struct {
	Context context.Context
}

func (self *StatsService) GetRepository(repositoryName string) (*model.RepositoryStats, error) {
	var entity model.RepositoryStats

	key := datastore.NewKey(self.Context, RepositoryStatsKind, repositoryName, 0, nil)
	err := datastore.Get(self.Context, key, &entity)

	if err != nil {
		log.Errorf(self.Context, "Failed to get element %s: %s", repositoryName, err.Error())
	}

	return &entity, err
}