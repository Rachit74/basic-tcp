package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	/*
		net.Listen() takes 2 args, type of conn and port number
	*/
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close() // defer used to close connection

	fmt.Println("Server is listening on port 8080")
	/*
		A defer statement defers the execution of a function until the surrounding function returns.
		defer ln.Close() will execute after everything else in the main function
		closing and cleaning after all communication
	*/

	// Accepting client connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
		}

		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New Client connection", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("Client says:", msg)

		// 5. Respond back
		conn.Write([]byte("Echo: " + msg + "\n"))
	}

	// // Read and process data from the client
	// scanner := bufio.NewScanner(conn)

	// // Write data back to the client
	// // ...
}
