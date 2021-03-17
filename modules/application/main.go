package main

import (
	. "view"
)

func main() {
	env := Env{}
	server := NewServer(env)
	// Start server
	server.Logger.Fatal(servere.Start(":8080"))
}
