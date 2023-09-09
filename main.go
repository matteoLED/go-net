package main

import (
	"fmt"
	"go-net/internal/server"
	"sync"
)

func main() {
	port := 8080
	serverInstance := server.NewServer(port)

	// Créez un WaitGroup pour synchroniser la goroutine de comptage
	var wg sync.WaitGroup
	packetChannel := make(chan int)

	serverInstance.PacketCount(&wg, packetChannel)

	// Lancez la goroutine qui compte les paquets
	wg.Add(1)
	go serverInstance.PacketCount(&wg, packetChannel)

	// Démarrer le serveur
	err := serverInstance.Start()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is running on port %d...\n", port)

	select {}
}
