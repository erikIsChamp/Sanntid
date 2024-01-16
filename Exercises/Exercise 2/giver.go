package main

import (
    "fmt"
    "net"
    
)

func main() {
    // The address of the server you want to send to
    serverIP := "10.100.23.129"
    workspaceNumber := 7 // replace with your workspace number
    port := 20000 + workspaceNumber
    serverAddr := fmt.Sprintf("%s:%d", serverIP, port)

    // The message you want to send
    message := []byte("Hello, Server!")

    // Resolve the server address
    addr, err := net.ResolveUDPAddr("udp", serverAddr)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Create a UDP connection
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    // Send the message
    _, err = conn.Write(message)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Message sent!")

	

}