package IPC

import (
	"fmt"
	"log"
	"net"
)

func CreateSocket(name string) net.Listener {
	socket := fmt.Sprintf("/tmp/hypr/%s/%s", GetSignature(), name)

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
