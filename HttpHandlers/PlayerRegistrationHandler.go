package HttpHandlers

import (
	"encoding/json"
	"gothello/Game"
	"gothello/Security"
	"net/http"
)

type PlayerRegistrationHandler struct{}

type PlayerIdRequest struct {
	PlayerName string
}

type PlayerIdResponse struct {
	PlayerId int
	PlayerToken string
}

func(s *PlayerRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response PlayerIdResponse
	var request PlayerIdRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	response.PlayerId = Game.GetGothelloServerState().ConnectPlayer(request.PlayerName)
	token, err := Security.CreateGothelloToken(response.PlayerId)
	if err != nil {
		WriteBadRequestResponse(w, err.Error())
		return
	}
	response.PlayerToken = token
	_ = WriteJsonResponse(w, response)
}