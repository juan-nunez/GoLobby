package main


import (
    "GoLobby/server"
)


func main() {
    s := server.New() 
    s.Listen()

}
