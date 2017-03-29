package server


import (
    "fmt"
    "net"
    "bufio"
    "encoding/json"
)


type Server struct {
    users map[string]User
}


func New() *Server {
    server := Server{ users:make(map[string]User) }
    return &server
}


func (s *Server) handleConnection(conn net.Conn) {
    reader:= bufio.NewReader(conn)

    for {
        message, error := reader.ReadString('\n')
        if error != nil {
            fmt.Println("ERROR")
            fmt.Println(error)
            continue
        }
        s.handleMessage(message, conn)
    } 
}


func (s *Server) handleMessage(message string, conn net.Conn) {
        values := make(map[string]string)
        e := json.Unmarshal([]byte(message), &values) 
        if e != nil {
            fmt.Println("Unmarshal error")
            fmt.Println(e)
            return
        }

        messageType := values["type"]
        fmt.Println(messageType)    
        switch messageType {
            case "LOGIN":
                s.handleLogin(values, conn)
            case "JOIN_LOBBY":
            case "SEND_LOBBY":
            case "SEND_USER":
                s.handleSendUser(values)


        } 


}

func (s *Server) handleLogin(values map[string]string, conn net.Conn) {
    userId := values["value"]
    user := User{userId, conn}
    s.users[values["value"]] = user
}

func (s *Server) handleSendUser(values map[string]string) {
    sendUserId := values["send_user_id"] 
    sendUserConn := s.users[sendUserId].conn
    message := values["message"]
    sendUserConn.Write([]byte(message))
}


func (s *Server) Listen() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go s.handleConnection(conn) 
    }      
}
