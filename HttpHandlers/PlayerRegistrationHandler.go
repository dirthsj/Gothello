package HttpHandlers

import (
	"encoding/json"
	"gothello/Game"
	"net/http"
)

type PlayerRegistrationHandler struct{}

type PlayerIdRequest struct {
	PlayerName string
}

type PlayerIdResponse struct {
	PlayerId int
}

func(s *PlayerRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response PlayerIdResponse
	var request PlayerIdRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	response.PlayerId = Game.GetGothelloServerState().ConnectPlayer(request.PlayerName)
	_ = WriteJsonResponse(w, response)
}