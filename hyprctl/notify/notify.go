package notify

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/internal/socket"
)

const (
	NotifyPrefix = ""
)

// SendNotification sends a Hyprland notification
func SendNotification(time int, msgType, msg string) error {
	icon := -1

	switch msgType {
	case "warning":
		icon = 0
	case "info":
		icon = 1
	case "notice":
		icon = 2
	case "error":
		icon = 3
	case "question":
		icon = 4
	case "checkmark":
		icon = 5
	default:
		icon = -1
	}

	color := "rgb(ff1ea3)"

	return socket.WriteCmd(fmt.Sprintf("notify %d %d %s %s: %s", icon, time, color, NotifyPrefix, msg))
}
