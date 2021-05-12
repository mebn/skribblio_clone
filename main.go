package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Running...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/Design/", http.StripPrefix("/Design/", http.FileServer(http.Dir("Design"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/join", joinPage)
	http.HandleFunc("/create", createPage)
	http.HandleFunc("/room", roomPage)
	http.HandleFunc("/ws", WSEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func joinPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/join.html")
}

func createPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/create.html")
}

func roomPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "POST err: %v", err)
			return
		}

		// send these to socket.go via cookies
		handleCookies(w, r)

		http.ServeFile(w, r, "./public/room.html")
	} else {
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func handleCookies(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	roomnumber := r.FormValue("roomnumber")
	roomtype := r.FormValue("roomtype")

	usernameCookie := http.Cookie{
		Name:  "username",
		Value: username,
	}

	roomnumberCookie := http.Cookie{
		Name:  "roomnumber",
		Value: roomnumber,
	}

	roomtypeCookie := http.Cookie{
		Name:  "roomtype",
		Value: roomtype,
	}

	http.SetCookie(w, &usernameCookie)
	http.SetCookie(w, &roomnumberCookie)
	http.SetCookie(w, &roomtypeCookie)
}
