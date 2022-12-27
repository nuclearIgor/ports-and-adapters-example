package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nuclearIgor/go-hexagonal/adapters/web/handlers"
	"github.com/nuclearIgor/go-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handlers.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
