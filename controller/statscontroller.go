package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/mlucchini/github-compare-backend/service"
)

type StatsController struct {}

func (self *StatsController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	repositoryName := params.ByName("org") + "/" + params.ByName("repository")

	ctx := appengine.NewContext(r)
	entity, err := (&service.StatsService{ ctx }).GetRepository(repositoryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(entity)
}