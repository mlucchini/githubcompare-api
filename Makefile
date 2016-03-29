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
	cd appengine; goapp deploy

.PHONY: vet build test serve deploy