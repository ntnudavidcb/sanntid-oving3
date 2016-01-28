package main

import (
    "log"
    "net"
    // "time"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func ListenForConnections(done chan bool) {
    local, err := net.ResolveTCPAddr("tcp", ":12345")
    CheckError(err)

    listener, err := net.ListenTCP("tcp", local)
    CheckError(err)

    conn, err := listener.AcceptTCP()
    CheckError(err)

    log.Println(conn.RemoteAddr(), "connected to arbeidsplass 12!")
    conn.Close()
    done <- true
}

func main() {
    remote, err := net.ResolveTCPAddr("tcp", "129.241.187.136:33546")
    CheckError(err)

    conn, err := net.DialTCP("tcp", nil, remote)
    CheckError(err)

    buffer := make([]byte, 1024)
    bytes_read, err := conn.Read(buffer)
    CheckError(err)

    log.Println("Read", bytes_read, "bytes:", string(buffer))

    done := make(chan bool)
    go ListenForConnections(done)

    // Fixed size message sending: Use port 34933 in remote
    // msg := "Connect to: 129.241.187.144:12345"
    // data := make([]byte, 1024)
    // copy(data[:], msg)
    // bytes_sent, err := conn.Write(data)

    // Variable size messages, use port 33546 in remote
    bytes_sent, err := conn.Write([]byte("Connect to: 129.241.187.144:12345"))
    conn.Write([]byte{0})
    CheckError(err)

    log.Println("Sent", bytes_sent, "bytes")

    <- done
    err = conn.Close()
    
    CheckError(err)
}