package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

func main() {
	httpListen := flag.String("http", "127.0.0.1:8080", "host:port to listen on")
	// basePath := flag.String("base", ".", "base path for web interface templates and static resources")
	// cvsPath := flag.String("cvs", ".", "path to the cvs repository checkout")
	flag.Parse()

	// http.Handle("/static/", http.FileServer(http.Dir(*basePath)))

	if strings.HasPrefix(*httpListen, "127.0.0.1") ||
		strings.HasPrefix(*httpListen, "localhost") {
		log.Print("Bind to your external IP address if you want to share the service with others")
	}

	log.Printf("Open your web browser and visit http://%s/", *httpListen)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}
