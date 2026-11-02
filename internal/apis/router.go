package apis

import (
	"net/http"
	"slices"
	"strings"
)

type Route struct {
	mux         *http.ServeMux
	prefix      string
	middlewares []func(http.Handler) http.Handler
}

func NewRoute(
	mux *http.ServeMux,
	prefix string,
	middlewares ...func(http.Handler) http.Handler,
) Route {
	return Route{
		mux:         mux,
		prefix:      prefix,
		middlewares: middlewares,
	}
}

func (rg Route) Handle(
	pattern string,
	handler func(http.ResponseWriter, *http.Request),
	middlewares ...func(http.Handler) http.Handler,
) {
	method := ""
	route := ""
	patternFields := strings.Fields(pattern)
	if len(patternFields) < 1 || len(patternFields) > 2 {
		panic("Invalid pattern: " + pattern)
	} else if len(patternFields) == 2 {
		method = patternFields[0] + " "
		route = patternFields[1]
	} else {
		route = patternFields[0]
	}

	next := http.Handler(http.HandlerFunc(handler))

	for _, v := range slices.Backward(middlewares) {
		next = v(next)
	}

	for _, v := range slices.Backward(rg.middlewares) {
		next = v(next)
	}

	rg.mux.HandleFunc(method+rg.prefix+route, next.ServeHTTP)
}

func (rg Route) Base(
	method string,
	handler func(http.ResponseWriter, *http.Request),
	middlewares ...func(http.Handler) http.Handler,
) {
	next := http.Handler(http.HandlerFunc(handler))

	for _, v := range slices.Backward(middlewares) {
		next = v(next)
	}

	for _, v := range slices.Backward(rg.middlewares) {
		next = v(next)
	}

	rg.mux.HandleFunc(method+" "+rg.prefix, next.ServeHTTP)
}

func (rg Route) Sub(
	prefix string,
	middlewares ...func(http.Handler) http.Handler,
) Route {
	return Route{
		mux:         rg.mux,
		prefix:      rg.prefix + prefix,
		middlewares: append(rg.middlewares, middlewares...),
	}
}
