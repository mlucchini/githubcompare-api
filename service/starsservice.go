package service

import (
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"github.com/mlucchini/github-compare-backend/model"
)

type StarsService struct {
	Context context.Context
}

func (self *StarsService) FilterOnRepositorySortByDate(repositoryName string) ([]*model.RepositoryStarEvent, error) {
	query := datastore.NewQuery(repositoryStarEventKind).Filter("RepositoryName =", repositoryName).Order("Date")

	events := make([]*model.RepositoryStarEvent, 0)
	_, err := query.GetAll(self.Context, &events)

	return events, err
}