package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get all environment variables and print them
		for _, e := range os.Environ() {
			fmt.Fprintf(w, e+"\n")
		}
	})

	http.ListenAndServe(":80", nil)
}
