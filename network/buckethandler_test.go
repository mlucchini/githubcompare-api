package network

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGivenInexistingFileWhenReadFromStorageThenFailsOnInexistingObject(t *testing.T) {
	ctx, ctxDone, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer ctxDone()

	bh := BucketHandler{ ctx }
	reader, bucketDone, err := bh.Reader("bucket", "file")
	fmt.Printf("%+v, %+v", reader, err)
	defer bucketDone()

	assert.Equal(t, err.Error(), "storage: object doesn't exist")
}