package HttpHandlers

import (
	"gothello/Game"
	"net/http"
)

type PlayerListHandler struct{}

func(s *PlayerListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response = Game.GetGothelloServerState().GetOnlinePlayerList()
	_ = WriteJsonResponse(w, response)
}