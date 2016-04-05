package controller

import (
	"net/http"
	"google.golang.org/appengine"
	"github.com/julienschmidt/httprouter"
	"github.com/mlucchini/githubcompare-api/network"
	"github.com/mlucchini/githubcompare-api/service"
	"io/ioutil"
	"github.com/mlucchini/githubcompare-api/lib"
	"fmt"
)

type LoadController struct {}

const entitiesPerTask = 100

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	ctx := appengine.NewContext(r)
	reader, done, err := (&network.BucketHandler{ Context: ctx }).Reader(bucket, file)
	defer done()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	loadService := service.LoadService{ Context: ctx }
	for group := range lib.GroupLinesIterator(reader, entitiesPerTask) {
		if err = loadService.SendTask(group); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "All tasks created successfully. Follow progress via Task Queue in the App Engine console")
}

func (self *LoadController) Task(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(r)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loadService := service.LoadService{ Context: ctx }
	if _, err = loadService.ReceiveTask(string(payload), entitiesPerTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}