package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)


func main() {
    l_router := mux.NewRouter()
    l_router.HandleFunc("/hello", hello).Methods("POST")
    fmt.Println("listening...")
    log.Fatal(http.ListenAndServe(":8000", l_router))
}


type Data struct {
    Message string
}

func hello(p_writer http.ResponseWriter, p_request *http.Request) {
    l_data_receive,_ := ioutil.ReadAll(p_request.Body)
    var l_data_json Data
    json.Unmarshal(l_data_receive, &l_data_json)

    l_response := process_hello(l_data_json)

    json.NewEncoder(p_writer).Encode(l_response)
}

func process_hello(p_data Data) Data {
    var l_response Data
    if p_data.Message == "hello" {
        l_response.Message = "hello you !"
    } else {
        l_response.Message = "hello?"
    }
    return l_response
}