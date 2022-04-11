package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

const indexFilename = "./views/**"

func (s defaultServer) index(c *gin.Context) {
    s.router.LoadHTMLGlob(indexFilename)
    c.HTML(http.StatusOK, "index", gin.H{
        "Title":    "Pastee",
        "Subtitle": "Securely share pastes!",
    })
}
