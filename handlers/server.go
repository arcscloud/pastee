package handlers

import (
    "github.com/gin-gonic/gin"

    "github.com/arcs/pastee/store"
)

var Engine *gin.Engine

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
    Engine = gin.Default()
    s := defaultServer{
        router: Engine,
        store:  store.New(),
    }
    s.routes()

    return s
}
