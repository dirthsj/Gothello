package HttpHandlers

import (
	"encoding/json"
	"gothello/Game"
	"gothello/Security"
	"net/http"
)

type MovePlayerHandler struct{}

type playerMove struct {
	X int
	Y int
}

func(s *MovePlayerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, err := Security.GetTokenFromRequestHeader(r.Header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	playerId := Security.GetPlayerIdFromToken(token)
	if playerId < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var playerMove playerMove
	_ = json.NewDecoder(r.Body).Decode(&playerMove)
	var game = Game.GetGothelloServerState().GetGame(playerId)
	if game.TryMove(playerId, playerMove.X, playerMove.Y) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}