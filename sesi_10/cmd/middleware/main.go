package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	greet := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}
	endendpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endendpoint)))

	fmt.Println("Listening to port 5000")

	err := http.ListenAndServe(":5000", mux)

	log.Fatal(err)
}
func middleware1(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 1")
		next.ServeHTTP(w, r)
	})

}
func middleware2(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 2")
		next.ServeHTTP(w, r)
	})

}
