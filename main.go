package main

import (
	"log"
	"net"
	"time"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	defer listener.Close()
	log.Printf("starter server on :8080 at %d-%d:%d,", time.Now().Day(), time.Now().Hour(), time.Now().Minute())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
