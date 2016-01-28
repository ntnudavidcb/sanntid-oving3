// Go 1.2
// go run helloworld_go.go
// my ip: 129.241.187.148

package main
 
import (
    "fmt"
    "net"
    "os"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 30000*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":30000")
    CheckError(err)
 
    fmt.Println("ServerAddr: ", ServerAddr)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}