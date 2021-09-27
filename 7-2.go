package main

import (
	"fmt"
	"log"
	"net/http"
)

type textHandler struct {
	responseText string
}

func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

type indexHandler struct {
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html> 
         <html>
	    	<head>
		 		<title>Biomass Commodity Exchange</title>
	    	</head>
			<body>
				<b>Wellcome to BCEX</b>
        		<p>
          		<a href="/menu1">page 1</a> | <a href="/menu2">page 2</a> | <a href="/menu3">page 3</a> | <a href="/menu4">page 4</a>
        		</p>
			</body>
		 </html>`
	fmt.Fprintf(w, html)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})

	menu1 := &textHandler{"page 111"}
	mux.Handle("/menu1", menu1)

	menu2 := &textHandler{"page 222"}
	mux.Handle("/menu2", menu2)

	menu3 := &textHandler{"page 333"}
	mux.Handle("/menu3", menu3)

	menu4 := &textHandler{"page 444"}
	mux.Handle("/menu4", menu4)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
