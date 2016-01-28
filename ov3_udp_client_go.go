// Go 1.2
// go run helloworld_go.go

package main
 
import (
    "net"
    "time"
    "log"
    // "fmt"
)

func ListenThing(conn *net.UDPConn, read_channel chan []byte) {
    for {
        data := make([]byte, 1024)
        log.Println("Attempt read")
        _, _, err := conn.ReadFromUDP(data)
        if err != nil {
            log.Fatal(err)
        }
        read_channel <- data

    }
}

func WriteThing(conn *net.UDPConn) {
    for {
        data := []byte("Testing 123")
        sent_bytes, err := conn.Write(data)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Sent", sent_bytes, "bytes")
        time.Sleep(1 * time.Second)
    }
}

func main() {
    // Listen for data through this address:port
    local, err :=  net.ResolveUDPAddr("udp4", ":20010")
    conn2, err := net.ListenUDP("udp4", local)
    if err != nil {
        log.Fatal(err)
    }

    // Send data to this address:port (broadcast address)
    remote, err := net.ResolveUDPAddr("udp4","129.241.187.23:20010")
    if err != nil {
        log.Fatal(err)
    }

    conn, err := net.DialUDP("udp4", nil, remote)
    if err != nil {
        log.Fatal(err)
    }

    read_channel := make(chan []byte)

    go ListenThing(conn2, read_channel)
    go WriteThing(conn)

    for {
        log.Println(<- read_channel)
    }
}