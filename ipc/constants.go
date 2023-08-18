package ipc

const (
	// Workspace
	// emitted on workspace change.
	// Is emitted ONLY when a user requests a workspace change, and is not emitted on mouse movements (see activemon)
	//
	// DATA:
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	Workspace = "workspace"

	// FocusedMonitor
	// emitted on the active monitor being changed.
	//
	// DATA:
	// - MONNAME (HyprlandMonitor.Name)
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	FocusedMonitor = "focusedmon"

	// ActiveWindow
	// emitted on the active window being changed.
	//
	// DATA:
	// - WINDOWCLASS (HyprlandClient.Class)
	// - WINDOWTITLE (HyprlandClient.Title)
	ActiveWindow = "activewindow"

	// ActiveWindowV2
	// emitted on the active window being changed.
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	ActiveWindowV2 = "activewindowv2"

	// Fullscreen
	// emitted when a fullscreen status of a window changes.
	//
	// DATA:
	// - 0 (Exit fullscreen)
	// - 1 (Enter fullscreen)
	Fullscreen = "fullscreen"

	// MonitorRemoved
	// emitted when a monitor is removed (disconnected)
	//
	// DATA:
	// - MONITORNAME (HyprlandMonitor.Name)
	MonitorRemoved = "monitorremoved"

	// MonitorAdded
	// emitted when a monitor is added (connected)
	//
	// DATA:
	// - MONITORNAME (HyprlandMonitor.Name)
	MonitorAdded = "monitoradded"

	// CreateWorkspace
	// 	emitted when a workspace is created
	//
	// DATA:
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	CreateWorkspace = "createworkspace"

	// DestroyWorkspace
	// emitted when a workspace is destroyed
	//
	// DATA:
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	DestroyWorkspace = "destroyworkspace"

	// MoveWorkspace
	// emitted when a workspace is moved to a different monitor
	//
	// DATA:
	// - WORKSPACENAME (HyprlandMonitor.Name)
	// - MONNAME (HyprlandMonitor.Name)
	MoveWorkspace = "moveworkspace"

	// ActiveLayout
	// emitted on a layout change of the active keyboard
	//
	// DATA:
	// - KEYBOARDNAME string (not mapped yet)
	// - LAYOUTNAME string (not mapped yet)
	ActiveLayout = "activelayout"

	// OpenWindow
	// emitted when a window is opened
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	// - WINDOWCLASS (HyprlandClient.Class)
	// - WINDOWTITLE (HyprlandClient.Title)
	OpenWindow = "openwindow"

	// CloseWindow
	// emitted when a window is closed
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	CloseWindow = "closewindow"

	// MoveWindow
	// emitted when a window is moved to a workspace
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	// - WORKSPACENAME (HyprlandWorkspace.Name)
	MoveWindow = "movewindow"

	// OpenLayer
	// emitted when a layerSurface is mapped
	//
	// DATA:
	// - NAMESPACE string (not mapped yet)
	OpenLayer = "openlayer"

	// CloseLayer
	// emitted when a layerSurface is unmapped
	//
	// DATA:
	// - NAMESPACE string (not mapped yet)
	CloseLayer = "closelayer"

	// SubMap
	// emitted when a keybind submap changes. Empty means default.
	//
	// DATA:
	// - SUBMAPNAME string (not mapped yet)
	SubMap = "submap"

	// ChangeFloatingMode
	// emitted when a window changes its floating mode. FLOATING is either 0 or 1.
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	// - FLOATING int (0 -> no | 1 -> yes)
	ChangeFloatingMode = "changefloatingmode"

	// Urgent
	// emitted when a window requests an urgent state
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	Urgent = "urgent"

	// Minimize
	// emitted when a window requests a change to its minimized state. MINIMIZED is either 0 or 1.
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	// - MINIMIZED (0 -> no | 1 -> yes)
	Minimize = "minimize"

	// Screencast
	// emitted when a screencopy state of a client changes.
	// Keep in mind there might be multiple separate clients.
	// - STATE:
	//   - 0 -> not active
	//   - 1 -> active
	// - OWNER:
	//   - 0 -> monitor share
	//   - 1 -> window share
	//
	// DATA:
	// - STATE int (not mapped yet)
	// - OWNER int (not mapped yet)
	Screencast = "screencast"

	// WindowTitle
	// emitted when a window title changes.
	//
	// DATA:
	// - WINDOWADDRESS (HyprlandClient.Address)
	WindowTitle = "windowtitle"
)
