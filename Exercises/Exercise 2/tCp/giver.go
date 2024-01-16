package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func main() {
    // The address of the server you want to connect to
    serverIP := "10.100.23.129"
    port := 33546 // use 34933 for fixed-size messages
    serverAddr := fmt.Sprintf("%s:%d", serverIP, port)

    // Your IP and the port you're listening on
    myIP := "10.100.23.17"
    myPort := 20007
    myAddr := fmt.Sprintf("%s:%d", myIP, myPort)

    // Set up a TCP listener on your side
    listener, err := net.Listen("tcp", myAddr)
    if err != nil {
        fmt.Println("Error setting up listener:", err)
        return
    }
    defer listener.Close()

    // Connect to the server
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Send a message to the server telling it to connect back to you
    message := fmt.Sprintf("Connect to: %s\000", myAddr)
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    fmt.Println("Message sent:", message)

    // Wait for the server to connect back to you
    serverConn, err := listener.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err)
        return
    }
    defer serverConn.Close()

    fmt.Println("Server connected!")

    // Now you can send messages and receive echoes on serverConn
    reader := bufio.NewReader(serverConn)
    message = "Hello, Server!\000"
    _, err = serverConn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    fmt.Println("Message sent:", message)

    // Read the server's response
    response, err := reader.ReadString('\000') // read until '\0'
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Remove the null character at the end
    response = strings.TrimSuffix(response, "\000")

    // Print the server's response
    fmt.Println("Received:", response)
}