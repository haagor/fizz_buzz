package main

import "testing"


func TestHello(t *testing.T) {
    var l_data1 Data
    l_data1.Message = "hello"
    var l_data2 Data
    l_data2.Message = "bla bla"

    if process_hello(l_data1).Message != "hello you !" {
        t.Fail()
    }
    if process_hello(l_data2).Message != "hello?" {
        t.Fail()
    }
}