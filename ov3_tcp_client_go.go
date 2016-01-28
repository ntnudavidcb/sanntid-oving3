package main

import (
    "log"
    "net"
    "os"
    // "time"
)

func CheckError(err error) {
    if err  != nil {
        log.Println("Error: " , err)
        os.Exit(0)
    }
}

func ListenForConnections(done chan bool) {
    local, err := net.ResolveTCPAddr("tcp", ":20010")
    CheckError(err)

    listener, err := net.ListenTCP("tcp", local)
    CheckError(err)

    conn, err := listener.AcceptTCP()
    CheckError(err)
    buffer := make([]byte, 1024)
    conn.Read(buffer)
    log.Println(buffer)

    log.Println(conn.RemoteAddr(), "connected to arbeidsplass 10!")
    conn.Close()
    done <- true
}

func main() {
    conn, err := net.Dial("tcp", "129.241.187.23:33546")
    log.Println("test")
    CheckError(err)
    log.Println("test2")

    buffer := make([]byte, 1024)
    bytes_read, err := conn.Read(buffer)
    CheckError(err)

    log.Println("Read", bytes_read, "bytes:", string(buffer))

    done := make(chan bool)
    go ListenForConnections(done)

    bytes_sent, err := conn.Write([]byte("Connect to: 129.241.187.158:20010\x00"))
    conn.Write([]byte{0})
    CheckError(err)

    log.Println("Sent", bytes_sent, "bytes")

    <- done
    err = conn.Close()
    
    CheckError(err)
}