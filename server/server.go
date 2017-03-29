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
            conn.Close() 
            break 
        }
        s.handleMessage(message, conn)
    } 
}


func (s *Server) handleMessage(message string, conn net.Conn) {
        values := make(map[string]string)
        e := json.Unmarshal([]byte(message), &values) 
        if e != nil {
            fmt.Println("Unmarshal error")
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
    username := values["username"]
    userId := "5"
    user := User{userId, username, conn}
    s.users[values["username"]] = user
}

func (s *Server) handleSendUser(values map[string]string) {
    sendUsername := values["to"] 
    me := values["from"]
    sendUserConn := s.users[sendUsername].conn
    message := values["message"]
    jsonifiedMessage := s.formatSendMessage(me, message)
    sendUserConn.Write(jsonifiedMessage)
}


func (s *Server) formatSendMessage(username, message string) []byte {
    formattedMessage := make(map[string]string)
    formattedMessage["from"] = username
    formattedMessage["message"] = message
    jsonifiedMessage, _ := json.Marshal(formattedMessage)
    return jsonifiedMessage
}

func (s *Server) Listen() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go s.handleConnection(conn) 
    }      
}
