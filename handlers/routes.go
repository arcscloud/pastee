package handlers

func (s *defaultServer) routes() {
    s.router.POST("/api/paste", s.pastePost)
    s.router.GET("/paste/:id", s.pasteGet)

    s.router.Static("/css", "./static/css")
    s.router.Static("/js", "./static/js")
    s.router.GET("/", s.index)
}
