package lib

import (
	"testing"
	"strings"
	"reflect"
)

func TestGivenReaderWithLine1WhenGroupLinesIteratorWithGroupOfSize3ThenReturnOneGroupWith1(t *testing.T) {
	reader := strings.NewReader("line1")
	nbLinesPerGroup := 3
	expected := [][]string{ []string{ "line1" } }

	result := make([][]string, 0)
	for group := range GroupLinesIterator(reader, nbLinesPerGroup) {
		result = append(result, group)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}

func TestGivenReaderWithLines123WhenGroupLinesIteratorWithGroupOfSize2ThenReturnTwoGroupsWith12And3(t *testing.T) {
	reader := strings.NewReader("line1\nline2\nline3")
	nbLinesPerGroup := 2
	expected := [][]string{ []string{ "line1", "line2" }, []string{ "line3" } }

	result := make([][]string, 0)
	for group := range GroupLinesIterator(reader, nbLinesPerGroup) {
		result = append(result, group)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result = %+v; expected = %+v", result, expected)
	}
}