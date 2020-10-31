package server

import (
	"github.com/ivanh/meido/store"
	"net/http"
)

type Route struct {
	path        string
	handlerFunc http.HandlerFunc
}

var routes = []Route{
	GetStaticHandle(StaticPath{"D:/", "/d"}),
	{"/items", store.GetItems},
}
