package main

import "log"

type Room struct {
	code        string
	players     map[*Player]*Player
	currentWord string
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

func (room *Room) SendEmptyToRoom(code string) {
	msg := "{\"code\":\"" + code + "\",\"data\":\"" + "" + "\"}"

	for _, client := range room.players {
		err := client.conn.WriteMessage(1, []byte(msg))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
