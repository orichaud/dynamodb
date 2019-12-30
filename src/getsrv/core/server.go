package core

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

const (
	HOST = "0.0.0.0:8080"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer(context *Context) *Server {
	r := mux.NewRouter()

	// /item...
	getHandler := func(w http.ResponseWriter, r *http.Request) {
		HandleGetItem(context, w, r)
	}
	r.HandleFunc("/item/{id:[0-9]+}/{subkey}", getHandler).
		Methods("GET").
		Schemes("http")

	// items...
	scanHandler := func(w http.ResponseWriter, r *http.Request) {
		HandleScanItems(context, w, r)
	}
	r.HandleFunc("/items", scanHandler).
		Queries("max", "{max}").
		Methods("GET").
		Schemes("http")
	r.HandleFunc("/items", scanHandler).
		Methods("GET").
		Schemes("http")
	getItemsForIdHandler := func(w http.ResponseWriter, r *http.Request) {
		HandleGetItems(context, w, r)
	}
	r.HandleFunc("/items/{id:[0-9]+}", getItemsForIdHandler).
		Methods("GET").
		Schemes("http")

	// Heath check
	healthzHandler := func(w http.ResponseWriter, r *http.Request) {
		HandleHealthz(context, w, r)
	}
	r.HandleFunc("/healthz", healthzHandler).
		Methods("GET").
		Schemes("http")

	// Error 404
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		msg := struct {
			Message string `json:Message`
		}{
			Message: "no resource found: 404",
		}

		if err := Send(msg, w); err != nil {
			glog.Errorf("Cannot transfer JSON payload: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	httpSrv := &http.Server{
		Handler:      r,
		Addr:         HOST,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	return &Server{HttpServer: httpSrv}
}

func (server *Server) Start() {
	glog.Info("server is ready to handle request")
	if err := server.HttpServer.ListenAndServe(); err != nil {
		glog.Error("failed to start the HTTP web server", err)
	}
}
