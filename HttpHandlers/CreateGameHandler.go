package HttpHandlers

import (
	"encoding/json"
	"gothello/Game"
	"gothello/Security"
	"log"
	"net/http"
)

type CreateGameHandler struct{}

type createGameRequest struct {
	OpponentId int
}

func(s *CreateGameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request createGameRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&request)
	if decodeErr != nil {
		log.Println("Decoding Error: ", decodeErr.Error())
	}
	token, err := Security.GetTokenFromRequestHeader(r.Header)
	if err != nil {
		log.Println("Error getting token: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	playerId := Security.GetPlayerIdFromToken(token)
	if playerId < 1 {
		log.Println("Error getting Id from token", playerId)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	player1Available := Game.GetGothelloServerState().IsPlayerOnline(playerId) && Game.GetGothelloServerState().GetGame(playerId) == nil
	player2Available := Game.GetGothelloServerState().IsPlayerOnline(request.OpponentId) && Game.GetGothelloServerState().GetGame(request.OpponentId) == nil

	if player1Available && player2Available {
		Game.GetGothelloServerState().CreateGame(playerId, request.OpponentId)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("Player 1 Available: ", player1Available)
		log.Println("Player 2 Available: ", player2Available)
		log.Println( "Opponent Id: ", request.OpponentId)
		w.WriteHeader(http.StatusBadRequest)
	}
}