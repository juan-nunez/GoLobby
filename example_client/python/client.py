import socket
import json
import threading
import sys


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
receiverThread = None

def start():
    print "(1) Register"
    print "(2) Exit"
    option = raw_input()
    

    if option == "1":
        register()
    elif option == "2":
        sys.exit()

def register():
    global sock
    global receiverThread
    sock.connect(("localhost", 8080))
    receiverThread = threading.Thread(target=receiver)
    receiverThread.start()
    name = raw_input("Your name: ")
    d = {"type": "REGISTER", "username": name}
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    menu()



def menu():
    print "(1) See User List"
    print "(2) Broadcast Message"
    print "(3) Send private Message"
    option = raw_input()
   
    if option == "1":
        seeUserList()

    if option == "2":
        broadcastMessage()
    if ption == "3":
        sendPrivateMessage()
    else:
        receiverThread.exit()
        sys.exit()



def sendPrivateMessage():
    username = raw_input("Enter name of user: ")
    message = raw_input("Message:")
    d = {"type" : "MSG_USER", "message": message, "to" : username} 
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    menu()

def broadcastMessage():
    message = raw_input("Message:")
    d = {"type": "MSG_ALL", "message": message}
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    menu()

def seeUserList():
    d = {"type": "USER_LIST"}
    encoded = json.dumps(d)
    sock.send(encoded + "\n")
    menu()
    

def receiver():
    while True:
        data = sock.recv(1024)
        print data


def main():
    start()


if __name__ == "__main__":
    main()
