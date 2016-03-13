package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/mlucchini/github-compare-backend/service"
)

type StarsController struct {}

func (self *StarsController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	repositoryName := params.ByName("org") + "/" + params.ByName("repository")

	ctx := appengine.NewContext(r)
	events, err := (&service.StarsService{ ctx }).FilterOnRepositorySortByDate(repositoryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(events)
}
