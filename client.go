package main

import (
	"bufio"
	"fmt"
	"net"
)

func RunClient() {
	fmt.Println("Start Client...")
	ln, _ := net.Listen("tcp", ":8000")
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		fmt.Fprintf(conn, "Message Received at subscriptionID, MessageID"+"\n")
	}
}
