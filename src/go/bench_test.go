package main

import (
    "testing"
)


var g_params1 = Params{"foo", "bar", 2, 3, 1000}

func BenchmarkFizzBuzz1(b *testing.B) {
    for c_n := 0; c_n < b.N; c_n++ {
        process_fizz_buzz(g_params1)
    }
}

func BenchmarkFizzBuzzNaive1(b *testing.B) {
    for c_n := 0; c_n < b.N; c_n++ {
        process_fizz_buzz_naive(g_params1)
    }
}

var g_params2 = Params{"foo", "bar", 20, 30, 1000}

func BenchmarkFizzBuzz2(b *testing.B) {
    for c_n := 0; c_n < b.N; c_n++ {
        process_fizz_buzz(g_params2)
    }
}

func BenchmarkFizzBuzzNaive2(b *testing.B) {
    for c_n := 0; c_n < b.N; c_n++ {
        process_fizz_buzz_naive(g_params2)
    }
}

