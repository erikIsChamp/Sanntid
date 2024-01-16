package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "time"
)

func main() {
    // The address of the server you want to connect to
    serverIP := "10.100.23.129"
    port := 33546 // use 34933 for fixed-size messages
    serverAddr := fmt.Sprintf("%s:%d", serverIP, port)

    // Connect to the server
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Disable the Nagle's Algorithm to prevent coalescing of small packets
    if tcpConn, ok := conn.(*net.TCPConn); ok {
        tcpConn.SetNoDelay(true)
    }

    // Create a reader to read the welcome message
    reader := bufio.NewReader(conn)

    // Read the welcome message
    welcomeMsg, err := reader.ReadString('\000') // read until '\0'
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Remove the null character at the end
    welcomeMsg = strings.TrimSuffix(welcomeMsg, "\000")

    fmt.Println("Received:", welcomeMsg)

    // The messages you want to send
    messages := []string{"Hello, Server!", "How are you?", "Goodbye, Server!"}

    for _, message := range messages {
        // Append a null character to the message
        message += "\000"

        // Send the message
        _, err = conn.Write([]byte(message))
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

        // Wait a bit before sending the next message
        time.Sleep(1 * time.Second)
    }
}