package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- ICommand
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nick":
			c.commands <- commandNick{
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- commandJoin{
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- commandRooms{
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- commandMsg{
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- commandQuit{
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("uknowncommand: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	log.Printf("Send error at %s", c.conn.RemoteAddr())
	c.conn.Write([]byte("ERR: " + err.Error()))
}

func (c *client) msg(msg string) {
	log.Printf("Send messend at %s", c.conn.RemoteAddr())
	c.conn.Write([]byte(">" + msg + "\n"))
}
