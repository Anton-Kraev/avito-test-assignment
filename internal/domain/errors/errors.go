package errors

import "errors"

var (
	ErrNotEnoughMoney  = errors.New("not enough money to complete a transaction")
	ErrInvalidPassword = errors.New("invalid password")
)
