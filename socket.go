package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]*Player)
var rooms = make(map[string]*Room)
var upgrader = websocket.Upgrader{}

var words = []string{
	"banana",
	"pineapple",
	"apple",
	"pear",
	"grape",
	"citrus",
}

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
		rooms[roomnumber] = &Room{roomnumber, map[*Player]*Player{}, ""}
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
	data := fmt.Sprintf("%v", obj["data"])

	// handle message based on code/channel

	receiveOn("is host", code, func() {
		msg := strconv.FormatBool(clients[conn].isHost)
		sendOn("is host", msg, conn)
	})

	receiveOn("should start game", code, func() {
		currentRoom.SendToRoom(msgType, msg)

		clients[conn].isTurn = true
		data2send := generateWordsToSend()
		sendOn("is turn", data2send, conn)
	})

	receiveOn("sendToRoom", code, func() {
		currentRoom.SendToRoom(msgType, msg)
	})

	receiveOn("drawing", code, func() {
		currentRoom.SendToRoom(msgType, msg)
	})

	receiveOn("clear canvas", code, func() {
		currentRoom.SendToRoom(msgType, msg)
	})

	receiveOn("picked word", code, func() {
		currentRoom.currentWord = data

		currentRoom.SendEmptyToRoom("game start")
	})

	receiveOn("next turn", code, func() {
		player.isTurn = false

		// chose new drawing player at random
		var newCurrentPlayer *Player
		// this works because the order is random.
		for v, _ := range currentRoom.players {
			newCurrentPlayer = v
			break
		}

		newCurrentPlayer.isTurn = true
		data2send := generateWordsToSend()
		obj := "{\"code\":\"" + "is turn" + "\",\"data\":\"" + data2send + "\"}"
		newCurrentPlayer.conn.WriteMessage(1, []byte(obj))
		currentRoom.SendEmptyToRoom("clear canvas")
		currentRoom.SendEmptyToRoom("new turn")

	})
}

func generateWordsToSend() string {
	// choose 3 words from words and send them on "is turn"
	rand1 := rand.Intn(len(words))
	rand2 := rand.Intn(len(words))
	rand3 := rand.Intn(len(words))

	word1 := words[rand1]
	word2 := words[rand2]
	word3 := words[rand3]

	data2send := word1 + " " + word2 + " " + word3
	return data2send
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
