package main

type Room struct {
	code    string
	players []*Player
}

func NewRoom(code string) *Room {
	return &Room{
		code: code,
	}
}
