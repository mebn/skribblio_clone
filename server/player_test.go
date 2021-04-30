package main

import "testing"

func TestNewPlayer(t *testing.T) {
	room1 := new(Room)
	player1 := NewPlayer("namn", room1)

	want := "namn"
	if got := player1.name; got != want {
		t.Errorf("NewPlayer(name, room1).name = %s, want %s \n", got, want)
	}
}
