package main

import "net/http"

import "fmt"

import "math/rand"

//CustomServeMux is a struct which can be a multiplexer
type CustomServeMux struct {
}

//this is the function handler to be overriden
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	//any struct that has serveHTTP function can be a multiplexer
	// mux := &CustomServeMux{}
	// http.ListenAndServe(":8000", mux)
	http.ListenAndServe(":8000", newMux)
}
