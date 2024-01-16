package main

import (
    "fmt"
    "net"
)


func main() {
    // Listen on UDP port 30000 on all available unicast and
    // anycast IP addresses of the local system.
    addr, err := net.ResolveUDPAddr("udp", ":30000")
    if err != nil {
        fmt.Println(err)
        return
    }
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    buffer := make([]byte, 1024)

    for {
        // Read from UDP connection.
        // This will block until client send data.
        n, fromWho, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Print the IP address of the sender
        fmt.Println("Received ", string(buffer[:n]), " from ", fromWho.IP)

        // Optional: filter out messages from ourselves
        // if fromWho.IP.String() != localIP {
        //     // do stuff with buffer
        // }
    }
}
