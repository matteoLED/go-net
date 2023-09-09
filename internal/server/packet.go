package server

import (
	"log"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	 "github.com/google/gopacket/pcap"
)

// HandlePacket traite un paquet réseau capturé.
func HandlePacket(packet gopacket.Packet, packetInfo *PacketInfo) {
	
    // Vérifiez si le paquet contient une couche Ethernet
    ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
    if ethernetLayer != nil {
        ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
        packetInfo.EthernetSrcMAC = ethernetPacket.SrcMAC.String()
        packetInfo.EthernetDstMAC = ethernetPacket.DstMAC.String()
    }

    // Ajoutez ici le code pour traiter d'autres couches de paquets (IP, TCP, UDP, etc.)
	
    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        ipPacket, _ := ipLayer.(*layers.IPv4)
        packetInfo.SourceIP = ipPacket.SrcIP.String()
        packetInfo.DestinationIP = ipPacket.DstIP.String()
        packetInfo.TTL = ipPacket.TTL
        packetInfo.Protocol = uint8(ipPacket.Protocol)
        packetInfo.TotalLength = ipPacket.Length
    }

    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer != nil {
        tcpPacket, _ := tcpLayer.(*layers.TCP)
        packetInfo.SourcePort = uint16(tcpPacket.SrcPort)
        packetInfo.DestinationPort = uint16(tcpPacket.DstPort)
        packetInfo.SequenceNumber = tcpPacket.Seq
        packetInfo.AcknowledgmentNumber = tcpPacket.Ack
    }

    udpLayer := packet.Layer(layers.LayerTypeUDP)
    if udpLayer != nil {
        udpPacket, _ := udpLayer.(*layers.UDP)
        packetInfo.SourcePort = uint16(udpPacket.SrcPort)
        packetInfo.DestinationPort = uint16(udpPacket.DstPort)
    }
}


func packetCount(wg *sync.WaitGroup, packetChannel chan int) {
    // Ouvrez l'interface Wi-Fi en mode promiscuité (assurez-vous d'avoir les autorisations nécessaires)
    handle, err := pcap.OpenLive("wlan0", 65536, true, pcap.BlockForever)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    packetCount := 0 // Compteur de paquets Wi-Fi

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        // Vérifiez si la couche est Wi-Fi
        wifiLayer := packet.Layer(layers.LayerTypeDot11)
        if wifiLayer != nil {
            // Vous pouvez extraire des informations Wi-Fi ici
            // Par exemple : adresse MAC source/destination, SSID, etc.
            packetCount++

            // Envoie la valeur mise à jour du compteur sur le canal
            packetChannel <- packetCount
        }
    }

    // Indiquez que cette goroutine est terminée
    wg.Done()
}
