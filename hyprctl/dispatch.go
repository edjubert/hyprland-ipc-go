package hyprctl

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/ipc"
	"github.com/edjubert/hyprland-ipc-go/types"
	"math/rand"
)

type Dispatch struct {
	Move   Move
	Focus  Focus
	Toggle Toggle
}

const (
	// DispatchKey is the dispatch keyword
	DispatchKey = "dispatch"
	// Exec
	Exec = "exec"
	// ExecR
	ExecR = "execr"
	// Pass
	Pass = "pass"
	// KillActive
	KillActive = "killactive"
	// CloseWindow
	CloseWindow = "closeWindow"
	// Workspace changes the workspace
	Workspace = "workspace"
	// MoveToWorkspace
	MoveToWorkspace = "movetoworkspace"
	// MoveToWorkspaceSilent
	MoveToWorkspaceSilent = "movetoworkspacesilent"
	// ToggleFloating
	ToggleFloating = "togglefloating"
	// Fullscreen
	Fullscreen = "fullscreen"
	// FakeFullscreen
	FakeFullscreen = "fakefullscreen"
	// Dpms
	Dpms = "dpms"
	// Pin
	Pin = "pin"
	// MoveFocus
	MoveFocus = "movefocus"
	// MoveWindow
	MoveWindow = "movewindow"
	// SwapWindow
	SwapWindow = "swapwindow"
	// CenterWindow
	CenterWindow = "centerwindow"
	// ResizeActive
	ResizeActive = "resizeactive"
	// MoveActive
	MoveActive = "moveactive"
	// ResizeWindowPixel
	ResizeWindowPixel = "resizewindowpixel"
	// ResizeWindowPixelExact
	ResizeWindowPixelExact = "resizewindowpixel exact"
	// MoveWindowPixel moves a selected window
	MoveWindowPixel = "movewindowpixel"
	// MoveWindowPixelExact
	MoveWindowPixelExact = "movewindowpixel exact"
	// CycleNext
	CycleNext = "cyclenext"
	// SwapNext
	SwapNext = "swapnext"
	// FocusWindow
	FocusWindow = "focuswindow"
	// FocusMonitor
	FocusMonitor = "focusmonitor"
	// SplitRatio
	SplitRatio = "splitratio"
	// ToggleOpaque
	ToggleOpaque = "toggleopaque"
	// MoveCursorToCorner
	MoveCursorToCorner = "movecursortocorner"
	// MoveCursor
	MoveCursor = "movecursor"
	// WorkspaceOpt
	WorkspaceOpt = "workspaceopt"
	// RenameWorkspace
	RenameWorkspace = "renameworkspace"
	// Exit
	Exit = "exit"
	// ForceRendererReload
	ForceRendererReload = "forcerendererreload"
	// MoveCurrentWorkspaceToMonitor
	MoveCurrentWorkspaceToMonitor = "movecurrentworkspacetomonitor"
	// MoveWorkspaceToMonitor
	MoveWorkspaceToMonitor = "moveworkspacetomonitor"
	// SwapActiveWorkspaces
	SwapActiveWorkspaces = "swapactiveworkspaces"
	// BringActiveToTop
	BringActiveToTop = "bringactivetotop"
	// ToggleSpecialWorskpace
	ToggleSpecialWorkspace = "togglespecialworkspace"
	// FocusUrgentOrLast
	FocusUrgentOrLast = "focusurgentorlast"
	// LockGroups
	LockGroups = "lockgroups"
	// LockActiveGroup
	LockActiveGroup = "lockactivegroup"
	// MoveIntoGroup
	MoveIntoGroup = "moveintogroup"
	// MoveOutOfGroup
	MoveOutOfGroup = "moveoutofgroup"
	// MoveGroupWindow
	MoveGroupWindow = "movegroupwindow"
	// Global
	Global = "global"
	// Submap
	Submap = "submap"
)

func runHyprctlSocket(cmd string) error {
	conn, err := ipc.ConnectHyprctl(0)
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

// FocusWorkspaceID focuses a HyprlandWorkspace.Id
func (d *Dispatch) FocusWorkspaceID(ID int) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %d", DispatchKey, Workspace, ID))
}

type Toggle struct{}

// Floating activate/deactivate the HyprlandClient.Floating mode for the given HyprlandClient.Address
func (t *Toggle) Floating(address string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s address:%s", DispatchKey, ToggleFloating, address))
}

// SpecialWorkspace show/hide the special HyprlandWorkspace.Name
func (t *Toggle) SpecialWorkspace(name string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %s", DispatchKey, ToggleSpecialWorkspace, name))
}

type Focus struct{}

// Window focuses a given HyprlandClient.Address
func (f *Focus) Window(address string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s address:%s", DispatchKey, FocusWindow, address))
}

// Monitor focuses a given HyprlandMonitor
func (f *Focus) Monitor(monitor types.HyprlandMonitor) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %s", DispatchKey, FocusMonitor, monitor.Name))
}

type Move struct{}

// WindowPixelExact move at precise HyprlandClient.At the HyprlandClient.Address
func (m *Move) WindowPixelExact(x, y int, address string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, MoveWindowPixelExact, x, y, address))
}

// ToWorkspaceName moves a given HyprlandClient.Address to a HyprlandWorkspace.Name and focus the HyprlandClient
func (m *Move) ToWorkspaceName(workspaceName, clientAddress string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %s,address:%s", DispatchKey, MoveToWorkspace, workspaceName, clientAddress))
}

// ToWorkspaceSilent moves a given HyprlandClient.Address to a HyprlandWorkspace.Name without focussing the HyprlandClient
func (m *Move) ToWorkspaceSilent(name, address string) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %s,address:%s", DispatchKey, MoveToWorkspaceSilent, name, address))
}

// ToSpecialNamed moves given HyprlandClient.Address to named special HyprlandWorkspace
func (m *Move) ToSpecialNamed(specialWorkspaceName, clientAddress string) error {
	if specialWorkspaceName != "" {
		specialWorkspaceName = fmt.Sprintf(":%s", specialWorkspaceName)
	}
	return m.ToWorkspaceSilent("special"+specialWorkspaceName, clientAddress)
}

// ResizeWindowExactPixel resize given HyprlandClient to specific width and height
func (d *Dispatch) ResizeWindowExactPixel(client types.HyprlandClient, intWidth, intHeight int) error {
	return runHyprctlSocket(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, ResizeWindowPixelExact, intWidth, intHeight, client.Address))
}

// ClientToCurrent moves a given HyprlandClient.Address to current HyprlandWorkspace
func (m *Move) ClientToCurrent(address string) error {
	getter := Get{}
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

	return runHyprctlSocket(fmt.Sprintf("%s %s %d %d,address:%s", DispatchKey, MoveWindowPixelExact, centerX, centerY, client.Address))
}
