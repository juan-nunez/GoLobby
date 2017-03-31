package server

import (
    "encoding/json"
)

type Message struct {
    Data map[string]string `json:"data"`   
    Error map[string]string `json:"error"`
}


func newSendMessage(username, message string) Message {
    m := Message{Data:make(map[string]string), Error:make(map[string]string)} 
    m.Data["from"] = username
    m.Data["message"] = message
    return m
}


func newRegisterErrorMessage(message string) Message{
    m := Message{Data:make(map[string]string), Error:make(map[string]string)} 
    m.Error["error"] = message 
    return m 
}


func messageToJson(m Message) []byte { 
    jsonified, _:= json.Marshal(m)
    return jsonified 
}




