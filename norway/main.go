package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mulander/norway"
)

var basePath string
var cvs norway.Entries

func Init(cvsPath string) {
	entriesFile, err := os.Open(cvsPath + "/CVS/Entries")
	if err != nil {
		log.Fatal(err)
	}
	defer entriesFile.Close()

	cvs, err = norway.ParseEntries(entriesFile)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	data := struct {
		Entries norway.EntriesSorted
	}{
		cvs.SortedByTimestamp(),
	}
	t, err := template.ParseFiles(filepath.Join(basePath, "templates", "root.html"))
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, data)
}

func main() {
	httpListen := flag.String("http", "127.0.0.1:8080", "host:port to listen on")
	flag.StringVar(&basePath, "base", ".", "base path for web interface templates and static resources")
	cvsPath := flag.String("cvs", ".", "path to the cvs repository checkout")
	flag.Parse()

	// http.Handle("/static/", http.FileServer(http.Dir(*basePath)))

	Init(*cvsPath)

	http.HandleFunc("/", Handler)

	if strings.HasPrefix(*httpListen, "127.0.0.1") ||
		strings.HasPrefix(*httpListen, "localhost") {
		log.Print("Bind to your external IP address if you want to share the service with others")
	}

	log.Printf("Open your web browser and visit http://%s/", *httpListen)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}
