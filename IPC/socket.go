package IPC

import (
	"fmt"
	"log"
	"net"
	"os"
)

// RemoveSocket removes the socket file
func RemoveSocket(name string) error {
	return os.Remove(fmt.Sprintf("/tmp/hypr/%s/%s", GetSignature(), name))
}

// CreateSocket creates the socket file
func CreateSocket(name string) net.Listener {
	socket := fmt.Sprintf("/tmp/hypr/%s/%s", GetSignature(), name)

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
