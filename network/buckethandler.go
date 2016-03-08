package network

import (
	"google.golang.org/cloud/storage"
	"golang.org/x/net/context"
)

type BucketHandler struct {
	Context context.Context
}

func (self *BucketHandler) Reader(bucket string, fileName string) (*storage.Reader, func(), error) {
	client, err := storage.NewClient(self.Context)
	if err != nil {
		return nil, nil, err
	}
	bucketHandle := client.Bucket(bucket)
	reader, err := bucketHandle.Object(fileName).NewReader(self.Context)
	return reader, func() {
		if reader != nil { reader.Close() }
		if client != nil { client.Close() }
	}, err
}