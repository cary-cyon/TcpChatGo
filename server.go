package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan ICommand
}

func (s *server) run() {
	for cmd := range s.commands {
		cmd.doAction(s)
	}
}
func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan ICommand),
	}
}
func (s *server) newClient(conn net.Conn) {
	log.Printf("new client at %s", conn.RemoteAddr().String())
	c := &client{
		conn:     conn,
		nick:     "anonymus",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the romm", c.nick))
	}
}
