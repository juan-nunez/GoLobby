package server


import (
    "fmt"
    "net"
    "bufio"
    "encoding/json"
    "math/rand"
    "strconv"
)


type Server struct {
    users map[string]User
}


func New() *Server {
    server := Server{ users:make(map[string]User) }
    return &server
}


func (s *Server) Listen() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go s.handleConnection(conn) 
    }      
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
            case "REGISTER":
                s.handleRegister(values, conn)
            case "SEND_USER":
                s.handleSendUser(values, conn)
        } 
}

func (s *Server) handleRegister(values map[string]string, conn net.Conn) {
    username := values["username"]
    userId := strconv.Itoa(rand.Int())
    user := User{userId, username, conn}
    s.users[values["username"]] = user
}

func (s *Server) handleSendUser(values map[string]string, conn net.Conn) {
    sendUsername := values["to"] 
    me := s.getNameByConn(conn) 
    sendUserConn := s.users[sendUsername].conn
    if sendUserConn == nil {
        return
    }
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

func (s *Server) getNameByConn(conn net.Conn) string {
    users := s.users
    for _, user := range users {
        if user.conn == conn {
            return user.name
        }
    }
    return "unknown"
}

