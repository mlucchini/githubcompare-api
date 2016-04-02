default: build

vet:
	go vet ./...

build: vet
	go build ./...

test: build
	goapp test ./...

serve:
	cd appengine; goapp serve

deploy: test
	cd frontend; npm run build; cp -R www/ ../appengine/web
	cd appengine; goapp deploy

update_queue:
	cd appengine; appcfg.py update_queues

eslint:
	cd frontend; ./node_modules/eslint/bin/eslint.js app/** --ext .jsx,.js

.PHONY: vet build test serve deploy update_queue eslint