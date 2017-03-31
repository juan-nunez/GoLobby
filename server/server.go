package server


import (
    "fmt"
    "net"
    "bufio"
    "encoding/json"
    "math/rand"
    "strconv"
    "errors"
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
            user, _ := s.getUserByConn(conn)
            delete(s.users, user.name)
            defer conn.Close()
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
            case "MSG_USER":
                s.handleSendUser(values, conn)
        } 
}

func (s *Server) handleRegister(values map[string]string, conn net.Conn) {
    username := values["username"]
    if _, ok := s.users[username]; ok {
        message := "Username already exists"
        m := newRegisterErrorMessage(message)
        jsonified := messageToJson(m)
        conn.Write(jsonified)
        return
    }
    userId := strconv.Itoa(rand.Int())
    user := User{userId, username, conn}
    s.users[values["username"]] = user
}

func (s *Server) handleSendUser(values map[string]string, conn net.Conn) {
    sendUsername := values["to"] 
    me, error := s.getUserByConn(conn) 
    var myName string
    if error != nil {
       myName = "unknown" 
    } else {
        myName = me.name
    }
    sendUserConn := s.users[sendUsername].conn
    if sendUserConn == nil {
        return
    }
    message := values["message"]
    m := newSendMessage(myName, message)
    jsonified := messageToJson(m)
    sendUserConn.Write(jsonified)
}



func (s *Server) getUserByConn(conn net.Conn) (User, error) {
    users := s.users
    for _, user := range users {
        if user.conn == conn {
            return user, nil
        }
    }
    error := errors.New("No user found")
    return User{}, error
}

