package main

import "fmt"

import "github.com/mtebourbi/lbc-fizzbuzz/pkg/server"

func main() {
	fmt.Println("FizzBuzz web service")
	server.ListenAndServe()
}
