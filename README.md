# GoLobby
GoLobby is a simple server side communications system. Clients can connect and talk to each other using JSON encoded messages. This project is a work in progress. I am using it to learn Golang.


GoLobby is only functional with TCP sockets.


#### Message types


###### Register
```
{"type": "REGISTER", "username": username}
```

###### Message user
```
{"type": "MSG_USER", "to": toUsername, "message": message}
```

###### Broadcast message to lobby
```
{"type": "MSG_ALL", "message": message}
```

###### View connected users
```
{"type": "USER_LIST"}
```
