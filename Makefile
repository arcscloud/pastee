PKG := github.com/arcs/pastee/version
GIT_COMMIT := $(shell git rev-list -1 HEAD)

build-prod:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o pastee -tags prod -ldflags "-X $(PKG).Hash=$(GIT_COMMIT)"

build-dev:
		go build -o pastee -ldflags "-X $(PKG).Hash=$(GIT_COMMIT)"

docker: build-prod
		docker build . --tag arcscloud/pastee

prod: docker

server: build-dev
		./pastee server

cleanup-dev: build-dev
		./pastee cleanup
