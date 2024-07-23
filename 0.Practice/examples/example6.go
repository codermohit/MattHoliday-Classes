/*
Program which imitates the http modules handler interface and HandlerFunc type
Instead of sending responses , it prints the pattern string and writes the name
to the specified io.Writer
*/
package main

import (
	"fmt"
	"io"
	"os"
)

// imitates handler interface {serverHTTP(ResponseWriter, *Request)}
type handlerI interface {
	//imitates the serveHTTP()
	serveMsg(w io.Writer, name string)
}

// imitates the HandlerFunc func(ResponseWriter, *Request)
type HandlerFuncF func(w io.Writer, name string)

// imitates serverHTTP method on HandlerFunc
func (f HandlerFuncF) serveMsg(w io.Writer, name string) {
	f(w, name)
}

func handleFunc(pattern string, handler handlerI, name string) {
	//original http.handleFunc registers the handler parameter to the DefaultServerMux
	//but here we will call the serveMsg on the passed in handler
	fmt.Println("Pattern string : ", pattern)
	handler.serveMsg(os.Stdout, name)
}

func main() {
	// Define a function that matches the HandlerFuncF type
	myHandler := HandlerFuncF(func(w io.Writer, name string) {
		fmt.Fprintf(w, "Name: %s\n", name)
	})

	// Use handleFunc with the defined handler and pass your name
	handleFunc("/greet", myHandler, "Gopher")
}
