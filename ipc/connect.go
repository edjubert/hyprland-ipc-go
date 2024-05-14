package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// HyprlandInstanceSignature is the hyprland instance signature
const HyprlandInstanceSignature = "HYPRLAND_INSTANCE_SIGNATURE"

// XdgRuntimeDir is the environment variable that lead to directory for XDG runtime
const XdgRuntimeDir = "XDG_RUNTIME_DIR"

const MaxBufferReadSize = 4096

type Echo struct {
	Length int
	Data   []byte
}

func GetSignature() string {
	return fmt.Sprintf("%s/%s", os.Getenv(XdgRuntimeDir), os.Getenv(HyprlandInstanceSignature))
}

// StartUnixConnection returns a connection to a socket name under <XDG_RUNTIME_DIR>/<HYPRLAND_INSTANCE_SIGNATURE>/.
// socketname.sock
func StartUnixConnection(name string) net.Conn {
	connection, err := net.Dial("unix", fmt.Sprintf("%s/%s", GetSignature(), name))
	if err != nil {
		panic(err)
	}

	return connection
}

// Write pushes a message to an opened connection
func Write(c net.Conn, msg string) error {
	length := len(msg)
	data := make([]byte, 0, 4+length)

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(length))
	data = append(data, buf...)

	w := bytes.Buffer{}
	err := binary.Write(&w, binary.BigEndian, data)
	if err != nil {
		return err
	}

	data = append(data, w.Bytes()...)

	_, err = c.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Read gets message from an opened connection
func Read(c net.Conn) (string, error) {
	buf := make([]byte, MaxBufferReadSize)

	_, err := c.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

type HyprSocketMessage map[string][]string

func getSocketMessage(messages string) HyprSocketMessage {
	message := strings.Split(messages, "\n")

	socketMessages := HyprSocketMessage{}
	for _, msg := range message {
		m := strings.Split(msg, ">>")
		if len(m) != 2 {
			continue
		}

		socketMessages[m[0]] = append(socketMessages[m[0]], m[1])
	}

	return socketMessages
}

const MaxDialAttemp = 10

// ConnectHyprctl opens a connection to the writable socket
// You must defer close it after opening it otherwise it will freeze your Hyprland session
// conn, err := ConnectHyprctl()
//
//	if err != nil {
//	  //handle error
//	}
//
//	defer func() {
//	  _ = conn.Close()
//	}
func ConnectHyprctl(attempt int) (net.Conn, error) {
	signature := GetSignature()
	hyprctl := signature + "/.socket.sock"

	conn, err := net.Dial("unix", hyprctl)
	if err != nil {
		fmt.Println("[HYPRCTL] listen error", err)
		time.Sleep(time.Second / 2)
		fmt.Println("[HYPRCTL] - Retrying")
		if attempt < MaxDialAttemp {
			return ConnectHyprctl(attempt + 1)
		}

		return nil, err
	}

	return conn, nil
}

func closeConn(conn net.Conn) {
	if err := conn.Close(); err != nil {
		fmt.Printf("[ERROR] - Could not close connection -> %v", err)
	}
}

type HyprlandCallback func(socketMessages HyprSocketMessage)

// ConnectEvents opens a connection to the readable hyprland socket
func ConnectEvents(callbacks []HyprlandCallback) {
	signature := GetSignature()
	socket := signature + "/.socket2.sock"

	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer closeConn(conn)

	for {
		msg, err := Read(conn)
		if err != nil {
			log.Fatal(err)
		}

		socketMessages := getSocketMessage(msg)

		for _, callback := range callbacks {
			go callback(socketMessages)
		}
	}
}
