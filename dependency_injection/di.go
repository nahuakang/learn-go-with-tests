package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet takes in a name and returns a greeting
// Both os.Stdout and bytes.Buffer implement io.Writer interface
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler greets
// http.ResponseWriter also implements io.Writer interface
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
