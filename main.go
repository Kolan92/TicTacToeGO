package main

import "ticktactoego/game"

func main() {
	println("TickTackToe game")

	gameManger := game.NewGameManger(game.NewGame(), game.NewOutput())
	gameManger.Run()
}
