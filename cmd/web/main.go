package main

import (
	"fmt"
	"github.com/ebonsage/learngo/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
