package main

import (
	lissajous "../lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles := 5
	if cyclesStr := r.FormValue("cycles"); cyclesStr != "" {
		if value, err := strconv.Atoi(cyclesStr); err != nil {
			cycles = value
		}
	}
	lissajous.Lissajous(w, cycles)
}
