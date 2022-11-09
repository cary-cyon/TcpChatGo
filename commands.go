package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type ICommand interface {
	doAction(s *server)
}

type commandNick struct {
	client *client
	args   []string
}

func (com commandNick) doAction(s *server) {
	com.client.nick = com.args[1]
	com.client.msg(fmt.Sprintf("NICK: %s", com.client.nick))
}

type commandJoin struct {
	client *client
	args   []string
}

func (com commandJoin) doAction(s *server) {
	roomName := com.args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[com.client.conn.RemoteAddr()] = com.client
	s.quitCurrentRoom(com.client)
	com.client.room = r
	r.broadcast(com.client, fmt.Sprintf("%s has join the room", com.client.nick))
	com.client.msg(fmt.Sprintf("welcom to %s", r.name))
}

type commandRooms struct {
	client *client
	args   []string
}

func (com commandRooms) doAction(s *server) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	com.client.msg(fmt.Sprintf("rooms are: %s", strings.Join(rooms, ", ")))
}

type commandMsg struct {
	client *client
	args   []string
}

func (com commandMsg) doAction(s *server) {
	if com.client.room == nil {
		com.client.err(errors.New("JOIN room befor chating"))
		return
	}
	com.client.room.broadcast(com.client, com.client.nick+": "+strings.Join(com.args[1:], " "))
}

type commandQuit struct {
	client *client
	args   []string
}

func (com commandQuit) doAction(s *server) {
	log.Printf("client has disconnected: %s", com.client.conn.RemoteAddr().String())
	s.quitCurrentRoom(com.client)
	com.client.msg("sad to see you go :( ")
	com.client.conn.Close()
}
