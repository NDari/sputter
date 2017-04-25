all: main

main: assets
	go build

test: main glide
	golint `glide novendor`
	go vet `glide novendor`
	go test `glide novendor`

assets: bindata
	go-bindata -o docstring/assets.go -pkg="docstring" \
	-prefix="docstring/" docstring/*.md

glide:
	go get github.com/Masterminds/glide

bindata:
	go get github.com/jteeuwen/go-bindata/...

lint:
	go get -u github.com/golang/lint/golint

init: glide bindata lint
	glide install

install: init main test
	go install
