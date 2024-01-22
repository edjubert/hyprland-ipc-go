package get

import (
	"encoding/json"
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/types"
	"os/exec"
)

type Get struct{}

// ActiveClient returns the active HyprlandClient
func (g *Get) ActiveClient() (types.HyprlandClient, error) {
	clientJSON, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		return types.HyprlandClient{}, err
	}

	var activewindow types.HyprlandClient
	if err := json.Unmarshal(clientJSON, &activewindow); err != nil {
		return types.HyprlandClient{}, err
	}

	return activewindow, nil
}

// WorkspaceFloatingClients returns all HyprlandClient that have HyprlandClient.Floating parameter set to true
func (g *Get) WorkspaceFloatingClients(workspace types.HyprlandWorkspace) ([]types.HyprlandClient, error) {
	clients, err := g.Clients()
	if err != nil {
		return nil, err
	}

	var workspaceClients []types.HyprlandClient
	for _, client := range clients {
		if client.Workspace.Id == workspace.Id {
			workspaceClients = append(workspaceClients, client)
		}
	}

	return workspaceClients, nil
}

// Clients returns all HyprlandClient
func (g *Get) Clients() ([]types.HyprlandClient, error) {
	clientsJSON, err := exec.Command("hyprctl", "clients", "-j").Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot execute command -> %w\n", err)
	}

	var clients []types.HyprlandClient
	if err := json.Unmarshal(clientsJSON, &clients); err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot unmarshall clients -> %w\n", err)
	}

	return clients, nil
}

// ClientByPID returns the HyprlandClient by its HyprlandClient.Pid
func (g *Get) ClientByPID(clients []types.HyprlandClient, pid int) (types.HyprlandClient, error) {
	for _, client := range clients {
		if client.Pid == pid {
			return client, nil
		}
	}

	return types.HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

// ClientByClassName returns the HyprlandClient by its HyprlandClient.Class
func (g *Get) ClientByClassName(clients []types.HyprlandClient, class string) (types.HyprlandClient, error) {
	for _, client := range clients {
		if client.Class == class {
			return client, nil
		}
	}

	return types.HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

// Monitors returns all HyprlandMonitor
func (g *Get) Monitors(format string) ([]types.HyprlandMonitor, error) {
	if format != "" && format != "-j" {
		return nil, fmt.Errorf("[ERROR] - wrong monitor formats")
	}

	monitorsJSON, err := exec.Command("hyprctl", "monitors", format).Output()
	if err != nil {
		return nil, err
	}

	var monitors []types.HyprlandMonitor
	if err := json.Unmarshal(monitorsJSON, &monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}

// ActiveMonitor returns the active HyprlandMonitor
func (g *Get) ActiveMonitor(monitors []types.HyprlandMonitor) (types.HyprlandMonitor, error) {
	for _, monitor := range monitors {
		if monitor.Focused {
			return monitor, nil
		}
	}

	return types.HyprlandMonitor{}, fmt.Errorf("[ERROR] - not found")
}

// ActiveWorkspace returns the active HyprlandWorkspace
func (g *Get) ActiveWorkspace() (types.HyprlandWorkspace, error) {
	activeClient, err := g.ActiveClient()
	if err != nil {
		return types.HyprlandWorkspace{}, err
	}

	return activeClient.Workspace, nil
}

// Workspaces returns all HyprlandWorkspace
func (g *Get) Workspaces() ([]types.HyprlandWorkspace, error) {
	ret, err := exec.Command("hyprctl", "workspaces", "-j").Output()
	if err != nil {
		return nil, err
	}

	var workspaces []types.HyprlandWorkspace
	if err := json.Unmarshal(ret, &workspaces); err != nil {
		return nil, err
	}

	fmt.Println(workspaces)
	return workspaces, nil
}

// MonitorByID return HyprlandMonitor for given HyprlandMonitor.Id
func (g *Get) MonitorByID(monitorId int) (types.HyprlandMonitor, error) {
	monitors, err := g.Monitors("-j")
	if err != nil {
		return types.HyprlandMonitor{}, err
	}

	for _, monitor := range monitors {
		if monitor.Id == monitorId {
			return monitor, nil
		}
	}

	return types.HyprlandMonitor{}, fmt.Errorf("[ERROR] - Could not find monitor %d\n", monitorId)
}
