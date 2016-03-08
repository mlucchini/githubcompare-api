package entry

import (
	"github.com/mlucchini/github-compare-backend/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Run() {
	starsController := controller.StarsController{}
	loadController := controller.LoadController{}

	router := httprouter.New()
	router.GET("/api/stars/:org/:repository", starsController.Get)
	router.POST("/api/storage/:bucket/:file", loadController.Update)
	http.Handle("/", router)
}