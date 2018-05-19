package main

import (
    "testing"
    "strings"
)


func TestFizzBuzz1(t *testing.T) {
    var l_params1 Params
    l_params1.String1 = "foo"
    l_params1.String2 = "bar"
    l_params1.Int1 = 2
    l_params1.Int2 = 3
    l_params1.Limit = 12

    l_res1 := []string{"1","foo","bar","foo","5","foobar","7","foo","bar","foo","11","foobar"}
    if strings.Compare(strings.Join(l_res1, ","), strings.Join(process_fizz_buzz(l_params1).Result, ",")) != 0 {
        t.Fail()
    }
}

func TestFizzBuzz2(t *testing.T) {
    var l_params2 Params
    l_params2.String1 = "fizz"
    l_params2.String2 = "buzz"
    l_params2.Int1 = 3
    l_params2.Int2 = 5
    l_params2.Limit = 15

    l_res2 := []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"}
    if strings.Compare(strings.Join(l_res2, ","), strings.Join(process_fizz_buzz(l_params2).Result, ",")) != 0 {
        t.Fail()
    }
}