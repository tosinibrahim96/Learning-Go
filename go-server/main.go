package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm error %v", err)
		return
	}

	fmt.Fprintf(res, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)

}

func helloHandler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/hello" {
		http.Error(res, "Page not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(res, "hello world!")

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting port at port 8001\n")

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
