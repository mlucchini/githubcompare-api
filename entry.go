package entry

import (
	"github.com/mlucchini/githubcompare/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Run() {
	statsController := controller.StatsController{}
	loadController := controller.LoadController{}
	curatedController := controller.CuratedController{}

	router := httprouter.New()
	router.GET("/api/stats/:org/:repository", statsController.Get)
	router.GET("/api/curated", curatedController.GetAll)
	router.POST("/api/admin/load/:bucket/:file", loadController.Update)
	router.POST("/api/admin/loadtask", loadController.Task)

	http.Handle("/", router)
}