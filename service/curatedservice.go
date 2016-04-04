package service

import (
	"io/ioutil"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

type CuratedService struct {
	Context context.Context
	json []byte
}

func (self *CuratedService) GetAll() ([]byte, error) {
	if self.json == nil {
		json, err := ioutil.ReadFile("../data/curated.json")
		if err != nil {
			log.Errorf(self.Context, "Failed to read the curated entries: %s", err.Error())
			return nil, err
		}
		self.json = json
	}
	return self.json, nil
}