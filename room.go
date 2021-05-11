package main

import "log"

type Room struct {
	code    string
	players map[*Player]*Player
}

func (room *Room) SendToRoom(msgType int, msg []byte) {
	for _, client := range room.players {
		err := client.conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
