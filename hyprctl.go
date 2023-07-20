package IPC

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
)

// GetActiveClient returns the active HyprlandClient
func GetActiveClient() (HyprlandClient, error) {
	clientJSON, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		return HyprlandClient{}, err
	}

	var activewindow HyprlandClient
	if err := json.Unmarshal(clientJSON, &activewindow); err != nil {
		return HyprlandClient{}, err
	}

	return activewindow, nil
}

// GetWorkspaceFloatingClients returns all HyprlandClient that have HyprlandClient.Floating parameter set to true
func GetWorkspaceFloatingClients(workspace HyprlandWorkspace) ([]HyprlandClient, error) {
	clients, err := GetClients()
	if err != nil {
		return nil, err
	}

	var workspaceClients []HyprlandClient
	for _, client := range clients {
		if client.Workspace.Id == workspace.Id {
			workspaceClients = append(workspaceClients, client)
		}
	}

	return workspaceClients, nil
}

// GetClients returns all HyprlandClient
func GetClients() ([]HyprlandClient, error) {
	clientsJSON, err := exec.Command("hyprctl", "clients", "-j").Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot execute command -> %w\n", err)
	}

	var clients []HyprlandClient
	if err := json.Unmarshal(clientsJSON, &clients); err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot unmarshall clients -> %w\n", err)
	}

	return clients, nil
}

// GetClientByPID returns the HyprlandClient by its HyprlandClient.Pid
func GetClientByPID(clients []HyprlandClient, pid int) (HyprlandClient, error) {
	for _, client := range clients {
		if client.Pid == pid {
			return client, nil
		}
	}

	return HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

// GetClientByClassName returns the HyprlandClient by its HyprlandClient.Class
func GetClientByClassName(clients []HyprlandClient, class string) (HyprlandClient, error) {
	for _, client := range clients {
		if client.Class == class {
			return client, nil
		}
	}

	return HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

// Monitors returns all HyprlandMonitor
func Monitors(format string) ([]HyprlandMonitor, error) {
	if format != "" && format != "-j" {
		return nil, fmt.Errorf("[ERROR] - wrong monitor formats")
	}

	monitorsJSON, err := exec.Command("hyprctl", "monitors", format).Output()
	if err != nil {
		return nil, err
	}

	var monitors []HyprlandMonitor
	if err := json.Unmarshal(monitorsJSON, &monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}

// ActiveMonitor returns the active HyprlandMonitor
func ActiveMonitor(monitors []HyprlandMonitor) (HyprlandMonitor, error) {
	for _, monitor := range monitors {
		if monitor.Focused {
			return monitor, nil
		}
	}

	return HyprlandMonitor{}, fmt.Errorf("[ERROR] - not found")
}

// SendNotification sends a Hyprland notification
func SendNotification(time int, msgType, msg string) error {
	icon := -1
	prefix := "  [Gophrland]"

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
	return runHyprctlCmd(fmt.Sprintf("notify %d %d %s %s: %s", icon, time, color, prefix, msg))
}

func runHyprctlCmd(cmd string) error {
	conn, err := ConnectHyprctl(0)
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

// MoveWindowPixelExact move at precise HyprlandClient.At the HyprlandClient.Address
func MoveWindowPixelExact(x, y int, address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch movewindowpixel exact %d %d,address:%s", x, y, address))
}

// ToggleFloating activate/deactivate the HyprlandClient.Floating mode for the given HyprlandClient.Address
func ToggleFloating(address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch togglefloating address:%s", address))
}

// ToggleSpecialWorkspace show/hide the special HyprlandWorkspace.Name
func ToggleSpecialWorkspace(name string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch togglespecialworkspace %s", name))
}

// CenterFloatingClient put a HyprlandClient at the center of a HyprlandMonitor.
// This applies an offset so windows are not stacked on the exact same position
func CenterFloatingClient(client HyprlandClient, monitor HyprlandMonitor, applyRand bool) error {
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

	return runHyprctlCmd(fmt.Sprintf("dispatch movewindowpixel exact %d %d,address:%s", centerX, centerY, client.Address))
}

// MoveToCurrent moves a given HyprlandClient.Address to current HyprlandWorkspace
func MoveToCurrent(address string) error {
	monitors, err := Monitors("-j")
	if err != nil {
		return err
	}

	monitor, err := ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	if err := MoveClientToWorkspaceIDSilent(monitor.ActiveWorkspace.Id, address); err != nil {
		return err
	}

	return nil
}

// FocusWorkspaceID focuses a HyprlandWorkspace.Id
func FocusWorkspaceID(ID int) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch workspace %d", ID))
}

// MoveClientToWorkspaceIDSilent moves a given HyprlandClient.Address to given HyprlandWorkspace.Id id but don't focus it
func MoveClientToWorkspaceIDSilent(workspaceID int, clientAddress string) error {
	return MoveClientToWorkspaceSilent(strconv.Itoa(workspaceID), clientAddress)
}

// FocusWindow focuses a given HyprlandClient.Address
func FocusWindow(address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch focuswindow address:%s", address))
}

// FocusMonitor focuses a given HyprlandMonitor
func FocusMonitor(monitor HyprlandMonitor) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch focusmonitor %s", monitor.Name))
}

// MoveClientToWorkspaceName moves a given HyprlandClient.Address to a HyprlandWorkspace.Name and focus the HyprlandClient
func MoveClientToWorkspaceName(workspaceName, clientAddress string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch movetoworkspace %s,clientAddress:%s", workspaceName, clientAddress))
}

// MoveClientToWorkspaceSilent moves a given HyprlandClient.Address to a HyprlandWorkspace.Name without focussing the HyprlandClient
func MoveClientToWorkspaceSilent(name, address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch movetoworkspacesilent %s,address:%s", name, address))
}

// GetActiveWorkspace returns the active HyprlandWorkspace
func GetActiveWorkspace() (HyprlandWorkspace, error) {
	activeClient, err := GetActiveClient()
	if err != nil {
		return HyprlandWorkspace{}, err
	}

	return activeClient.Workspace, nil
}

// GetWorkspaces returns all HyprlandWorkspace
func GetWorkspaces() ([]HyprlandWorkspace, error) {
	ret, err := exec.Command("hyprctl", "workspaces", "-j").Output()
	if err != nil {
		return nil, err
	}

	var workspaces []HyprlandWorkspace
	if err := json.Unmarshal(ret, &workspaces); err != nil {
		return nil, err
	}

	fmt.Println(workspaces)
	return workspaces, nil
}

// MoveToSpecialNamed moves given HyprlandClient.Address to named special HyprlandWorkspace
func MoveToSpecialNamed(specialWorkspaceName, clientAddress string) error {
	if specialWorkspaceName != "" {
		specialWorkspaceName = fmt.Sprintf(":%s", specialWorkspaceName)
	}
	return MoveClientToWorkspaceSilent("special"+specialWorkspaceName, clientAddress)
}

// GetMonitorByID return HyprlandMonitor for given HyprlandMonitor.Id
func GetMonitorByID(monitorId int) (HyprlandMonitor, error) {
	monitors, err := Monitors("-j")
	if err != nil {
		return HyprlandMonitor{}, err
	}

	for _, monitor := range monitors {
		if monitor.Id == monitorId {
			return monitor, nil
		}
	}

	return HyprlandMonitor{}, fmt.Errorf("[ERROR] - Could not find monitor %d\n", monitorId)
}

// ResizeWindowPixel resize given HyprlandClient to specific width and height
func ResizeWindowPixel(client HyprlandClient, intWidth, intHeight int) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch resizewindowpixel %d %d,address:%s", intWidth, intHeight, client.Address))
}
