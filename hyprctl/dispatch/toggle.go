package dispatch

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/internal/socket"
)

type Toggle struct{}

// Floating activate/deactivate the HyprlandClient.Floating mode for the given HyprlandClient.Address
func (t *Toggle) Floating(address string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s address:%s", DispatchKey, ToggleFloating, address))
}

// SpecialWorkspace show/hide the special HyprlandWorkspace.Name
func (t *Toggle) SpecialWorkspace(name string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %s", DispatchKey, ToggleSpecialWorkspace, name))
}
