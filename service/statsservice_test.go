package service

import (
	"testing"
	"github.com/mlucchini/githubcompare-api/model"
	"github.com/mlucchini/githubcompare-api/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyStoreWhenGetRepositoryThenReturnError(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()

	_, err := (&StatsService{ ctx }).GetRepository("repo")

	assert.NotNil(t, err)
}

func TestGivenTwoEntitiesWhenGetRepositoryThenReturnOneEntity(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	testutil.LoadStore(ctx, []*model.RepositoryStats{
		&model.RepositoryStats{ RepositoryName: "repo1", Stars: []int{ 41, 42, 43 } },
		&model.RepositoryStats{ RepositoryName: "repo2", Stars: []int{ 41, 42, 43 } },
	}, t)

	result, err := (&StatsService{ ctx }).GetRepository("repo1")

	assert.Nil(t, err)
	assert.Equal(t, "repo1", result.RepositoryName)
}