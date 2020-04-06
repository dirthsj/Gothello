package main

import (
	"gothello/HttpHandlers"
	"gothello/WebsocketHandlers"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handlerMap := make(map[string]http.Handler)
	fs := http.FileServer(http.Dir("./www"))
	handlerMap["/"] = fs
	handlerMap["/time"] = &HttpHandlers.TimeHandler{}
	handlerMap["/register"] = &HttpHandlers.PlayerRegistrationHandler{}
	handlerMap["/playerList"] = &HttpHandlers.PlayerListHandler{}
	handlerMap["/createGame"] = &HttpHandlers.CreateGameHandler{}
	chatHub := WebsocketHandlers.NewChatHub()
	go chatHub.Run()
	http.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) {
		WebsocketHandlers.ServeChatWs(chatHub, w, r)
	})
	gameHub := WebsocketHandlers.NewGameHub()
	go gameHub.Run()
	http.HandleFunc("/ws/game/", func(w http.ResponseWriter, r *http.Request) {
		WebsocketHandlers.ServeGameWs(gameHub, w, r)
	})
	handlerMap["/playerMove"] = &HttpHandlers.MovePlayerHandler{GameHub:gameHub}
	for key, element := range handlerMap {
		http.Handle(key, element)
	}
	log.Fatal(http.ListenAndServe(":" + port, nil))
}