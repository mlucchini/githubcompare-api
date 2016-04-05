package service

import (
	"testing"
	"github.com/mlucchini/githubcompare-api/model"
	"github.com/mlucchini/githubcompare-api/testutil"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/datastore"
)

func TestGivenEmptyStoreWhenPutOneEntityThenStoreHasOneEntity(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	assert.Empty(t, testutil.GetAllEntities(ctx, RepositoryStatsKind, t))

	key, err := (&LoadService{ ctx }).Put(&model.RepositoryStats{ RepositoryName: "repo", Stars: []int{ 41, 42, 43 } })
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, []*datastore.Key{key}, t)

	assert.Equal(t, 1, len(testutil.GetAllEntities(ctx, RepositoryStatsKind, t)))
}

func TestGivenEmptyStoreWhenPutMultiTwoEntitiesThenStoreHasTwoEntities(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	assert.Empty(t, testutil.GetAllEntities(ctx, RepositoryStatsKind, t))

	keys, err := (&LoadService{ ctx }).PutMulti([]*model.RepositoryStats{
		&model.RepositoryStats{ RepositoryName: "repo1", Stars: []int{ 41, 42, 43 } },
		&model.RepositoryStats{ RepositoryName: "repo2", Stars: []int{ 41, 42, 43 } },
	})
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, keys, t)

	assert.Equal(t, 2, len(testutil.GetAllEntities(ctx, RepositoryStatsKind, t)))
}

func TestWhenPutMulti501EntitiesThenReturnErrorAsDatastoreProductionDoesntSupportIt(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	entities := make([]*model.RepositoryStats, 0)
	for i := 0; i < 501; i++ {
		entities = append(entities, &model.RepositoryStats{ RepositoryName: "repo", Stars: []int{ 41, 42, 43 } })
	}

	_, err := (&LoadService{ ctx }).PutMulti(entities)

	assert.NotNil(t, err)
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
	assert.Empty(t, testutil.GetAllEntities(ctx, RepositoryStatsKind, t))

	keys, err := (&LoadService{ ctx }).ReceiveTask("repo1,41;42\nrepo2,43", 1000)
	assert.Nil(t, err)
	testutil.EnsureEntitiesAreCommitted(ctx, keys, t)

	assert.Equal(t, 2, len(testutil.GetAllEntities(ctx, RepositoryStatsKind, t)))
}