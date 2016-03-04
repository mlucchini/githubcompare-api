package model

import (
	"testing"
	"reflect"
	"time"
)

func TestString(t *testing.T) {
	now := time.Now()
	event := RepositoryStarEvent{ RepositoryName: "MyRepo", Date: now, Stars: 42 }
	expected := "MyRepo," + now.Format(YearMonthDayFormat) + ",42"

	result := event.String()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}

func TestParse(t *testing.T) {
	now, _ := time.Parse(YearMonthDayFormat, time.Now().Format(YearMonthDayFormat))
	expected := RepositoryStarEvent{ RepositoryName: "MyRepo", Date: now, Stars: 42 }

	result := RepositoryStarEvent{}
	err := result.Parse("MyRepo," + now.Format(YearMonthDayFormat) + ",42")

	if err != nil {
		t.Errorf("Error while parsing; result = %+v", result)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}