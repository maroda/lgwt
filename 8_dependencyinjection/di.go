package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// this first line demonstrates that we can send an interface here via bytes.Buffer
// but if we send os.Stdout, the function won't work. so instead,
// we tell it to use the interface that those types implement.
// func Greet(writer *bytes.Buffer, name string) {
func Greet(writer io.Writer, name string) {
	// This needs to be fmt.Fprintf because
	// it takes a Writer (interface)
	// which is implemented by bytes.Buffer
	// so the string ends up there to test instead of stdout (via fmt.Printf)
	// 	in other words, we tell Fprintf to fill the buffer 'writer' with the print
	fmt.Fprintf(writer, "Hello, %s", name)
	// fmt.Printf("Hello, %s", name) // <<< this just goes to stdout
}

/*
// This will work, because we can pass the function anything that implements Writer
func main() {
	Greet(os.Stdout, "Elodie")
}
*/

// in this example, http.ResponseWriter implements the Writer interface,
// so it can also be passed to Greet just like bytes.Buffer and os.Stdout
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
