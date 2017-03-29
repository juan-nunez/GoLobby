package server

import (
    "net"
)

type User struct {
    id string
    name string
    conn net.Conn
}
