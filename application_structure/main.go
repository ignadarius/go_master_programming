package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleRequest(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	server, err := net.Listen("tcp", ";8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
