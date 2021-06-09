package internal

import (
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1"
	port = 1203
)

func HandleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from: " + remoteAddr)

	buf := make([]byte, 1024)
	for {
		reqLen, err := conn.Read(buf)
		if err != nil {

			if err.Error() == "EOF" {
				fmt.Println("Disconned from ", remoteAddr)
				break
			} else {
				fmt.Println("Error reading:", err.Error())
				break
			}
		}
		conn.Write([]byte("Message received.\n"))

		fmt.Printf("len: %d, recv: %s\n", reqLen, string(buf[:reqLen]))
	}
	conn.Close()
}

func StartServer ()
{
	src := addr + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()
	fmt.Printf("TCP server start and listening on %s.\n", src)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go HandleConnection(conn)
	}
}