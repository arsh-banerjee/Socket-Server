/*****************************************************************************
 * server.go
 * Name: Arsh Banerjee
 * NetId: arshb
 *****************************************************************************/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Open socket and wait for client to connect
 * Print received message to stdout
 */
func server(server_port string) {
	ln, err := net.Listen("tcp", ":"+server_port)
	if err != nil {
		log.Fatalf("Failed to setup a listener - %v\n", err)
	}
	defer ln.Close()
	buffer := make([]byte, RECV_BUFFER_SIZE)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection - %v\n", err)
		}
		defer conn.Close()

		for {
			message, err := conn.Read(buffer)
			fmt.Print(string(buffer[:message]))
			if err != nil {
				break
			}
		}
	}

}

// Main parses command-line arguments and calls server function
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server [server port]")
	}
	server_port := os.Args[1]
	server(server_port)
}
