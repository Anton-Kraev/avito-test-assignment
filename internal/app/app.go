package app

import (
	"github.com/Anton-Kraev/avito-test-assignment/internal/http/handler"
	"github.com/Anton-Kraev/avito-test-assignment/internal/http/server"
	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/logger"
	"github.com/Anton-Kraev/avito-test-assignment/internal/service/auth"
	"github.com/Anton-Kraev/avito-test-assignment/internal/service/info"
	"github.com/Anton-Kraev/avito-test-assignment/internal/service/payment"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

const env = "local"

func Run() {
	log := logger.Setup(env)

	authService := auth.NewService(nil)
	infoService := info.NewService(nil, nil, nil)
	paymentService := payment.NewService(nil, nil, nil, nil)

	srv := server.New(log)
	hndl := handler.New(log, authService, infoService, paymentService)

	api.RegisterHandlers(srv, hndl)

	srv.Logger.Fatal(srv.Start(":8080"))
}
