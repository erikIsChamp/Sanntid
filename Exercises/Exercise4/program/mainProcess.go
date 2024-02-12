package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

var filename string = "/Users/larsleopold/Desktop/exercise/Sanntid/Exercises/Exercise4/program/mainProcess"

var Addr string = "localhost:6969"

func backmyshiup() {
	fmt.Println("initilaizing backup")
	udpAddr, err := net.ResolveUDPAddr("udp", Addr)
    if err != nil {
        fmt.Println("Error resolving UDP address:", err)
        return
    }
    conn, err := net.ListenUDP("udp", udpAddr)
    if err != nil {
        fmt.Println("Error listening on UDP:", err)
        return
    }
	buffer := make([]byte, 1024)
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		_, _, err := conn.ReadFromUDP(buffer)
		i := buffer[0]
		if err != nil {
			go process(i)
			break
		}
	}
}


func process(Count byte) {
	fmt.Println("I am the captain now")
	exec.Command("osascript", "-e", `tell app "Terminal" to do script "go run `+filename+`.go"`).Run()
	udpAddr, _ := net.ResolveUDPAddr("udp", Addr)
	conn, _ := net.DialUDP("udp",nil, udpAddr)
	defer conn.Close()

	for {
		time.Sleep(1 * time.Second)
		Count++
		fmt.Println("Count:", Count)
		conn.Write([]byte{Count}) //write to server
	}
}

func main() {

	go backmyshiup()

	select {}

}
