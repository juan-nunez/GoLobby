import socket
import json
import threading

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.connect(("localhost", 8080))

name = raw_input("Your name: ")
d = {"type": "REGISTER", "username": name}
encoded = json.dumps(d)
sock.send(encoded + "\n")
username = raw_input("username of receiver: ")

 
def receiver():
    while True:
        data = sock.recv(1024)
        print data


def sender():
    while True:
        message = raw_input("Message:")
        d = {"type" : "SEND_USER", "message": message, "to" : username} 
        encoded = json.dumps(d)
        sock.send(encoded + "\n")




threading.Thread(target=receiver).start()
threading.Thread(target=sender).start()  

