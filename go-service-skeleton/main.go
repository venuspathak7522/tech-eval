package main

import (
	"sre.qlik.com/palindrome/server"
)

func main() {
	// we could also load from config files
	// set up all the dependencies of the server by calling NewServer
	s := server.NewServer()
	// start the server
	s.Start()
}
