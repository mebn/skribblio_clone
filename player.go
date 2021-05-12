package main

import "github.com/gorilla/websocket"

type Player struct {
	name   string
	conn   *websocket.Conn
	room   *Room
	isHost bool
}

func NewPlayer(name string, conn *websocket.Conn, room *Room) *Player {
	temp := &Player{name, conn, room, false}
	temp.room.players[temp] = temp

	return temp
}
