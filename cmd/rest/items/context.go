package items

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/baggyapp/depto/pkg/errors"
	"github.com/baggyapp/depto/pkg/responses"
	"github.com/baggyapp/depto/pkg/router"
	"github.com/baggyapp/depto/services"
	"github.com/sirupsen/logrus"
)

// NewRouter ...
func NewRouter(c Context) http.Handler {
	r := router.New()

	r.POST("/items", c.Handle(createItems))
	r.GET("/items", c.Handle(getItems))
	r.GET("/items/:id", c.Handle(getItem))
	r.PATCH("/items/:id", c.Handle(updateItems))
	r.POST("/items/:id/info", c.Handle(addInfoToItem))
	r.DELETE("/items/:id", c.Handle(deleteItems))

	return r
}

// Context ...
type Context struct {
	ItemsService      services.Items
	ItemInfoService   services.ItemInfo
	ItemImagesService services.ItemImages

	Logger *logrus.Logger
}

// Handle ...
func (c Context) Handle(h HandlerFunc) Handler {
	return Handler{
		log:    c.Logger,
		ctx:    c,
		handle: h,
	}
}

// Handle2 ...
func (c Context) Handle2(h HandlerFunc) router.Handler {
	return func(ctx *router.Context) {
		rCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
		defer cancel()

		// in case of client disconnect, cancel context
		if cn, ok := ctx.ResponseWriter.(http.CloseNotifier); ok {
			go func() {
				<-cn.CloseNotify()
				fmt.Println("Cancelling request")
				cancel()
			}()
		}

		resp, err := h(c, ctx.ResponseWriter, ctx.Request.WithContext(rCtx))
		if err != nil {
			aErr, ok := err.(*errors.API)
			if !ok {
				log.Printf("[ERROR]: unexpected unhandled error: %v", err)
				return
			}

			aErr.Write(c.Logger, ctx.ResponseWriter)
			return
		}

		ctx.JSON(resp.Status, resp)
	}
}

// HandlerFunc ...
type HandlerFunc func(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error)

// Handler ...
type Handler struct {
	ctx    Context
	log    *logrus.Logger
	handle HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// in case of client disconnect, cancel context
	if cn, ok := w.(http.CloseNotifier); ok {
		go func() {
			<-cn.CloseNotify()
			fmt.Println("Cancelling request")
			cancel()
		}()
	}

	resp, err := h.handle(h.ctx, w, r.WithContext(rCtx))
	if err != nil {
		aErr, ok := err.(*errors.API)
		if !ok {
			log.Printf("[ERROR]: unexpected unhandled error: %v", err)
			return
		}

		aErr.Write(h.log, w)
		return
	}

	if err := resp.Write(w); err != nil {
		log.Printf("[ERROR]: %v, encoding response: %v", err, resp)
	}
}
