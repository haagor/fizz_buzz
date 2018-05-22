# Fizz_Buzz

*"Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”*

## Getting Started


### Prerequisites

Install [Go](https://golang.org/dl/)

### Installing

Clone the repository into `$GOPATH/src/github.com/`

To check my project you can execute the scenario `fizz_buzz/src$ ./scenario.sh`

## Running the tests

```
fizz_buzz/src/go$ go test -v
=== RUN   TestFizzBuzz1
--- PASS: TestFizzBuzz1 (0.00s)
=== RUN   TestFizzBuzz2
--- PASS: TestFizzBuzz2 (0.00s)
=== RUN   TestFizzBuzzNaive1
--- PASS: TestFizzBuzzNaive1 (0.00s)
=== RUN   TestFizzBuzzNaive2
--- PASS: TestFizzBuzzNaive2 (0.00s)
PASS
ok  	github.com/fizz_buzz/src/go	0.002s

```

## Run server

Run with Go :

```
fizz_buzz/src/go$ go run fizz_buzz.go 
listening...

```

Or execute binary :

```
fizz_buzz/bin$ ./fizz_buzz 
listening...

```


## Details

### REST API

Here an example of command curl :

`curl -H "Content-Type: application/json" -XGET "http://localhost:8000/fizz_buzz/v1" -d '{"string1":"fizz","string2":"buzz","int1":3,"int2":5,"limit":15}'`

### Naming convention
My naming convention that I use is, if we take the variable "hello" for example :

* `g_hello` global variable
* `l_hello` local variable
* `c_hello` loop variable
* `p_hello` parameter

Then I use snake_case convention, even if there is a collision with the test naming convention of Go.

### Tests
My tests miss the REST API. They test the business part of fizz_buzz. The next step will be to mock the http.Request to implement the API tests. With these tests I should have a code coverage near of 100%.
Waiting that, I implemented a Bash script scenario to check my REST API with `curl`.

### Dependency gestion
I use [dep](https://golang.github.io/dep/docs/daily-dep.html) to manage my dependencies.

### Go bench
`fizz_buzz/src/go$ go test -bench=. -benchtime=60s`
I wanted to initialize a benchmark with Go. So I write another version of fizz_buzz. I named it with the keyword "naive" because I thinking it was less optimize. My first benchmark didn't confirm my intuition. Indeed I believed reduce the amount of test but with a small `int1`, the naive version will leave the if statements offen at the second. My first version, in all case, will execute all of its if statement.

The other point could be the string concatenation. I tried with `bytes.Buffer` but the difference was not perceptible. This is something that I could more investigate later.

```
fizz_buzz/src/go$ go test -bench=. -benchtime=60s
goos: linux
goarch: amd64
pkg: github.com/fizz_buzz/src/go
BenchmarkFizzBuzz1-4        	 2000000	     58235 ns/op
BenchmarkFizzBuzzNaive1-4   	 2000000	     50160 ns/op
BenchmarkFizzBuzz2-4        	 1000000	     64501 ns/op
BenchmarkFizzBuzzNaive2-4   	 1000000	     70405 ns/op
PASS
ok  	github.com/fizz_buzz/src/go	462.988s
```

## Authors

* **Paris Simon** 
