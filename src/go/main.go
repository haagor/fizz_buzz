package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// our main function
func main() {
    l_router := mux.NewRouter()
    l_router.HandleFunc("/hello", hello).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", l_router))
}

func hello(p_writer http.ResponseWriter, p_request *http.Request) {
	fmt.Fprintln(p_writer, "hello you!")
}