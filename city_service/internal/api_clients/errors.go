package apiclients

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("an error has occurred and your request was not completed. please try later")
)
