package server


import (
    "fmt"
    "net"
    "bufio"
)


type Server struct {

}


func New() *Server {
    return &Server{}
}


func (s *Server) handleConnection(conn net.Conn) {
    for {
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println(string(message))
        conn.Write([]byte("received\n"))
    } 
}


func handleMessage(message string, conn net.Conn) {
       

}


func (s *Server) Listen() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go s.handleConnection(conn) 
    }      
}
