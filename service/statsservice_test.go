package service

import (
	"testing"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/testutil"
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
		&model.RepositoryStats{ "repo1", []int{ 41, 42, 43 } },
		&model.RepositoryStats{ "repo2", []int{ 41, 42, 43 } },
	}, t)

	result, err := (&StatsService{ ctx }).GetRepository("repo1")

	assert.Nil(t, err)
	assert.Equal(t, "repo1", result.RepositoryName)
}