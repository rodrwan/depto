package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baggyapp/depto/cmd/rest/items"
	"github.com/baggyapp/depto/database"
	itemimagessvc "github.com/baggyapp/depto/services/itemimages"
	iteminfosvc "github.com/baggyapp/depto/services/iteminfo"
	itemsvc "github.com/baggyapp/depto/services/items"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := database.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	itemsSvc := itemsvc.NewService(itemsvc.Config{DB: db})
	iteminfoSvc := iteminfosvc.NewService(iteminfosvc.Config{DB: db})
	iteminfoimageSvc := itemimagessvc.NewService(itemimagessvc.Config{DB: db})

	logger := logrus.New()
	logger.Out = os.Stdout

	// save log to file in production.
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}

	mux.Handle("/api/", http.StripPrefix("/api", items.NewRouter(items.Context{
		ItemsService:      itemsSvc,
		ItemInfoService:   iteminfoSvc,
		ItemImagesService: iteminfoimageSvc,
		Logger:            logger,
	})))

	w := logger.Writer()
	defer w.Close()

	listenAddr := ":8000"
	done := make(chan bool)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add logger to handlers
	handler := &logHandler{
		log:  logger,
		next: mux,
	}

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      cors.AllowAll().Handler(handler),
		ErrorLog:     log.New(w, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

type ctxKeyLog struct{}
type ctxKeyRequestID struct{}

type logHandler struct {
	log  *logrus.Logger
	next http.Handler
}

type responseRecorder struct {
	b      int
	status int
	w      http.ResponseWriter
}

func (r *responseRecorder) Header() http.Header { return r.w.Header() }

func (r *responseRecorder) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.w.Write(p)
	r.b += n
	return n, err
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.w.WriteHeader(statusCode)
}

func (lh *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID, _ := uuid.NewRandom()
	ctx = context.WithValue(ctx, ctxKeyRequestID{}, requestID.String())

	start := time.Now()
	rr := &responseRecorder{w: w}
	log := lh.log.WithFields(logrus.Fields{
		"http.req.path":   r.URL.Path,
		"http.req.method": r.Method,
		"http.req.id":     requestID.String(),
		"http.req.host":   r.Host,
	})

	log.Info("request started")
	defer func() {
		log.WithFields(logrus.Fields{
			"http.resp.took_ms": int64(time.Since(start) / time.Millisecond),
			"http.resp.status":  rr.status,
			"http.resp.bytes":   rr.b}).Infof("request complete")
	}()

	ctx = context.WithValue(ctx, ctxKeyLog{}, log)
	r = r.WithContext(ctx)
	lh.next.ServeHTTP(rr, r)
}

// Custom middleware handler logs user agent
func addSecureHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Security-Policy", "default-src 'self'")
		rw.Header().Add("X-Frame-Options", "SAMEORIGIN")
		rw.Header().Add("X-XSS-Protection", "1; mode=block")
		rw.Header().Add("Strict-Transport-Security", "max-age=10000, includeSubdomains; preload")
		rw.Header().Add("X-Content-Type-Options", "nosniff")
		rw.Header().Add("X-Powered-By", "")

		next(rw, r) // Pass on to next middleware handler
	}
}
