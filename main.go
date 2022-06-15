package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	money := addValues(10, 40)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 10 + 40 = %d", money))
}

func Divide(w http.ResponseWriter, r *http.Request) {

	f, err := divideValues(100.0, 0.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f is divided by %f is %f", 100.0, 0.0, f))

}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the application on line: %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
