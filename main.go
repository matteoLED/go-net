package main

import (
	"fmt"
	"go-net/internal/server"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	port := 8080
	serverInstance := server.NewServer(port)

	packet := gopacket.NewPacket([]byte{0, 1, 2, 3}, layers.LayerTypeEthernet, gopacket.Default)

	server.NetworkHandler(packet)

	err := serverInstance.Start()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is running on port %d...\n", port)

	select {}
}
