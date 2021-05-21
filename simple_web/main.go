package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"fmt"
	"flag"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

// Initiate web server
func main() {
	// https://zetcode.com/golang/flag/
	var port string
	flag.StringVar(&port, "port", "9100", "the port the server will listen on")
	flag.Parse()

	router := router()
	fmt.Println("Starting WebServer on port " + port)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
