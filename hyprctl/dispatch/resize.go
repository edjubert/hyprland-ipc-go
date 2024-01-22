package dispatch

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/internal/socket"
	"github.com/edjubert/hyprland-ipc-go/types"
)

type Resize struct{}

// WindowExactPixel resize given HyprlandClient to specific width and height
func (r *Resize) WindowExactPixel(client types.HyprlandClient, intWidth, intHeight int) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, ResizeWindowPixelExact, intWidth, intHeight, client.Address))
}
