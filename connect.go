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

func StartUnixConnection() net.Conn {
	connection, err := net.Dial("unix", "/tmp/hypr/"+GetSignature()+"/.gophrland.sock")
	if err != nil {
		panic(err)
	}

	return connection
}

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
