package socket

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go/ipc"
)

// WriteCmd allow you to write any command on the Hyprland socket.
func WriteCmd(cmd string) error {
	conn, err := IPC.ConnectHyprctl(0)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("[ERROR] - closing hyprctl conn", err)
		}
	}()

	if _, err := conn.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
