package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	root := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hiii")
	}

	http.HandleFunc("/", root)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
