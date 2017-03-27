package server


import (
    "fmt"
    "net"
    "bufio"
)


type server struct {

}


func New() *server {
    return &server{}
}


func (s *server) handleConnection(conn net.Conn) {
    for {
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println(string(message))
        conn.Write([]byte("received\n"))
    } 
}



func (s *server) Listen() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go s.handleConnection(conn) 
    }      
}
