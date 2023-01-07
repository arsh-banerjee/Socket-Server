/*****************************************************************************
 * client.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const SEND_BUFFER_SIZE = 2048

/* TODO: client()
 * Open socket and send message from stdin.
 */
func client(server_ip string, server_port string) {
	conn, err := net.Dial("tcp", server_ip+":"+server_port)
	if err != nil {
		log.Fatalf("Failed to connect to server - %v\n", err)
	}

	buffer := make([]byte, SEND_BUFFER_SIZE)
	reader := bufio.NewReader(os.Stdin)

	for {
		n, err := reader.Read(buffer)
		fmt.Println(string(buffer[:n]))
		_, _ = conn.Write(buffer[0:n])
		if err == io.EOF {
			break
		}
	}

	conn.Close()
}

// Main parses command-line arguments and calls client function
func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./client [server IP] [server port] < [message file]")
	}
	server_ip := os.Args[1]
	server_port := os.Args[2]
	client(server_ip, server_port)
}
