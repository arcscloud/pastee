PKG := github.com/arcs/pastee/version
GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
		GOOS=linux GOARCH=arm64 go build -o server -ldflags "-X $(PKG).Hash=$(GIT_COMMIT)"

docker: build
		docker build . --platform linux/arm64/v8 --tag arcscloud/pastee

prod: docker