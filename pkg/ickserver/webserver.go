// spa web server
package ickserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type WebServer struct {
	staticfiledir      string
	http_port          string
	http_rwTimeout     int
	http_idleTimeout   int
	http_cache_control bool
	http_logger        bool

	WebRouter *mux.Router
	ApiRouter *mux.Router
}

func MakeWebserver() WebServer {
	ws := new(WebServer)

	// get .env variables
	ws.staticfiledir = strings.ToLower(strings.Trim(os.Getenv("SPA_STATICFILEDIR"), " "))
	if ws.staticfiledir == "" {
		ws.staticfiledir = "./web/static"
	}

	ws.http_port = strings.Trim(os.Getenv("HTTP_PORT"), " ")
	if ws.http_port == "" {
		ws.http_port = "5432"
	}

	ws.http_rwTimeout, _ = strconv.Atoi(os.Getenv("HTTP_RWTIMEOUT"))
	if ws.http_rwTimeout <= 0 {
		ws.http_rwTimeout = 15
	}

	ws.http_idleTimeout, _ = strconv.Atoi(os.Getenv("HTTP_IDLETIMEOUT"))
	if ws.http_idleTimeout <= 0 {
		ws.http_idleTimeout = 15
	}

	ws.http_cache_control = false
	if strings.ToLower(strings.Trim(os.Getenv("HTTP_CACHE_CONTROL"), " ")) == "true" {
		ws.http_cache_control = true
	}

	ws.http_logger = false
	if strings.ToLower(strings.Trim(os.Getenv("HTTP_LOGGER"), " ")) == "true" {
		ws.http_logger = true
	}

	// configure the server, with or without trailing slash is the same route
	ws.WebRouter = mux.NewRouter().StrictSlash(true)

	// configure the /api subrouter
	ws.ApiRouter = ws.WebRouter.PathPrefix("/api").Subrouter()
	ws.ApiRouter.HandleFunc("/health", GetHealthHandle())

	return *ws
}

func (ws WebServer) Run() {

	// let's go
	fmt.Printf("Starting the SPA serving assets from %q and /api on port %s\n", ws.staticfiledir, ws.http_port)

	// the main handler serving spa static files
	// force content-type header for wasm files
	ws.WebRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		http.FileServer(http.Dir(ws.staticfiledir)).ServeHTTP(w, r)
	})

	// add middleware to remove cache if requested in config file
	if !ws.http_cache_control {
		fmt.Println("spa server: no-cache forced in response header")
		ws.WebRouter.Use(middlewareNoCache)
	}

	// setup timeouts
	srv := &http.Server{
		Addr:         ws.http_port,
		WriteTimeout: time.Duration(ws.http_rwTimeout) * time.Second,
		ReadTimeout:  time.Duration(ws.http_rwTimeout) * time.Second,
		IdleTimeout:  time.Duration(ws.http_idleTimeout) * time.Second,
	}

	// add middleware to log every request
	if ws.http_logger {
		fmt.Println("spa server: http logger is on")
		srv.Handler = newLogger(ws.WebRouter)
	} else {
		srv.Handler = ws.WebRouter
	}

	// listen and serve in a go routine to allow catching shutdown clean request in parallel
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL will not be caught.
	chansig := make(chan os.Signal, 1)
	signal.Notify(chansig, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	// Block until we receive a shutdown signal.
	<-chansig

	// Start the clean shutdown process.
	// Create a deadline to wait for, longer than the rwTimeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ws.http_rwTimeout)+time.Second*10)
	defer func() {
		signal.Stop(chansig) // clean stop listening os
		cancel()             // ensure clean cancel the context, so write ctx.Done()
	}()

	// Doesn't block if no connections,
	// but will otherwise wait clean shutdown until the timeout deadline.
	srv.Shutdown(ctx)

	fmt.Println("SPA web Server is down")
}

// GetHealthHandle responds to a GET Health api request
func GetHealthHandle() func(http.ResponseWriter, *http.Request) {
	counter := 0
	return func(w http.ResponseWriter, r *http.Request) {
		counter++
		json.NewEncoder(w).Encode(map[string]string{"health": "live", "counter": strconv.Itoa(counter)})
	}
}
