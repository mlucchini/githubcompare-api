default: test

vet:
	go vet ./...

build: vet
	go build ./...

test: build
	goapp test ./...

serve:
	cd appengine; goapp serve

update_queue:
	cd appengine; appcfg.py update_queues

esbuild:
	cd frontend; npm run build

eslint:
	cd frontend; npm run lint

deploy: test esbuild eslint
	cp -R frontend/www/ appengine/web
	cd appengine; goapp deploy

.PHONY: vet build test serve update_queue esbuild eslint deploy