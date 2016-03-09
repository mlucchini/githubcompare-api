package controller

import (
	"net/http"
	"google.golang.org/appengine"
	"github.com/julienschmidt/httprouter"
	"github.com/mlucchini/github-compare-backend/network"
	"github.com/mlucchini/github-compare-backend/service"
	"io/ioutil"
	"github.com/mlucchini/github-compare-backend/lib"
	"fmt"
)

type LoadController struct {}

const entitiesPerTask = 500

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	ctx := appengine.NewContext(r)
	reader, done, err := (&network.BucketHandler{ ctx }).Reader(bucket, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer done()

	loadService := service.LoadService{ ctx }
	for group := range lib.GroupLinesIterator(reader, entitiesPerTask) {
		if err = loadService.SendTask(group); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "All tasks created successfully. Follow progress via the App Engine console")
}

func (self *LoadController) Task(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(r)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loadService := service.LoadService{ ctx }
	if _, err = loadService.ReceiveTask(string(payload), entitiesPerTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}