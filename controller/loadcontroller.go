package controller

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/julienschmidt/httprouter"
	"bufio"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/network"
	"github.com/mlucchini/github-compare-backend/service"
)

type LoadController struct {}

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	ctx := appengine.NewContext(r)
	loadService := service.LoadService{ ctx }

	reader, done, err := (&network.BucketHandler{ ctx }).Reader(bucket, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer done()

	added := 0; err = nil
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && err == nil {
		var entity model.RepositoryStarEvent
		entity.Parse(scanner.Text())
		_, err = loadService.Put(&entity)
		added += 1
		if added % 100 == 0 {
			log.Infof(ctx, "Number of events added so far: %d", added)
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Number of events loaded: %d\n", added)
}