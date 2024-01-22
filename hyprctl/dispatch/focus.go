package dispatch

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/internal/socket"
	"github.com/edjubert/hyprland-ipc-go/types"
)

type Focus struct{}

// WorkspaceID focuses a HyprlandWorkspace.Id
func (f *Focus) WorkspaceID(ID int) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %d", DispatchKey, Workspace, ID))
}

// Window focuses a given HyprlandClient.Address
func (f *Focus) Window(address string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s address:%s", DispatchKey, FocusWindow, address))
}

// Monitor focuses a given HyprlandMonitor
func (f *Focus) Monitor(monitor types.HyprlandMonitor) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %s", DispatchKey, FocusMonitor, monitor.Name))
}
