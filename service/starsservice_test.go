package service

import (
	"testing"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/testutil"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyStoreWhenFilterOnRepositorySortByDateThenReturnEmptySlice(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()

	result, err := (&StarsService{}).FilterOnRepositorySortByDate(ctx, "MyRepo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, len(result))
}

func TestGivenThreeEntitiesWhenFilterOnRepositorySortByDateThenReturnTwoEntities(t *testing.T) {
	ctx, done := testutil.MockContext(t)
	defer done()
	testutil.LoadStore(ctx, []*model.RepositoryStarEvent{
		&model.RepositoryStarEvent{ "Repo1", time.Now(), 42 },
		&model.RepositoryStarEvent{ "Repo2", time.Now(), 43 },
		&model.RepositoryStarEvent{ "Repo1", time.Now().Add(time.Hour * 24), 44 },
	}, t)

	result, err := (&StarsService{}).FilterOnRepositorySortByDate(ctx, "Repo1")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(result))
}