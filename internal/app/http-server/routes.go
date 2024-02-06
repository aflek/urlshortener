package httpserver

func (server *UsApp) Routes() {
	server.Router.POST("/", server.CreateShortUrl)
	server.Router.GET("/:id", server.FindShortUrl)
}
