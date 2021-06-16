package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartServer() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// defer li.Close()

	log.Println("Start tcp server")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("連線中斷.")
}
