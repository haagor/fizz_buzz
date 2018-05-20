#!/bin/bash

echo "launch server :"
go run go/main.go &
sleep 2

echo -e "\nrun curl :"
echo -e "\n[bad url]"
curl -H "Content-Type: application/json" -XPOST "http://localhost:8000/fizz_buzz" -d '{"string1":"fizz","string2":"buzz","int1":3,"int2":5,"limit":15}'
echo -e "\n[good curl]"
curl -H "Content-Type: application/json" -XPOST "http://localhost:8000/fizz_buzz/v1" -d '{"string1":"fizz","string2":"buzz","int1":3,"int2":5,"limit":15}'
echo -e "\n[bad type inside Json]"
curl -H "Content-Type: application/json" -XPOST "http://localhost:8000/fizz_buzz/v1" -d '{"string1":"fizz","string2":"buzz","int1":3,"int2":5,"limit":"100"}'

echo -e "\n"
sleep 2
kill 0