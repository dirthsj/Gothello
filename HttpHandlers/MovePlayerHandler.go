package HttpHandlers

import (
	"encoding/json"
	"gothello/Game"
	"gothello/Security"
	"gothello/WebsocketHandlers"
	"log"
	"net/http"
)

type MovePlayerHandler struct{
	GameHub *WebsocketHandlers.GameHub
}

type playerMove struct {
	X int
	Y int
}

func(s *MovePlayerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	var playerMove playerMove
	_ = json.NewDecoder(r.Body).Decode(&playerMove)
	var game = Game.GetGothelloServerState().GetGame(playerId)
	if game == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if game.TryMove(playerId, playerMove.X, playerMove.Y) {
		w.WriteHeader(http.StatusOK)
		player1, player2 := game.GetPlayerIds()
		update := game.GetBoardUpdate()
		s.GameHub.MessageClientByPlayerId(player1, update)
		s.GameHub.MessageClientByPlayerId(player2, update)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}