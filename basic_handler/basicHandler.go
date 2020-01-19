package main

import (
	"io"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hellow world!\n")
}

func main() {
	http.HandleFunc("/hello", MyHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// curl -X GET http://localhost:8000/hello
