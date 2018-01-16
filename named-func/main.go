package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	host := flag.String("host", ":9001", "Port to listen on")
	flag.Parse()

	fmt.Printf("Listening on %s\n", *host)

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/cheese", handleCheese)

	http.ListenAndServe(*host, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index Handler: %s %v", r.Method, r.URL.Path)
	fmt.Fprint(w, `<a href="/cheese">Cheese</a>`)
}

func handleCheese(w http.ResponseWriter, r *http.Request) {
	log.Printf("Cheese Handler: %s %v", r.Method, r.URL.Path)

	w.Header().Add("Content-Type", "text/html")

	t, err := template.New("cheese").Parse(cheeseTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Title, Image string }{
		Title: "Wensleydale",
		Image: `http://vignette2.wikia.nocookie.net/wallaceandgromit/images/5/5d/Wensleydale.jpg`,
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

const cheeseTemplate = `
<html>
<head><title>A Template</title><head>
<body>
<h1>templated {{.Title}} </h1>
<img src="{{.Image}}">
</body>
</html>
`
