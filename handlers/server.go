package handlers

import (
    "github.com/arcs/pastee/store"
    "github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
    Engine = gin.Default()
}

type Server interface {
    Router() *gin.Engine
}

type defaultServer struct {
    router *gin.Engine
    store  store.Store
}

func (s defaultServer) Router() *gin.Engine {
    return s.router
}

func New() Server {
    s := defaultServer{
        router: Engine,
        store:  store.New(),
    }
    s.routes()

    return s
}
