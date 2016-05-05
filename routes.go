package main

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"net/http"
)

type Route struct {
	name    string
	path    string
	method  string
	handler httprouter.Handle
	ctx     context.Context
	storage storage.Storage
}

var routes = []Route{
	{name: "createSecret", path: "/secret", method: "POST", handler: CreateSecret},
	{name: "readSecret", path: "/secret/:uuid", method: "GET", handler: ReadSecret},
	{name: "updateSecret", path: "/secret/:uuid", method: "PATCH", handler: UpdateSecret},
	{name: "deleteSecret", path: "/secret/:uuid", method: "DELETE", handler: DeleteSecret},
}
