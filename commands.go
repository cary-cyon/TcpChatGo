package main

type commandId int

const (
	CMD_NiCK commandId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandId
	client *client
	args   []string
}
