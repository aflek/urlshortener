package httpserver

func (server *UsServer) Routes() {
	server.Router.POST("/", server.CreateShortURL)
	server.Router.GET("/:id", server.FindShortURL)
}
