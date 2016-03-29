package network

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	ctx, ctxDone, err := aetest.NewContext()
	assert.Nil(t, err)
	defer ctxDone()

	bh := BucketHandler{ ctx }
	_, bucketDone, err := bh.Reader("bucket", "file")
	defer bucketDone()

	assert.Equal(t, err.Error(), "storage: object doesn't exist")
}