package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleRequest(c net.Conn) {
	scanner := bufio.NewScanner(c)
	defer c.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		tokens := strings.Split(line, ":")
		if tokens[0] == "Host" {
			c.Write([]byte(tokens[1]))
		}

		fmt.Println(scanner.Text())
	}

}

func main() {
	server, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
