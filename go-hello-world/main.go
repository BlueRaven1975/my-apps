package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func goQuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, quote.Go())
}

func optQuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, quote.Opt())
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/go-quote", goQuoteHandler)
	http.HandleFunc("/opt-quote", optQuoteHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
