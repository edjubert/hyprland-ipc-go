package socket

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go"
	"log"
	"net"
)

// Create creates the socket file
func Create(name string) net.Listener {
	socket := fmt.Sprintf("/tmp/hypr/%s/%s", IPC.GetSignature(), name)

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
