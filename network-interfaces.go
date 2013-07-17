package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type NetworkInterface struct {
	Name       string   `json:"name"`
	MacAddress string   `json:"macaddress"`
	Addrs      []string `json:"ipaddress"`
}

func NetworkInterfacesToMap() (networkInterfaceMap map[string]NetworkInterface) {
	networkInterfaceMap = make(map[string]NetworkInterface)

	networkInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, networkInterface := range networkInterfaces {
		var n NetworkInterface
		n.Name = networkInterface.Name
		n.MacAddress = networkInterface.HardwareAddr.String()

		addrs, _ := networkInterface.Addrs()
		for _, addr := range addrs {
			n.Addrs = append(n.Addrs, addr.String())
		}

		networkInterfaceMap[networkInterface.Name] = n
	}

	return
}

func main() {
	networkInterfaceMap := NetworkInterfacesToMap()
	networkInterfaceJSON, err := json.MarshalIndent(networkInterfaceMap, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(networkInterfaceJSON))
}
