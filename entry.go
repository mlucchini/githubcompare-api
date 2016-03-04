package entry

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/mlucchini/github-compare-backend/controller"
)

func init() {
    starsController := controller.StarsController{}
    storageController := controller.StorageController{}

    router := httprouter.New()
    router.GET("/api/stars/:org/:repository", starsController.Get)
    router.POST("/api/storage/:bucket/:file", storageController.Update)
    http.Handle("/", router)
}