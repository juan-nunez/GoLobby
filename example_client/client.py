import socket
import json

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.connect(("localhost", 8080))


d = {"type": "LOGIN", "value": "75"}

encoded = json.dumps(d)
print encoded
sock.send(encoded + "\n")


