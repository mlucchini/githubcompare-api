default: test

get:
	go get ./...

build: get
	go build ./...

test: build
	go test ./... -cover

serve:
	cd appengine; goapp serve

update_queue:
	cd appengine; appcfg.py update_queues

.PHONY: get build test serve update_queue
