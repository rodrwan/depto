package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Handler func(*Context)

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (c *Context) JSON(code int, body interface{}) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.WriteHeader(code)

	b, _ := json.Marshal(body)

	c.ResponseWriter.Write(b)
}

type Route2 struct {
	Pattern *regexp.Regexp
	Handler Handler
	Method  string
}

type Router2 struct {
	Routes map[string]map[string]Route2
}

func New2() *Router2 {
	return &Router2{
		Routes: make(map[string]map[string]Route2, 0),
	}
}

func (r *Router2) GET(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "GET"}

	r.Handle("GET", path, route)
}

func (r *Router2) HEAD(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "HEAD"}

	r.Handle("HEAD", path, route)
}

func (r *Router2) OPTIONS(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "OPTIONS"}

	r.Handle("OPTIONS", path, route)
}

func (r *Router2) POST(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "POST"}

	r.Handle("POST", path, route)
}

func (r *Router2) PUT(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "PUT"}

	r.Handle("PUT", path, route)
}

func (r *Router2) PATCH(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "PATCH"}

	r.Handle("PATCH", path, route)
}

func (r *Router2) DELETE(path string, handler Handler) {
	re := regexp.MustCompile(path)
	route := Route2{Pattern: re, Handler: handler, Method: "DELETE"}

	r.Handle("DELETE", path, route)
}

func (r *Router2) Handle(method, path string, route Route2) {
	r.Routes[method][path] = route
}

func (r *Router2) ServeHTTP(w http.ResponseWriter, rr *http.Request) {
	ctx := &Context{Request: rr, ResponseWriter: w}

	for _, rt := range r.Routes[rr.Method] {
		if rt.Method != rr.Method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Println(rt.Pattern.String())
		fmt.Println(ctx.URL.Path)
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			fmt.Println(ctx.Params)
			rt.Handler(ctx)
			return
		}
	}
}
