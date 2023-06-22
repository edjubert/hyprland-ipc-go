package IPC

import (
	"fmt"
	"log"
	"net"
)

func CreateSocket() net.Listener {
	signature := GetSignature()
	socket := "/tmp/hypr/" + signature + "/.gophrland.sock"

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
