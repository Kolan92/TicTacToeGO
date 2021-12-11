package game

import (
	"fmt"
	"strings"
)

type IGame interface {
	MakeMove(coordinateX, coordinateY int) error
	GetStatus() (bool, Player)
	GetCurrentPlayer() Player
	GetBoard() string
}

type Player uint8

const (
	None Player = iota
	PlayerX
	PlayerO
)

func (player Player) toString() string {
	switch player {
	case PlayerX:
		return "PlayerX"
	case PlayerO:
		return "PlayerO"
	}
	return ""
}

func (player Player) toShortString() string {
	switch player {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	case None:
		return " "
	}
	return ""
}

type Game struct {
	currentPlayer Player
	board         [3][3]Player
}

func NewGame() *Game {
	return &Game{
		currentPlayer: PlayerX,
	}
}

func (game *Game) MakeMove(coordinateX, coordinateY int) error {
	if coordinateX > 2 || coordinateX < 0 || coordinateY > 2 || coordinateY < 0 {
		return fmt.Errorf("invalid coordinates x: {%d} y: {%d}", coordinateX, coordinateY)
	}

	if game.board[coordinateX][coordinateY] != None {
		return fmt.Errorf("filed with coordinates x: {%d} y: {%d} already occupied", coordinateX, coordinateY)
	}
	game.board[coordinateX][coordinateY] = game.currentPlayer
	if game.currentPlayer == PlayerX {
		game.currentPlayer = PlayerO
	} else {
		game.currentPlayer = PlayerX
	}

	return nil
}

func (game *Game) GetStatus() (bool, Player) {
	for i := 0; i < 3; i++ {
		//columns check
		row0 := game.board[i][0]
		row1 := game.board[i][1]
		row2 := game.board[i][2]

		if row0 != None && row0 == row1 && row0 == row2 {
			return true, row0
		}

		//rows check
		col0 := game.board[0][i]
		col1 := game.board[1][i]
		col2 := game.board[2][i]

		if col0 != None && col0 == col1 && col0 == col2 {
			return true, col0
		}
	}

	//diagonals
	{
		filed0 := game.board[0][0]
		filed1 := game.board[1][1]
		filed2 := game.board[2][2]

		if filed0 != None && filed0 == filed1 && filed0 == filed2 {
			return true, filed0
		}
	}
	{
		filed0 := game.board[0][2]
		filed1 := game.board[1][1]
		filed2 := game.board[2][0]

		if filed0 != None && filed0 == filed1 && filed0 == filed2 {
			return true, filed0
		}
	}

	for _, row := range game.board {
		for _, field := range row {
			if field == None {
				return false, None
			}
		}
	}

	return true, None
}

func (game *Game) GetCurrentPlayer() Player {
	return game.currentPlayer
}

func (game *Game) GetBoard() string {
	var sb strings.Builder
	for rowIndex, row := range game.board {
		for _, field := range row {
			sb.WriteString("[")
			sb.WriteString(field.toShortString())
			sb.WriteString("]")
		}
		if rowIndex != 2 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
