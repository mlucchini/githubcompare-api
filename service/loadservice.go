package service

import (
	"github.com/mlucchini/github-compare-backend/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"strings"
	"google.golang.org/appengine/taskqueue"
)

type LoadService struct {
	Context context.Context
}

const defaultQueue = ""
const separator = "\n"

func (self *LoadService) Put(entity *model.RepositoryStarEvent) (*datastore.Key, error) {
	key := datastore.NewKey(self.Context, repositoryStarEventKind, self.stringId(entity), 0, nil)
	_, err := datastore.Put(self.Context, key, entity)
	return key, err
}

func (self *LoadService) PutMulti(entities []*model.RepositoryStarEvent) ([]*datastore.Key, error) {
	keys := make([]*datastore.Key, 0, len(entities))
	for _, entity := range entities {
		key := datastore.NewKey(self.Context, repositoryStarEventKind, self.stringId(entity), 0, nil)
		keys = append(keys, key)
	}
	_, err := datastore.PutMulti(self.Context, keys, entities)
	return keys, err
}

func (self *LoadService) SendTask(elements []string) (error) {
	payload := strings.Join(elements, separator)
	t := &taskqueue.Task{ Path: "/api/admin/loadtask", Payload: []byte(payload), Header: nil, Method: "POST" }
	if _, err := taskqueue.Add(self.Context, t, defaultQueue); err != nil {
		return err
	}
	log.Infof(self.Context, "Added task with %d elements", len(elements))
	return nil
}

func (self *LoadService) ReceiveTask(payload string, entitiesPerTask int) ([]*datastore.Key, error) {
	elements := strings.Split(string(payload), separator)
	entities := make([]*model.RepositoryStarEvent, 0, entitiesPerTask)

	for _, event := range elements {
		var entity model.RepositoryStarEvent
		if err := entity.Parse(event); err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	log.Infof(self.Context, "Received task with %d elements", len(elements))

	keys, err := self.PutMulti(entities)
	if err == nil {
		log.Infof(self.Context, "Stored %d elements", len(elements))
	}

	return keys, err
}

func (self *LoadService) stringId(entity *model.RepositoryStarEvent) string {
	return entity.RepositoryName + "," + entity.Date.Format(model.YearMonthDayFormat)
}