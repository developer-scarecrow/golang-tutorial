package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(writer http.ResponseWriter, r *http.Request) {
	// プリントする
	fmt.Fprintf(writer, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)

	// クロージャを渡しても良い
	// http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
	// 	// プリントする
	// 	fmt.Fprintf(writer, "Hello World")
	// })

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
