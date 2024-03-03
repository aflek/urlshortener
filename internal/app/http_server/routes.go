package httpserver

func (server *UsServer) Routes() {
	server.Router.POST("/", server.CreateShortURL)
	server.Router.GET("/:id", server.FindShortURL)
	server.Router.POST("/api/shorten", server.Shorten)
}
