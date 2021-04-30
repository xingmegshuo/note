package main

import (
	"log"
	"net"
	"time"
)

func main() {
	connOut, err := net.DialTimeout("tcp", "127.0.0.1:8001", time.Duration(1)*time.Second)
	
}
