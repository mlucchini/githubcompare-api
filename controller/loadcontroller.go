package controller

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	"github.com/julienschmidt/httprouter"
	"bufio"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/network"
	"github.com/mlucchini/github-compare-backend/service"
	"golang.org/x/net/context"
	"strings"
	"io/ioutil"
)

type LoadController struct {}

const entitiesPerTask = 1000
const defaultQueue = ""
const separator = "\n"

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	ctx := appengine.NewContext(r)
	reader, done, err := (&network.BucketHandler{ ctx }).Reader(bucket, file)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return
	}
	defer done()

	lines := make([]string, 0, entitiesPerTask)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && err == nil {
		lines = append(lines, scanner.Text())
		if (len(lines) == entitiesPerTask) {
			err = self.sendTask(w, ctx, lines)
			lines = lines[:0]
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		self.sendTask(w, ctx, lines)
	}
}

func (self *LoadController) sendTask(w http.ResponseWriter, ctx context.Context, elements []string) (error) {
	payload := strings.Join(elements, separator)
	t := &taskqueue.Task{ Path: "/api/loadtask", Payload: []byte(payload), Header: nil, Method: "POST" }
	if _, err := taskqueue.Add(ctx, t, defaultQueue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	log.Infof(ctx, "Added task with %d elements", len(elements))
	return nil
}

func (self *LoadController) Task(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(r)
	payload, err := ioutil.ReadAll(r.Body)

	events := strings.Split(string(payload), separator)
	entities := make([]*model.RepositoryStarEvent, 0, entitiesPerTask)

	for _, event := range events {
		var entity model.RepositoryStarEvent
		if err := entity.Parse(event); err != nil {
			log.Errorf(ctx, err.Error())
			return
		}
		entities = append(entities, &entity)
	}

	_, err = (&service.LoadService{ ctx }).PutMulti(entities)
	if err != nil {
		log.Errorf(ctx, err.Error())
	}
}