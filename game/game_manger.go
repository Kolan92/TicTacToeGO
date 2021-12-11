package game

import (
	"strconv"
	"strings"
)

type GameManger struct {
	game   IGame
	output Output
}

func NewGameManger(game IGame, output Output) *GameManger {
	return &GameManger{
		game:   game,
		output: output,
	}
}

func (this *GameManger) Run() {
	this.output.Write("Hello to tic tac toe game")
	hasEnded, winner := this.game.GetStatus()

	for !hasEnded {
		currentBoard := this.game.GetBoard()
		this.output.Write(currentBoard)

		this.handleUserMove()
		hasEnded, winner = this.game.GetStatus()
	}

	if winner == None {
		this.output.Write("Draw")
	} else {
		this.output.Write(winner.toString() + " has won!")
	}
}

func (this *GameManger) handleUserMove() {
	player := this.game.GetCurrentPlayer()
	this.output.Write(player.toString() + " move")

	coordinateX := this.getCoordinate("X")
	coordinateY := this.getCoordinate("Y")
	err := this.game.MakeMove(coordinateX, coordinateY)

	if err != nil {
		this.output.Write("Filed already occupied, please try again")
		this.handleUserMove()
	}
}

func (this *GameManger) getCoordinate(coordinateType string) int {
	this.output.Write("Please type number between 0 and 2 for coordinate " + coordinateType)
	coordinateText, err := this.output.Read()

	if err != nil {
		this.output.Write("Error reading input, please try again")
		return this.getCoordinate(coordinateType)
	}

	coordinate, err := strconv.Atoi(strings.TrimSuffix(coordinateText, "\n"))
	if err != nil {
		this.output.Write("Incorrect number format, please try again")
		return this.getCoordinate(coordinateType)
	}

	if coordinate < 0 || coordinate > 2 {
		this.output.Write("Value outside of correct range, please try again")
		return this.getCoordinate(coordinateType)
	}

	return coordinate
}
