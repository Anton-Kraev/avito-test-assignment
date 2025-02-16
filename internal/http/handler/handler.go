package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) PostApiAuth(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetApiBuyItem(ctx echo.Context, item string) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetApiInfo(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostApiSendCoin(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
