package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := flag.String("host", ":9001", "Port to listen on")
	static := flag.String("static", "static", "location for static content")
	flag.Parse()

	fmt.Printf("Listening on %s\nStatic in %s\n", *host, *static)

	var c Cheese

	c.Image = "stilton.jpg"
	c.Title = "Stilton Blue"

	http.Handle("/cheese", c)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(*host, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index Handler: %s %v", r.Method, r.URL.Path)
	fmt.Fprint(w, `<a href="/cheese">Cheese</a>`)
}

////////////////////////////////////////
// cheese.go
//

// Cheese handler
type Cheese struct {
	//Image of the product
	Image string `json:"image"`
	//Title for the page
	Title string
}

// ServeHTTP handler, all your cheese needs are here
func (c Cheese) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Cheese Handler: %s %v", r.Method, r.URL.Path)

	if r.URL.Query().Get("f") == "json" {
		c.jsonHandler(w, r)
	} else {
		c.htmlHandler(w, r)
	}
}

func (c Cheese) htmlHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Cheese Handler: %s %v", r.Method, r.URL.Path)

	w.Header().Add("Content-Type", "text/html")

	fmt.Fprintf(w, `<h1>%s</h1><img src="/static/%s">`, c.Title, c.Image)
}

func (c Cheese) jsonHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Cheese Handler: %s %v", r.Method, r.URL.Path)

	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	if err := enc.Encode(c); err != nil {
		log.Printf("Unable to encode json response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
