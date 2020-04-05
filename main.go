package main

import (
	"gothello/HttpHandlers"
	"gothello/WebsocketHandlers"
	"log"
	"net/http"

	_ "github.com/heroku/x/hmetrics/onload"
)


func main() {
	handlerMap := make(map[string]http.Handler)
	fs := http.FileServer(http.Dir("./www"))
	handlerMap["/"] = fs
	handlerMap["/time"] = &HttpHandlers.TimeHandler{}
	handlerMap["/playerMove"] = &HttpHandlers.MovePlayerHandler{}
	handlerMap["/register"] = &HttpHandlers.PlayerRegistrationHandler{}
	handlerMap["/playerList"] = &HttpHandlers.PlayerListHandler{}
	for key, element := range handlerMap {
		http.Handle(key, element)
	}
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
	log.Fatal(http.ListenAndServe(":80", nil))
}