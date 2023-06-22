package IPC

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// HyprlandInstanceSignature is the hyprland instance signature
const HyprlandInstanceSignature = "HYPRLAND_INSTANCE_SIGNATURE"

type Echo struct {
	Length int
	Data   []byte
}

func GetSignature() string {
	return os.Getenv(HyprlandInstanceSignature)
}

// StartUnixConnection returns a connection to a socket name under /tmp/hypr/<HYPRLAND_INSTANCE_SIGNATURE>/.socketname.sock
func StartUnixConnection(name string) net.Conn {
	connection, err := net.Dial("unix", fmt.Sprintf("/tmp/hypr/%s/%s", GetSignature(), name))
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
	buf := make([]byte, 4)

	_, err := c.Read(buf)
	if err != nil {
		return "", err
	}

	byteCount := binary.BigEndian.Uint32(buf)
	length := int(byteCount)
	data := make([]byte, length)

	_, err = c.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func handleMessage(msg string) {
	m := strings.Split(msg, ">>")
	if len(m) < 1 {
		fmt.Println("[WARN] - Not enough args")
		return
	}
}

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
func ConnectHyprctl() (net.Conn, error) {
	signature := GetSignature()
	hyprctl := "/tmp/hypr/" + signature + "/.socket.sock"

	conn, err := net.Dial("unix", hyprctl)
	if err != nil {
		log.Fatal("[HYPRCTL] listen error", err)
		return nil, err
	}

	return conn, nil
}

func closeConn(conn net.Conn) {
	if err := conn.Close(); err != nil {
		fmt.Printf("[ERROR] - Could not close connection -> %v", err)
	}
}

// ConnectEvents opens a connection to the readable hyprland socket
func ConnectEvents() {
	signature := GetSignature()
	socket := "/tmp/hypr/" + signature + "/.socket2.sock"

	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer closeConn(conn)

	for {
		_, err := Read(conn)
		if err != nil {
			log.Fatal(err)
		}

		//handleMessage(msg)
	}
}
