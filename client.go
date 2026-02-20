package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	message := "Hello from the Client!"
	fmt.Println("Sending message to server...")
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	reply := make([]byte, 1024)
	n, _ := conn.Read(reply)
	fmt.Printf("Server replied: %s\n", string(reply[:n]))
}
