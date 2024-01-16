package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    // The address of the server you want to send to
    serverIP := "10.100.23.129"
    port := "20007"
    serverAddr := serverIP + ":" + port

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

    // Buffer to store the server's response
    buffer := make([]byte, 1024)

    // Read the server's response
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the server's response
    fmt.Println("Received: ", string(buffer[:n]))

    // Sleep for a while before sending the next message
    time.Sleep(2 * time.Second)
}