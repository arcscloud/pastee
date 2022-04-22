//go:build prod

package handlers

import (
    "github.com/gin-gonic/gin"
)

func init() {
    gin.SetMode(gin.ReleaseMode)

    eng := gin.Default()
    eng.SetTrustedProxies(nil)

    Engine = eng
}
