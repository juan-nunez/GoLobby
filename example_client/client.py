import socket
import json

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.connect(("localhost", 8080))

name = raw_input("name: ")
d = {"type": "LOGIN", "username": name}
encoded = json.dumps(d)
sock.send(encoded + "\n")
username = raw_input("username: ")
while True:
    message = raw_input("Message:")
    d = {"type" : "SEND_USER", "message": message, "to" : username, "from": name} 
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    data = sock.recv(1024)
    print data
    
  
    





