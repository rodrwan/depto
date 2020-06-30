package router

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	httprouter.Router
}

func New() *Router {
	return &Router{
		Router: httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
		},
	}
}

func (r *Router) GET(path string, handler http.Handler) {
	r.Handle("GET", path, handler)
}

func (r *Router) HEAD(path string, handler http.Handler) {
	r.Handle("HEAD", path, handler)
}

func (r *Router) OPTIONS(path string, handler http.Handler) {
	r.Handle("OPTIONS", path, handler)
}

func (r *Router) POST(path string, handler http.Handler) {
	r.Handle("POST", path, handler)
}

func (r *Router) PUT(path string, handler http.Handler) {
	r.Handle("PUT", path, handler)
}

func (r *Router) PATCH(path string, handler http.Handler) {
	r.Handle("PATCH", path, handler)
}

func (r *Router) DELETE(path string, handler http.Handler) {
	r.Handle("DELETE", path, handler)
}

func (r *Router) Handle(method, path string, handler http.Handler) {
	r.Router.Handle(method, path, r.wrapHandler(handler))
}

func (r *Router) wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := r.Context()
		for _, p := range params {
			ctx = context.WithValue(ctx, p.Key, p.Value)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
