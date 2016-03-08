package controller

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/cloud/storage"
	"github.com/julienschmidt/httprouter"
	"bufio"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/network"
)

type LoadController struct {}

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	ctx := appengine.NewContext(r)
	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	fmt.Fprintf(w, "Bucket name: %v\n\n", bucket)
	fmt.Fprintf(w, "File name: %v\n\n", file)

	bucketHandler := &network.BucketHandler{ Bucket: client.Bucket(bucket), Context: ctx }
	reader, err := bucketHandler.ReadFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		var repositoryStarEvent model.RepositoryStarEvent
		repositoryStarEvent.Parse(line)
		fmt.Fprintf(w, repositoryStarEvent.String())

		/*key := datastore.NewKey(ctx, kind, "stringID", 0, nil)
		entity := new(RepositoryStarEvent)
		if _, err := datastore.Put(ctx, key, entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}*/
	}
}