package main

import (
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

	http.HandleFunc("/", homePage)
	http.HandleFunc("/room", roomPage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func roomPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/room.html")
}
