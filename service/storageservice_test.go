package service

import (
	"testing"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/testutil"
	"github.com/stretchr/testify/assert"
	"time"
"google.golang.org/appengine/datastore"
)

func TestGivenEmptyStoreWhenPutOneEntityThenStoreHasOneEntity(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	assert.Empty(t, testutil.GetAllEntities(ctx, repositoryStarEventKind, t))

	key, err := (&StorageService{}).Put(ctx, &model.RepositoryStarEvent{ RepositoryName: "MyRepo", Date: time.Now(), Stars: 42 })
	if err != nil {
		t.Fatal(err)
	}
	testutil.EnsureEntitiesAreCommitted(ctx, []*datastore.Key{key}, t)

	assert.Equal(t, 1, len(testutil.GetAllEntities(ctx, repositoryStarEventKind, t)))
}