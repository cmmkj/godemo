package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, t *http.Request) {
	Greet(w, "world")
}

func main() {
	//Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
