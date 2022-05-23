PKG := github.com/arcs/pastee/version
GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
		go build -o server -tags prod -ldflags "-X $(PKG).Hash=$(GIT_COMMIT)"

docker: build
		docker build . --tag arcscloud/pastee

prod: docker

dev:
		go build -o server -ldflags "-X $(PKG).Hash=$(GIT_COMMIT)" && ./server
