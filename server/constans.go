package server

import "net/http"

const (
	// 3xx Redirection
	/* 302 */ FOUND = http.StatusFound

	// 2xx Success
	/* 200 */ OK = http.StatusOK
	/* 201 */ CREATED = http.StatusCreated

	// 4xx Client Error
	/* 400 */ BAD_REQUEST = http.StatusBadRequest
	/* 401 */ UNAUTHORIZED = http.StatusUnauthorized
	/* 404 */ NOT_FOUND = http.StatusNotFound

	// 5xx Server Error
	/* 500 */ INTERNAL_SERVER_ERROR = http.StatusInternalServerError
)

