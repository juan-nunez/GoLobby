package server

import (
    "net"
)

type User struct {
    id string
    conn net.Conn
}
