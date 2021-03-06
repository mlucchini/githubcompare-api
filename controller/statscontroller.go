package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/mlucchini/githubcompare-api/service"
)

type StatsController struct {}

func (self *StatsController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600, s-maxage=3600")

	repositoryName := params.ByName("org") + "/" + params.ByName("repository")

	ctx := appengine.NewContext(r)
	entity, err := (&service.StatsService{ Context: ctx }).GetRepository(repositoryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(entity)
}