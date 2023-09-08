package main

import (
	"go-net/internal/server" // Use the correct relative path
)

func main() {
	port := 8080
	server := server.NewServer(port) // Use NewServer here

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
