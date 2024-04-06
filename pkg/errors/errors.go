package errors

import "errors"

var (
	InvalidRequest = errors.New("invalid request")
	FailedRequest  = errors.New("failed request")
	DataNotFound   = errors.New("data not found")

	CarExists   = errors.New("car already exists")
	CarNotFound = errors.New("car not found")
)
