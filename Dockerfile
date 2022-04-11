#FROM golang:1.18-buster as backend_builder
#
#WORKDIR /app
#
#COPY go.* ./
#RUN go mod download
#
#COPY ./handlers /app/handlers
#COPY ./rand /app/rand
#COPY ./store /app/store
#COPY main.go ./
#
#RUN GOOS=linux GOARCH=amd64 \
#  go build \
#  -o /app/server \
#  ./main.go
#
#FROM arm32v6/alpine:3.15
#
#RUN apk add --no-cache bash
#
#COPY --from=backend_builder /app/server /app/server
#COPY ./docker_entrypoint /app/docker_entrypoint
#COPY ./static /app/static
#COPY ./views /app/views
#
#WORKDIR /app
#
#ENTRYPOINT ["/app/docker_entrypoint"]

#FROM golang:1.18-bullseye
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#
#COPY . .
#
#CMD ["/app/server"]

FROM arm64v8/alpine
COPY . /app
CMD ["/app/server"]