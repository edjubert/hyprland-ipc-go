package socket

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go"
	"os"
)

// Remove removes the socket file
func Remove(name string) error {
	return os.Remove(fmt.Sprintf("/tmp/hypr/%s/%s", IPC.GetSignature(), name))
}
