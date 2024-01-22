package dispatch

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
