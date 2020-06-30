package inventory

import (
	"log"
	"net/http"

	"github.com/baggyapp/depto/pkg/responses"
	"github.com/baggyapp/depto/pkg/router"
	"github.com/baggyapp/depto/services"
)

// NewRouter ...
func NewRouter(c Context) http.Handler {
	r := router.New()

	r.GET("/:id", c.Handle(getInventory))
	r.POST("/", c.Handle(createInventory))
	r.POST("/:id/add", c.Handle(addItemToInventory))
	r.DELETE("/:id", c.Handle(deleteInventory))
	// r.PATCH("/:id", c.Handle(updateItems))

	return r
}

// Context ...
type Context struct {
	InventoryService      services.Inventory
	InventoryItemsService services.InventoryItem
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
		log.Printf("[ERROR]: unexpected unhandled error: %v", err)
		return
	}

	resp.Write(w)
}
