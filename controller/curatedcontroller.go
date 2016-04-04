package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type CuratedController struct {}

type CuratedEntry struct {
	Category string `json:"category"`
	Description string `json:"description"`
	Repositories []string `json:"repositories"`
}

var curatedEntries = []CuratedEntry{
	CuratedEntry{ "Front-end Javascript frameworks", "Some tools that help you build single-page applications", []string{ "angular/angular.js", "facebook/react", "jashkenas/backbone", "emberjs/ember.js" }},
	CuratedEntry{ "Server-side web frameworks", "Productivity and performance to build scalable web applications", []string{ "meteor/meteor", "rails/rails", "expressjs/express", "laravel/laravel", "pallets/flask", "playframework/playframework" }},
	CuratedEntry{ "Front-end build systems", "Simple automation", []string{ "gruntjs/grunt", "gulpjs/gulp", "webpack/webpack" }},
	CuratedEntry{ "Configuration Management tools", "Server automation framework and applications", []string{ "ansible/ansible", "chef/chef", "puppetlabs/puppet", "saltstack/salt", "capistrano/capistrano" }},
	CuratedEntry{ "Unikernels", "Open-source library operating systems", []string{ "cloudius-systems/osv", "runtimejs/runtime", "mirage/mirage", "rumpkernel/rumprun" }},
	CuratedEntry{ "API gateways", "Scalable, high-performance open-source API layers", []string{ "strongloop/loopback", "mashape/kong", "TykTechnologies/tyk" }},
	CuratedEntry{ "Containers", "Container engines to pack, ship and run your applications as lightweight containers", []string{ "docker/docker", "cloudfoundry/warden", "cloudfoundry-incubator/garden", "coreos/rkt" }},
	CuratedEntry{ "Design frameworks", "Popular HTML, CSS, and JavaScript frameworks for developing responsive, mobile first projects on the web", []string{ "twbs/bootstrap", "h5bp/html5-boilerplate", "daneden/animate.css", "zurb/foundation-sites", "necolas/normalize.css" }},
	CuratedEntry{ "Text editors", "Some amazingly awesome open source editors", []string{ "atom/atom", "adobe/brackets", "neovim/neovim", "limetext/lime", "textmate/textmate" }},
	CuratedEntry{ "NoSQL databases", "Non-relational, distributed, open-source and horizontally scalable databases", []string{ "antirez/redis", "rethinkdb/rethinkdb", "mongodb/mongo", "pouchdb/pouchdb" }},
	CuratedEntry{ "CSS preprocessors", "Helps you write maintainable, future-proof CSS code", []string{ "less/less.js", "sass/sass", "stylus/stylus" }},
}

func (self *CuratedController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600, s-maxage=3600")

	encoder := json.NewEncoder(w)
	encoder.Encode(curatedEntries)
}