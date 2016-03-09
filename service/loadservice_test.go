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

	key, err := (&LoadService{ ctx }).Put(&model.RepositoryStarEvent{ "repo", time.Now(), 42 })
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, []*datastore.Key{key}, t)

	assert.Equal(t, 1, len(testutil.GetAllEntities(ctx, repositoryStarEventKind, t)))
}

func TestGivenEmptyStoreWhenPutMultiTwoEntitiesThenStoreHasTwoEntities(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	assert.Empty(t, testutil.GetAllEntities(ctx, repositoryStarEventKind, t))

	keys, err := (&LoadService{ ctx}).PutMulti([]*model.RepositoryStarEvent{
		&model.RepositoryStarEvent{ "repo1", time.Now(), 42 },
		&model.RepositoryStarEvent{ "repo2", time.Now(), 43 },
	})
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, keys, t)

	assert.Equal(t, 2, len(testutil.GetAllEntities(ctx, repositoryStarEventKind, t)))
}

func TestGivenTaskWithOneElementWhenSendTaskThenNoError(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	input := []string{ "line1" }

	err := (&LoadService{ ctx }).SendTask(input)

	assert.Nil(t, err)
}

func TestGivenPayloadWithTwoElementsWhenReceiveTaskThenStoreHasTwoEntities(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	assert.Empty(t, testutil.GetAllEntities(ctx, repositoryStarEventKind, t))

	keys, err := (&LoadService{ ctx }).ReceiveTask("repo1,2016-01-01,42\nrepo2,2016-01-02,43", 1000)
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, keys, t)

	assert.Equal(t, 2, len(testutil.GetAllEntities(ctx, repositoryStarEventKind, t)))
}