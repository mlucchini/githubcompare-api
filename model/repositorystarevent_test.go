package model

import (
	"testing"
	"reflect"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	now := time.Now()
	event := RepositoryStarEvent{ RepositoryName: "repo", Date: now, Stars: 42 }
	expected := "repo," + now.Format(YearMonthDayFormat) + ",42"

	result := event.String()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}

func TestParse(t *testing.T) {
	now, _ := time.Parse(YearMonthDayFormat, time.Now().Format(YearMonthDayFormat))
	expected := RepositoryStarEvent{ "repo", now, 42 }

	result := RepositoryStarEvent{}
	err := result.Parse("repo," + now.Format(YearMonthDayFormat) + ",42")

	assert.Nil(t, err)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}