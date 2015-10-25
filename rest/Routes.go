package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/locations",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/locations/{locationId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/locations",
		TodoCreate,
	},
	Route{
		"TodoUpdate",
		"PUT",
		"/locations/{locationId}",
		TodoUpdate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/locations/{locationId}",
		TodoDelete,
	},
}
