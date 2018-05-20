package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)


func main() {
    l_router := mux.NewRouter()
    l_router.HandleFunc("/fizz_buzz/v1", fizz_buzz).Methods("POST")
    fmt.Println("listening...")
    log.Fatal(http.ListenAndServe(":8000", l_router))
}


type Params struct {
    String1 string
    String2 string
    Int1    int
    Int2    int
    Limit   int
}

type Response struct {
    Result []string
}

func fizz_buzz(p_writer http.ResponseWriter, p_request *http.Request) {
    l_params_receive,_ := ioutil.ReadAll(p_request.Body)
    var l_params_json Params
    if err := json.Unmarshal(l_params_receive, &l_params_json); err != nil {
        fmt.Println("error:", err)
    }

    l_response := process_fizz_buzz(l_params_json)

    json.NewEncoder(p_writer).Encode(l_response)
}

func process_fizz_buzz(p_fizz_buzz Params) Response {
    l_response := Response{Result: make([]string, p_fizz_buzz.Limit)}
    for c_i := 1; c_i <= p_fizz_buzz.Limit; c_i++ {
        l_res := ""
        if c_i % p_fizz_buzz.Int1 == 0 {
            l_res += p_fizz_buzz.String1
        } 
        if c_i % p_fizz_buzz.Int2 == 0 {
            l_res += p_fizz_buzz.String2
        }
        if l_res == "" {
            l_res = strconv.Itoa(c_i)
        }
        l_response.Result[c_i-1] = l_res
    }
    return l_response
}

func process_fizz_buzz_naive(p_fizz_buzz Params) Response {
    l_response := Response{Result: make([]string, p_fizz_buzz.Limit)}
    l_res := ""
    for c_i := 1; c_i <= p_fizz_buzz.Limit; c_i++ {
        if (c_i % (p_fizz_buzz.Int1 * p_fizz_buzz.Int2)) == 0 {
            l_res = p_fizz_buzz.String1 + p_fizz_buzz.String2
        } else if c_i % p_fizz_buzz.Int1 == 0 {
            l_res = p_fizz_buzz.String1
        } else if c_i % p_fizz_buzz.Int2 == 0 {
            l_res = p_fizz_buzz.String2
        } else {
            l_res = strconv.Itoa(c_i)
        }
        l_response.Result[c_i-1] = l_res
    }
    return l_response
}
