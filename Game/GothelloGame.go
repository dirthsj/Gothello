package Game

import "encoding/json"

type GothelloGame struct {
	whitePlayerId int
	blackPlayerId int
	board [BoardDimension][BoardDimension]BoardSquareState
	turn int
}

type BoardSquareState string

const(
	Empty BoardSquareState = "empty"
	Black BoardSquareState = "black"
	White BoardSquareState = "white"
	BoardDimension int = 8
)

func (game *GothelloGame) init() {
	game.turn = 0
	game.board = [BoardDimension][BoardDimension]BoardSquareState{}
	for x := 0; x < BoardDimension; x++ {
		game.board[x] = [BoardDimension]BoardSquareState{}
		for y := 0; y < BoardDimension; y++ {
			game.board[x][y] = Empty
		}
	}
	game.board[3][3] = White
	game.board[3][4] = Black
	game.board[4][3] = Black
	game.board[4][4] = White
}

func (game *GothelloGame) get(x int, y int) BoardSquareState {
	return game.board[x][y]
}

func (game *GothelloGame) set(x int, y int, state BoardSquareState) {
	game.board[x][y] = state
}

func (game *GothelloGame) TryMove(playerId int, x int, y int) bool {
	if game.getCurrentPlayerId() != playerId {
		return false
	}
	moveState := Empty
	if playerId == game.whitePlayerId {
		moveState = White
	} else if playerId == game.blackPlayerId {
		moveState = Black
	} else {
		return false
	}
	if game.get(x, y) != Empty {
		return false
	}
	result := false
	result = game.tryFlip(x    , y + 1,  0,  1, moveState, Empty) || result
	result = game.tryFlip(x + 1, y + 1,  1,  1, moveState, Empty) || result
	result = game.tryFlip(x + 1, y    ,  1,  0, moveState, Empty) || result
	result = game.tryFlip(x + 1, y - 1,  1, -1, moveState, Empty) || result
	result = game.tryFlip(x    , y - 1,  0, -1, moveState, Empty) || result
	result = game.tryFlip(x - 1, y - 1, -1, -1, moveState, Empty) || result
	result = game.tryFlip(x - 1, y    , -1,  0, moveState, Empty) || result
	result = game.tryFlip(x - 1, y + 1, -1,  1, moveState, Empty) || result
	if result {
		game.set(x, y, moveState)
		game.turn++
	}
	return result
}

func (game *GothelloGame) tryFlip(x int, y int, xi int, yi int, newSquareState BoardSquareState, previousSquareState BoardSquareState) bool {
	// base case: out of bounds
	if x < 0 || x >= BoardDimension || y < 0 || y >= BoardDimension {
		return false
	}
	// base case: find a piece of the same color
	if game.get(x, y) == newSquareState {
		return previousSquareState != Empty
	}
	// base case: find an empty square
	if game.get(x,y) == Empty {
		return false
	}
	// recursive case: find a piece of the other color
	if game.tryFlip(x + xi, y + yi, xi, yi, newSquareState, game.get(x,y)) {
		game.set(x, y, newSquareState)
		return true
	}
	return false
}

func (game *GothelloGame) getCurrentPlayerId() int {
	if game.turn % 2 == 0 {
		return game.whitePlayerId
	}
	return game.blackPlayerId
}

func (game *GothelloGame) GetPlayerIds() (int, int) {
	return game.whitePlayerId, game.blackPlayerId;
}

func (game *GothelloGame) GetBoardUpdate() []byte {
	update, _ := json.Marshal(game.board)
	return update
}