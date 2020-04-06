package Game

var (
	state *gothelloServerState
)

type gothelloServerState struct {
	onlinePlayers map[int]PlayerRegister
	onlineGames map[int]*GothelloGame
	nextPlayerId  int
}

func GetGothelloServerState() *gothelloServerState {
	if state == nil {
		state = &gothelloServerState{
			onlinePlayers: map[int]PlayerRegister{},
			onlineGames: map[int]*GothelloGame{},
			nextPlayerId:  1,
		}
	}
	return state
}

type PlayerRegister struct {
	PlayerId int
	PlayerName string
}

func(s *gothelloServerState) DisconnectPlayer(playerId int) {
	delete(s.onlinePlayers, playerId)
}

func(s *gothelloServerState) ConnectPlayer(playerName string) int {
	player := PlayerRegister{
		PlayerId:   s.nextPlayerId,
		PlayerName: playerName,
	}
	s.nextPlayerId++
	s.onlinePlayers[player.PlayerId] = player
	return player.PlayerId
}

func(s *gothelloServerState) IsPlayerOnline(playerId int) bool {
	if _, ok := s.onlinePlayers[playerId]; ok {
		return true
	}
	return false
}

func(s *gothelloServerState) GetOnlinePlayerList() []PlayerRegister {
	var playerList []PlayerRegister
	for i := 0; i < s.nextPlayerId; i++ {
		if val, ok := s.onlinePlayers[i]; ok {
			playerList = append(playerList, val)
		}
	}
	return playerList
}

func(s *gothelloServerState) CreateGame(userAId int, userBId int) {
	var game = GothelloGame{whitePlayerId:userAId, blackPlayerId:userBId}
	game.init()
	s.onlineGames[userAId] = &game
	s.onlineGames[userBId] = &game
}

func(s *gothelloServerState) GetGame(userId int) *GothelloGame {
	return s.onlineGames[userId]
}

func(s *gothelloServerState) DeleteGame(userId int) {
	var toDelete = s.onlineGames[userId]
	var userId1, userId2 = toDelete.GetPlayerIds()
	delete(s.onlineGames, userId1)
	delete(s.onlineGames, userId2)
}