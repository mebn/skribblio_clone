package main

import "github.com/gorilla/websocket"

type Player struct {
	name   string
	conn   *websocket.Conn
	room   *Room
	isHost bool
	isTurn bool
	score  int
}

func NewPlayer(name string, conn *websocket.Conn, room *Room) *Player {
	temp := &Player{name, conn, room, false, false, 0}
	temp.room.players[temp] = temp

	return temp
}
