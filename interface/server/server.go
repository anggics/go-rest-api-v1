package server

import (
	"rest-api/interface/container"

	"github.com/gin-gonic/gin"
)

type server struct {
	engine *gin.Engine
	port   string
}

func StartApp(container container.Container) *server {
	app := gin.Default()

	handler := setupHandler(&container)
	setupRouter(app, handler)

	port := ":9000"
	return &server{
		engine: app,
		port:   port,
	}
}

func (s *server) Run() {
	//fmt.Println(fmt.Sprintf("Server running in addr : http://localhost%s/", s.port))
	s.engine.Run(s.port)
}
