package network

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/cloud/storage"
)

func TestReadFile(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	client, err := storage.NewClient(ctx)
	bucket := client.Bucket("MyBucket")

	bh := BucketHandler{ Bucket: bucket, Context: ctx }
	bh.ReadFile("MyFile")
}