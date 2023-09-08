package server

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// HandlePacket traite un paquet réseau capturé.
func HandlePacket(packet gopacket.Packet) {
	// Vérifiez si le paquet contient une couche Ethernet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Printf("Ethernet Source MAC: %s, Destination MAC: %s\n", ethernetPacket.SrcMAC, ethernetPacket.DstMAC)
	}

	// Ajoutez ici le code pour traiter d'autres couches de paquets (IP, TCP, UDP, etc.)

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ipPacket, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("Source IP: %s, Destination IP: %s\n", ipPacket.SrcIP, ipPacket.DstIP)
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("Source Port: %d, Destination Port: %d\n", tcpPacket.SrcPort, tcpPacket.DstPort)
	}

	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udpPacket, _ := udpLayer.(*layers.UDP)
		fmt.Printf("Source Port: %d, Destination Port: %d\n", udpPacket.SrcPort, udpPacket.DstPort)
	}
}
