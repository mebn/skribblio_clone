package main

import "testing"

func TestNewRoom(t *testing.T) {
	room1 := NewRoom("abc123")

	want := "abc123"
	if got := room1.code; got != want {
		t.Errorf("NewRoom().code = %s, want %s \n", got, want)
	}
}
