package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/new-endpoint", handleNewEndPoint)

	fmt.Println("added print statement")

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	response := []byte("Hello World!")
	fmt.Println((response)) //  We can use Print as string for better readability

	// Handle both return values
	if _, err := writer.Write(response); err != nil {
		fmt.Println(err)
	}

}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Println("handle health endpoint triggered")
}

func handleNewEndPoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	response := []byte("Hello World!")
	fmt.Println((response)) //  We can use Print as string for better readability

	// Handle both return values
	if _, err := writer.Write(response); err != nil {
		fmt.Println(err)
	}

}
