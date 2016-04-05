package testutil

import (
	"google.golang.org/appengine/aetest"
	"testing"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"github.com/mlucchini/githubcompare-api/model"
)

func MockContext(t *testing.T) (context.Context, func()) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	return ctx, done
}

func GetAllEntities(ctx context.Context, kind string, t *testing.T) ([]model.RepositoryStats) {
	query := datastore.NewQuery(kind)
	entities := make([]model.RepositoryStats, 0)
	_, err := query.GetAll(ctx, &entities)
	if err != nil {
		t.Fatal(err)
	}
	return entities
}

func EnsureEntitiesAreCommitted(ctx context.Context, keys []*datastore.Key, t *testing.T) {
	for _, key := range keys {
		err := datastore.Get(ctx, key, &model.RepositoryStats{})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func LoadStore(ctx context.Context, entities []*model.RepositoryStats, t *testing.T) {
	keys := make([]*datastore.Key, len(entities))
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		keys[i] = datastore.NewKey(ctx, "RepositoryStats", entity.RepositoryName, 0, nil)
	}
	_, err := datastore.PutMulti(ctx, keys, entities)
	if err != nil {
		t.Fatal(err)
	}
	EnsureEntitiesAreCommitted(ctx, keys, t)
}
