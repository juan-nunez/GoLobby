import socket
import json

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.connect(("localhost", 8080))

value = raw_input("id: ")

d = {"type": "LOGIN", "value": value}
encoded = json.dumps(d)
sock.send(encoded + "\n")
user_id = raw_input("userid: ")
while True:
    message = raw_input("Message:")
    d = {"type" : "SEND_USER", "message": message, "send_user_id" : user_id} 
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    data = sock.recv(1024)
    print data
    
  
    





