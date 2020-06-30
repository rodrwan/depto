package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/baggyapp/depto/cmd/rest/inventory"
	"github.com/baggyapp/depto/cmd/server/items"
	"github.com/baggyapp/depto/database"
	"github.com/baggyapp/depto/database/postgres"
	inventorysvc "github.com/baggyapp/depto/services/inventory"
	inventoryitemsvc "github.com/baggyapp/depto/services/inventoryitem"
	iteminfosvc "github.com/baggyapp/depto/services/iteminfo"
	itemsvc "github.com/baggyapp/depto/services/items"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	db, err := database.NewPostgres("postgres://user:pass@localhost:5435/inventory?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	itemsSvc := itemsvc.Service{
		Store: postgres.ItemsStore{
			DB: db,
		},
	}
	iteminfoSvc := iteminfosvc.Service{
		Store: postgres.ItemInfoStore{
			DB: db,
		},
	}
	inventorySvc := inventorysvc.Service{
		Store: postgres.InventoryStore{
			DB: db,
		},
	}
	invetoryItemsSvc := inventoryitemsvc.Service{
		Store: postgres.InventoryItemStore{
			DB: db,
		},
	}

	mux.Handle("/items/", http.StripPrefix("/items", items.NewRouter(items.Context{
		ItemsService:    itemsSvc,
		ItemInfoService: iteminfoSvc,
	})))

	mux.Handle("/inventory", inventory.NewRouter(inventory.Context{
		InventoryService:      inventorySvc,
		InventoryItemsService: invetoryItemsSvc,
	}))

	logger := logrus.New()

	listenAddr := ":9000"
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	w := logger.Writer()
	defer w.Close()
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      cors.AllowAll().Handler(mux),
		ErrorLog:     log.New(w, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Info("Server is ready to handle requests at ", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Info("Server stopped")
}
