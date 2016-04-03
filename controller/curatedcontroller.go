package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type CuratedController struct {}

type CuratedEntry struct {
	Category string `json:"category"`
	Repositories []string `json:"repositories"`
}

var curatedEntries = []CuratedEntry{
	CuratedEntry{ "Front-end javascript frameworks", []string{ "angular/angular.js", "facebook/react", "jashkenas/backbone", "emberjs/ember.js" }},
	CuratedEntry{ "Server-side web frameworks", []string{ "meteor/meteor", "rails/rails", "expressjs/express", "laravel/laravel", "pallets/flask", "playframework/playframework" }},
	CuratedEntry{ "Front-end build systems", []string{ "gruntjs/grunt", "gulpjs/gulp", "webpack/webpack" }},
	CuratedEntry{ "Configuration management tools", []string{ "ansible/ansible", "chef/chef", "puppetlabs/puppet", "saltstack/salt", "capistrano/capistrano" }},
	CuratedEntry{ "Unikernels", []string{ "cloudius-systems/osv", "runtimejs/runtime", "mirage/mirage", "rumpkernel/rumprun" }},
	CuratedEntry{ "API gateways", []string{ "strongloop/loopback", "mashape/kong", "TykTechnologies/tyk" }},
	CuratedEntry{ "Containers", []string{ "docker/docker", "cloudfoundry/warden", "cloudfoundry-incubator/garden", "coreos/rkt" }},
	CuratedEntry{ "Design frameworks", []string{ "twbs/bootstrap", "h5bp/html5-boilerplate", "daneden/animate.css", "zurb/foundation-sites", "necolas/normalize.css" }},
	CuratedEntry{ "Text editors", []string{ "atom/atom", "adobe/brackets", "neovim/neovim", "limetext/lime", "textmate/textmate" }},
	CuratedEntry{ "NoSQL databases", []string{ "antirez/redis", "rethinkdb/rethinkdb", "mongodb/mongo", "pouchdb/pouchdb" }},
	CuratedEntry{ "CSS preprocessors", []string{ "less/less.js", "sass/sass", "stylus/stylus" }},
}

func (self *CuratedController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600, s-maxage=3600")

	encoder := json.NewEncoder(w)
	encoder.Encode(curatedEntries)
}