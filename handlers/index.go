package handlers

import (
    "github.com/arcs/pastee/utl"
    "github.com/gin-gonic/gin"
    "net/http"
)

const indexFilename = "./views/**"

func (s defaultServer) index(c *gin.Context) {
    s.router.LoadHTMLGlob(indexFilename)

    hash, _ := utl.Run("git", "rev-parse", "HEAD")

    c.HTML(http.StatusOK, "index", gin.H{
        "Title":    "Pastee",
        "Subtitle": "Securely share pastes!",
        "Hash":     string(hash)[:8],
    })
}
