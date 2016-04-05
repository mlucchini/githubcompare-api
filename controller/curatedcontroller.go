package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/mlucchini/githubcompare-api/service"
)

type CuratedController struct {}

func (self *CuratedController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600, s-maxage=3600")

	json, err := (&service.CuratedService{}).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}