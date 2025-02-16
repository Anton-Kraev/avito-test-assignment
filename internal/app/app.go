package app

import (
	"github.com/Anton-Kraev/avito-test-assignment/internal/http/handler"
	"github.com/Anton-Kraev/avito-test-assignment/internal/http/server"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

func Run() {
	srv := server.New()
	hndl := handler.New()

	api.RegisterHandlers(srv, hndl)

	srv.Logger.Fatal(srv.Start(":8080"))
}
