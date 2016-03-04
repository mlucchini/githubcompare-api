package network

import (
	"google.golang.org/cloud/storage"
	"golang.org/x/net/context"
)

type BucketHandler struct {
	Bucket *storage.BucketHandle
	Context context.Context
}

func (bh *BucketHandler) ReadFile(fileName string) (*storage.Reader, error) {
	return bh.Bucket.Object(fileName).NewReader(bh.Context)
}