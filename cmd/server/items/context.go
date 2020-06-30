package items

import (
	"log"
	"net/http"

	"github.com/baggyapp/depto/pkg/errors"
	"github.com/baggyapp/depto/pkg/responses"
	"github.com/baggyapp/depto/pkg/router"
	"github.com/baggyapp/depto/services"
)

// NewRouter ...
func NewRouter(c Context) http.Handler {
	r := router.New()

	r.GET("/", c.Handle(getItems))
	r.POST("/", c.Handle(createItems))
	// r.GET("/new", c.Handle(formItems))

	r.GET("/:id", c.Handle(getItemInfo))
	r.GET("/:id/info", c.Handle(addItemInfoIndex))
	r.POST("/:id/info", c.Handle(addItemInfo))

	return r
}

// Context ...
type Context struct {
	ItemsService    services.Items
	ItemInfoService services.ItemInfo
}

// Handle ...
func (c Context) Handle(h HandlerFunc) Handler {
	return Handler{
		ctx:    c,
		handle: h,
	}
}

// HandlerFunc ...
type HandlerFunc func(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error)

// Handler ...
type Handler struct {
	ctx    Context
	handle HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := h.handle(h.ctx, w, r)
	if err != nil {
		iErr, ok := err.(*errors.HTML)
		if !ok {
			log.Fatal("could not cast error")
		}

		iErr.Write(w)
		return
	}

	resp.Write(w)
}
