package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]*Player)
var rooms = map[string]*Room{}
var upgrader = websocket.Upgrader{}

func WSEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// receive username and roomnumber from main.go via cookies
	var username string
	var roomnumber string
	var roomtype string
	for _, v := range r.Cookies() {
		if v.Name == "username" {
			username = v.Value
		} else if v.Name == "roomnumber" {
			roomnumber = v.Value
		} else if v.Name == "roomtype" {
			roomtype = v.Value
		}
	}

	fmt.Println(username, roomnumber, roomtype)

	var tempPlayer *Player
	if roomtype == "Join" {
		_, ok := rooms[roomnumber]
		if !ok {
			// room does not exist
			// redirect player back to /
		}
	} else if roomtype == "Create" {
		rooms[roomnumber] = &Room{roomnumber, map[*Player]*Player{}}
	}

	tempPlayer = NewPlayer(username, ws, rooms[roomnumber])
	if roomtype == "Create" {
		tempPlayer.isHost = true
	}
	clients[ws] = tempPlayer
	fmt.Println(rooms)

	// send username and roomnumber back to client
	sendOn("username", username, ws)
	sendOn("roomnumber", roomnumber, ws)

	log.Println("Client connected!")
	msgListener(ws)
}

// every client has one of these
func msgListener(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()

		// client disconnect
		if err != nil {
			delete(clients[conn].room.players, clients[conn])
			delete(clients, conn)
			log.Println(err)
			return
		}

		handleMessage(conn, msgType, msg)
	}
}

func handleMessage(conn *websocket.Conn, msgType int, msg []byte) {
	player := clients[conn]
	currentRoom := player.room

	// receiving string to map
	obj := make(map[string]interface{})
	err := json.Unmarshal(msg, &obj)
	if err != nil {
		panic(err)
	}

	// interface to string
	code := fmt.Sprintf("%v", obj["code"])
	// data := fmt.Sprintf("%v", obj["data"])

	// handle message based on code/channel

	receiveOn("sendToRoom", code, func() {
		currentRoom.SendToRoom(msgType, msg)
	})
}

// Send on a channel name to server. For example:
//     sendOn("say hello", "hello")
// will send "hello" to the receiver "say hello"
// on the server.
//
// name: The channel name to send on.
//
// msg: The message/data to send.
func sendOn(name, msg string, conn *websocket.Conn) {
	obj := "{\"code\":\"" + name + "\",\"data\":\"" + msg + "\"}"
	conn.WriteMessage(1, []byte(obj))
}

// Receives on a channel name, sent from server.
// For example:
//     receiveOn("say hello", code, func() {
//	   // do something with data and/or msg
//     })
//
// name: The channel name to send to.
//
// msg: The message/data to send,
//
// cb: A callback function to do something with received data from server.
func receiveOn(name string, code string, f func()) {
	if name == code {
		f()
	}
}
