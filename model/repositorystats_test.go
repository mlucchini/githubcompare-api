package model

import (
	"testing"
	"reflect"
	"bytes"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	stat := RepositoryStats{ RepositoryName: "repo", Stars: []int{ 1, 2, 3 } }
	expected := "repo,[1 2 3]"

	result := stat.String()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}

func TestJson(t *testing.T) {
	stat := RepositoryStats{ RepositoryName: "repo", Stars: []int{ 1, 2, 3 } }
	expected := "{\"repository\":\"repo\",\"stars\":[1,2,3]}"

	var b bytes.Buffer
	stat.Json(&b)
	result := strings.Trim(b.String(), "\n")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}

func TestParse(t *testing.T) {
	expected := RepositoryStats{ "repo", []int{ 41, 42, 43 } }

	result := RepositoryStats{}
	err := result.Parse("repo,41;42;43")

	assert.Nil(t, err)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}