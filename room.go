package main

type Room struct {
	code    string
	players map[*Player]*Player
}
