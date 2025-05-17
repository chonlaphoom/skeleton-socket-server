package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func main() {
	flag.Parse()

	fmt.Printf("static serve at %s", *addr)
	http.HandleFunc("/", serveApp)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
