package entry

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
    "encoding/json"
)

const kind = "RepositoryDateStars"

func init() {
    router := httprouter.New()
    router.GET("/api/stars/:org/:repository", handler)
    http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    c := appengine.NewContext(r)

    repositoryName := params.ByName("org") + "/" + params.ByName("repository")
    query := datastore.NewQuery(kind).Filter("RepositoryName =", repositoryName).Order("Date")

    events := make([]RepositoryDateStars, 0)
    if _, err := query.GetAll(c, &events); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    encoder := json.NewEncoder(w)
    encoder.Encode(events)
}
