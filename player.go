package main

type Player struct {
	name   string
	isHost bool
	points int
	room   *Room
}

func NewPlayer(name string, room *Room) *Player {
	return &Player{
		name:   name,
		room:   room,
		isHost: false,
	}
}
