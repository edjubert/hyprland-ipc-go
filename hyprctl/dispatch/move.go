package dispatch

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/hyprctl/get"
	"github.com/edjubert/hyprland-ipc-go/internal/socket"
	"github.com/edjubert/hyprland-ipc-go/types"
	"math/rand"
)

type Move struct{}

// WindowPixelExact move at precise HyprlandClient.At the HyprlandClient.Address
func (m *Move) WindowPixelExact(x, y int, address string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, MoveWindowPixelExact, x, y, address))
}

// ToWorkspaceName moves a given HyprlandClient.Address to a HyprlandWorkspace.Name and focus the HyprlandClient
func (m *Move) ToWorkspaceName(workspaceName, clientAddress string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %s,address:%s", DispatchKey, MoveToWorkspace, workspaceName, clientAddress))
}

// ToWorkspaceSilent moves a given HyprlandClient.Address to a HyprlandWorkspace.Name without focussing the HyprlandClient
func (m *Move) ToWorkspaceSilent(name, address string) error {
	return socket.WriteCmd(fmt.Sprintf("%s %s %s,address:%s", DispatchKey, MoveToWorkspaceSilent, name, address))
}

// ToSpecialNamed moves given HyprlandClient.Address to named special HyprlandWorkspace
func (m *Move) ToSpecialNamed(specialWorkspaceName, clientAddress string) error {
	if specialWorkspaceName != "" {
		specialWorkspaceName = fmt.Sprintf(":%s", specialWorkspaceName)
	}
	return m.ToWorkspaceSilent("special"+specialWorkspaceName, clientAddress)
}

// ClientToCurrent moves a given HyprlandClient.Address to current HyprlandWorkspace
func (m *Move) ClientToCurrent(address string) error {
	getter := get.Get{}
	monitors, err := getter.Monitors("-j")
	if err != nil {
		return err
	}

	monitor, err := getter.ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	if err := m.ToWorkspaceName(monitor.ActiveWorkspace.Name, address); err != nil {
		return err
	}

	return nil
}

// CenterFloatingClient put a HyprlandClient at the center of a HyprlandMonitor.
// This applies an offset so windows are not stacked on the exact same position
func (m *Move) CenterFloatingClient(client types.HyprlandClient, monitor types.HyprlandMonitor, applyRand bool) error {
	margin := 100
	randFactorX := client.Size[0]
	randFactorY := client.Size[1]
	randX := rand.Intn(randFactorX)
	randY := rand.Intn(randFactorY)
	if !applyRand {
		randFactorX = 0
		randFactorY = 0
		randX = 0
		randY = 0
	}
	centerX := (monitor.X + monitor.Width - monitor.Width/2) - client.Size[0]/2 - randFactorX/2 + randX
	centerY := (monitor.Y + monitor.Height - monitor.Height/2) - client.Size[1]/2 - randFactorY/2 + randY + margin

	return socket.WriteCmd(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, MoveWindowPixelExact, centerX, centerY, client.Address))
}
