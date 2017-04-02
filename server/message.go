package server

import (
    "encoding/json"
)
type Message interface {
}


type SingleMessage struct {
    Data map[string]string `json:"data"`   
    Error map[string]string `json:"error"`
}


type ListMessage struct {
    Data map[string][]string `json:"data"`   
    Error map[string]string `json:"error"`
}


func newSendMessage(username, message string) Message {
    m := SingleMessage{Data:make(map[string]string), Error:make(map[string]string)} 
    m.Data["from"] = username
    m.Data["message"] = message
    return m
}


func newRegisterErrorMessage(message string) Message{
    m := SingleMessage{Data:make(map[string]string), Error:make(map[string]string)} 
    m.Error["error"] = message 
    return m 
}


func newUserListMessage(usernames []string) Message{
    m := ListMessage{Data:make(map[string][]string), Error:make(map[string]string)}
    m.Data["users"] = usernames
    return m
}


func messageToJson(m Message) []byte { 
    jsonified, _:= json.Marshal(m)
    return jsonified 
}




