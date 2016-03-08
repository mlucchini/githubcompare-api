package controller

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/julienschmidt/httprouter"
	"bufio"
	"github.com/mlucchini/github-compare-backend/model"
	"github.com/mlucchini/github-compare-backend/network"
	"github.com/mlucchini/github-compare-backend/service"
	"google.golang.org/appengine/runtime"
	"golang.org/x/net/context"
)

type LoadController struct {}

func (self *LoadController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucket, file := params.ByName("bucket"), params.ByName("file")

	runtime.RunInBackground(appengine.NewContext(r), func(ctx context.Context) {
		loadService := service.LoadService{ ctx }

		reader, done, err := (&network.BucketHandler{ ctx }).Reader(bucket, file)
		if err != nil {
			log.Errorf(ctx, err.Error())
			return
		}
		defer done()

		added := 0; err = nil
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() && err == nil {
			var entity model.RepositoryStarEvent
			entity.Parse(scanner.Text())
			_, err = loadService.Put(&entity)
			if err != nil {
				log.Errorf(ctx, err.Error())
				return
			}
			added += 1
			if added % 100 == 0 {
				log.Infof(ctx, "Number of events added so far: %d", added)
			}
		}
	})

	fmt.Fprintf(w, "Launched loading in background")
}