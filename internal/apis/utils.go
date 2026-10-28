package apis

import (
	"net/http"
)

func subRoute(pattern string, mux http.Handler) *http.ServeMux {
	handler := http.NewServeMux()
	handler.Handle(pattern+"/", mux)

	return handler
}
