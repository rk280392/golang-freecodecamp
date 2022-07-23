package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Your name is %s \n", name)
	fmt.Fprintf(w, "Your address is %s \n", address)

}

func main() {
	port := 8700
	fileserver := http.FileServer(http.Dir("./static")) //httpd.Dir converts string into http.Dir type. This type has one method "OPEN"
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port %v\n", port)
	if err := http.ListenAndServe(":8700", nil); err != nil {
		log.Fatal(err)
	}
}
