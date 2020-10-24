package server

import "net/http"

type Route struct {
	path        string
	handlerFunc http.HandlerFunc
}
