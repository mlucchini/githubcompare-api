package entry

import (
	"github.com/mlucchini/github-compare-backend/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Run() {
	statsController := controller.StatsController{}
	loadController := controller.LoadController{}

	router := httprouter.New()
	router.GET("/api/stats/:org/:repository", statsController.Get)
	router.POST("/api/admin/load/:bucket/:file", loadController.Update)
	router.POST("/api/admin/loadtask", loadController.Task)

	http.Handle("/", router)
}