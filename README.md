# Hyprland IPC Go wrapper
This is a simple IPC module to interract with Hyprland in your Go projects.
The purpose of this project is to be used in my main project [Gophrland](https://github.com/edjubert/gophrland)

As it is to be used with Gophrland, I haven't implemented all functions yet.

Feel free to [open a PR](#how-to-open-a-pr) to add some functions

## Install
```bash
go get github.com/edjubert/hyprland-ipc-go
```

## How to use
This is a collection of functions and types

### Types
#### HyprlandWorkspace
```go
type HyprlandWorkspace struct {
    Id              int    `json:"id"`
    Windows         int    `json:"windows"`
    Monitor         string `json:"monitor"`
    Name            string `json:"name"`
    HasFullscreen   bool   `json:"hasfullscreen"`
    LastWindow      string `json:"lastwindow"`
    LastWindowTitle string `json:"lastwindowtitle"`
}
```

#### HyprlandMonitor
```go
type HyprlandMonitor struct {
    Id              int               `json:"id"`
    Name            string            `json:"name"`
    Description     string            `json:"description"`
    Make            string            `json:"make"`
    Model           string            `json:"model"`
    Serial          string            `json:"serial"`
    Width           int               `json:"width"`
    Height          int               `json:"height"`
    RefreshRate     float64           `json:"refreshRate"`
    X               int               `json:"x"`
    Y               int               `json:"y"`
    ActiveWorkspace HyprlandWorkspace `json:"activeWorkspace"`
    Reserved        []int             `json:"reserved"`
    Scale           float64           `json:"scale"`
    Transform       int               `json:"transform"`
    Focused         bool              `json:"focused"`
    DpmsStatus      bool              `json:"dpmsStatus"`
    Vrr             bool              `json:"vrr"`
}
```

#### HyprlandClient
```go
type HyprlandClient struct {
    Address        string            `json:"address,omitempty"`
    Mapped         bool              `json:"mapped,omitempty"`
    Hidden         bool              `json:"hidden,omitempty"`
    At             []int             `json:"at,omitempty"`
    Size           []int             `json:"size,omitempty"`
    Workspace      HyprlandWorkspace `json:"workspace,omitempty"`
    Floating       bool              `json:"floating,omitempty"`
    Monitor        int               `json:"monitor,omitempty"`
    Class          string            `json:"class,omitempty"`
    InitialClass   string            `json:"initialClass,omitempty"`
    Title          string            `json:"title,omitempty"`
    InitialTitle   string            `json:"initialTitle,omitempty"`
    Pid            int               `json:"pid,omitempty"`
    XWayland       bool              `json:"xwayland,omitempty"`
    Pinned         bool              `json:"pinned,omitempty"`
    Fullscreen     bool              `json:"fullscreen,omitempty"`
    FullscreenMode int               `json:"fullscreenMode,omitempty"`
    FakeFullscreen bool              `json:"fakeFullscreen,omitempty"`
}
```

### Dispatchers
All dispatchers are using [Hyprland IPC](https://wiki.hyprland.org/IPC/) so all calls are asynchronously executed
This prevents you to block your Hyprland session if you are running a heavy command

#### Current implementation
- Move
  - `WindowPixelExact(x int, y int, clientAddress string) error`
  - `ToWorkspaceName(workspaceName string, clientAddress string) error`
  - `ToWorkspaceSilen(workspaceName string, clientAddress string) error`
  - `ToSpecialNamed(specialWorkspaceName string, clientAddress string) error`
  - `ClientToCurrent(clientAddress string) error`
  - `CenterFloatingClient(client HyprlandClient, monitor HyprlandMonitor, applyRand bool) error`
- Focus
  - `Window(clientAddress string) error`
  - `Monitor(monitor HyprlandMonitor) error`
  - `WorkspaceID(workspaceId int) error`
- Toggle
  - `Floating(clientAddress string) error`
  - `SpecialWorkspace(specialWorkspaceName string) error`
-  Resize
    - `WindowExactPixel(client HyprlandClient,width int, height int) error`

#### Example

```go
package main

import (
  "github.com/edjubert/hyprland-ipc-go/hyprctl/dispatch"
)

func main() {
  move := dispatch.Move{}
  _ = move.ToWorkspaceName("my workspace name", "0x000000")
}
```

### Getters
- Get
  - `ActiveClient() (HyprlandClient, error)`
  - `WorkspaceFloatingClients(workspace HyprlandWorkspace) ([]HyprlandClient, error)`
  - `Clients() ([]HyprlandClient, error)`
  - `ClientByPID(clients []HyprlandClient, pid int) (HyprlandClient, error)`
  - `ClientByClassName(clients []HyprlandClient, class string) (HyprlandClient, error)`
  - `Monitors(format string) ([]HyprlandMonitor, error)`
  - `ActiveMonitor(monitors []HyprlandMonitor) (HyprlandMonitor, error)`
  - `ActiveWorkspace() (HyprlandWorkspace, error)`
  - `Workspaces() ([]HyprlandWorkspace, error)`
  - `MonitorByID(monitorId int) (HyprlandMonitor, error)`

#### Example

```go
package main

import (
	"github.com/edjubert/hyprctl-ipc/get"
)

func main() {
	getter := get.Get{}
    if _, err := getter.ActiveClient(); err != nil {
		// Do something
    }
}
```

### Send Hyprland notification
Like [dispatchers](#dispatchers), `SendNotification` uses [Hyprland IPC](https://wiki.hyprland.org/IPC/)
- `SendNotification(time int, msgType string, msg string) error`

```go
package main

import "github.com/edjubert/hyprland-ipc-go/hyprctl/notify"

func main() {
	_ = notify.SendNotification(2000, "info", "Hello hyprland-ipc-go")
}
```

## How to open a PR
To open a PR, fork the GitHub project, work on your additions and go to:
- Pull Request -> New pull request
- compare across forks
- Select your fork and branch