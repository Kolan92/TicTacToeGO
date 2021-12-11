package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerEnum(t *testing.T) {
	t.Run("ToString", func(t *testing.T) {
		t.Run("PlayerX", func(t *testing.T) {
			player := PlayerX
			stringValue := player.toString()
			expectedValue := "PlayerX"
			assert.Equal(t, expectedValue, stringValue, "Incorrect string value {%q} of player, expected {%q}", stringValue, expectedValue)
		})
		t.Run("PlayerO", func(t *testing.T) {
			player := PlayerO
			stringValue := player.toString()
			expectedValue := "PlayerO"
			assert.Equal(t, expectedValue, stringValue, "Incorrect string value {%q} of player, expected {%q}", stringValue, expectedValue)
		})
	})

	t.Run("ToShortString", func(t *testing.T) {
		t.Run("PlayerX", func(t *testing.T) {
			player := PlayerX
			stringValue := player.toShortString()
			expectedValue := "X"
			assert.Equal(t, expectedValue, stringValue, "Incorrect string value {%q} of player, expected {%q}", stringValue, expectedValue)
		})
		t.Run("PlayerO", func(t *testing.T) {
			player := PlayerO
			stringValue := player.toShortString()
			expectedValue := "O"
			assert.Equal(t, expectedValue, stringValue, "Incorrect string value {%q} of player, expected {%q}", stringValue, expectedValue)
		})
		t.Run("None", func(t *testing.T) {
			player := None
			stringValue := player.toShortString()
			expectedValue := " "
			assert.Equal(t, expectedValue, stringValue, "Incorrect string value {%q} of player, expected {%q}", stringValue, expectedValue)
		})
	})
}

func TestStartingPlayer(t *testing.T) {
	game := NewGame()

	assert.Equal(t, PlayerX, game.currentPlayer)
}

func TestInitialBoard(t *testing.T) {
	game := NewGame()
	outerDimension := len(game.board)
	assert.Equal(t, 3, outerDimension)
	innerDimension := len(game.board[0])
	assert.Equal(t, 3, innerDimension)

	for _, row := range game.board {
		for _, player := range row {
			assert.Equal(t, None, player)
		}
	}
}

func TestFirstMoveUpdatesBoard(t *testing.T) {
	t.Run("0,0", func(t *testing.T) {
		game := NewGame()
		_ = game.MakeMove(0, 0)
		changedFiled := game.board[0][0]
		assert.Equal(t, PlayerX, changedFiled)
	})
	t.Run("1,2", func(t *testing.T) {
		game := NewGame()
		_ = game.MakeMove(1, 2)
		changedFiled := game.board[1][2]
		assert.Equal(t, PlayerX, changedFiled)
	})
}

func TestFirstMoveUpdatesCurrentPlayer(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(1, 2)
	assert.Equal(t, PlayerO, game.currentPlayer)
}

func TestMoveOutsideOfBordReturnsError(t *testing.T) {
	game := NewGame()
	err := game.MakeMove(-1, 2)
	assert.NotNil(t, err)
}

func TestPlayerCantMoveToAlreadySelectedFiled(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(1, 2)
	err := game.MakeMove(1, 2)
	assert.NotNil(t, err)
}

func TestSecondMoveShouldChangeSecondFiled(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(1, 2)
	_ = game.MakeMove(2, 2)

	assert.Equal(t, PlayerX, game.board[1][2])
	assert.Equal(t, PlayerO, game.board[2][2])
}

func TestGameIsNotEndedAfterFirstMove(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(1, 2)

	hasEnded, winner := game.GetStatus()
	assert.Equal(t, false, hasEnded)
	assert.Equal(t, None, winner)
}

func TestWinningConditions(t *testing.T) {
	t.Run("PlayerX", func(t *testing.T) {
		t.Run("Rows", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(0, 0) //PlayerX
			_ = game.MakeMove(1, 0) //PlayerO
			_ = game.MakeMove(0, 1) //PlayerX
			_ = game.MakeMove(2, 0) //PlayerO
			_ = game.MakeMove(0, 2) //PlayerX

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerX, winner)
		})

		t.Run("Columns", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(0, 0) //PlayerX
			_ = game.MakeMove(1, 0) //PlayerO
			_ = game.MakeMove(0, 1) //PlayerX
			_ = game.MakeMove(2, 0) //PlayerO
			_ = game.MakeMove(0, 2) //PlayerX

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerX, winner)
		})

		t.Run("Diagonal1", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(0, 0) //PlayerX
			_ = game.MakeMove(1, 0) //PlayerO
			_ = game.MakeMove(1, 1) //PlayerX
			_ = game.MakeMove(2, 0) //PlayerO
			_ = game.MakeMove(2, 2) //PlayerX

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerX, winner)
		})

		t.Run("Diagonal2", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(0, 2) //PlayerX
			_ = game.MakeMove(1, 0) //PlayerO
			_ = game.MakeMove(1, 1) //PlayerX
			_ = game.MakeMove(2, 2) //PlayerO
			_ = game.MakeMove(2, 0) //PlayerX

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerX, winner)
		})
	})

	t.Run("PlayerY", func(t *testing.T) {
		t.Run("Rows", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(1, 2) //PlayerX
			_ = game.MakeMove(0, 0) //PlayerO
			_ = game.MakeMove(1, 0) //PlayerX
			_ = game.MakeMove(0, 1) //PlayerO
			_ = game.MakeMove(2, 0) //PlayerX
			_ = game.MakeMove(0, 2) //PlayerO

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerO, winner)
		})

		t.Run("Columns", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(2, 1) //PlayerX
			_ = game.MakeMove(0, 0) //PlayerO
			_ = game.MakeMove(1, 0) //PlayerX
			_ = game.MakeMove(0, 1) //PlayerO
			_ = game.MakeMove(2, 0) //PlayerX
			_ = game.MakeMove(0, 2) //PlayerO

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerO, winner)
		})

		t.Run("Diagonal1", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(1, 2) //PlayerX
			_ = game.MakeMove(0, 0) //PlayerO
			_ = game.MakeMove(2, 1) //PlayerX
			_ = game.MakeMove(1, 0) //PlayerO
			_ = game.MakeMove(1, 1) //PlayerX
			_ = game.MakeMove(2, 0) //PlayerO

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerO, winner)
		})

		t.Run("Diagonal2", func(t *testing.T) {
			game := NewGame()
			_ = game.MakeMove(1, 2) //PlayerX
			_ = game.MakeMove(0, 2) //PlayerO
			_ = game.MakeMove(1, 0) //PlayerX
			_ = game.MakeMove(1, 1) //PlayerO
			_ = game.MakeMove(2, 2) //PlayerX
			_ = game.MakeMove(2, 0) //PlayerO

			hasEnded, winner := game.GetStatus()
			assert.True(t, hasEnded)
			assert.Equal(t, PlayerO, winner)
		})
	})

}

func TestDraw(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(0, 0) //PlayerX
	_ = game.MakeMove(0, 1) //PlayerO
	_ = game.MakeMove(0, 2) //PlayerX
	_ = game.MakeMove(1, 2) //PlayerO
	_ = game.MakeMove(1, 0) //PlayerX
	_ = game.MakeMove(2, 0) //PlayerO
	_ = game.MakeMove(1, 1) //PlayerX
	_ = game.MakeMove(2, 2) //PlayerO
	_ = game.MakeMove(2, 1) //PlayerX

	hasEnded, winner := game.GetStatus()
	assert.True(t, hasEnded)
	assert.Equal(t, None, winner)
}
func TestGame_EmptyGetBoard(t *testing.T) {
	game := NewGame()
	expectedBoardString :=
		`[ ][ ][ ]
[ ][ ][ ]
[ ][ ][ ]`

	actualBoardString := game.GetBoard()
	assert.Equal(t, expectedBoardString, actualBoardString)
}

func TestGame_GetBoard(t *testing.T) {
	game := NewGame()
	_ = game.MakeMove(0, 0) //PlayerX
	_ = game.MakeMove(0, 1) //PlayerO
	_ = game.MakeMove(0, 2) //PlayerX
	_ = game.MakeMove(1, 2) //PlayerO
	_ = game.MakeMove(1, 0) //PlayerX
	_ = game.MakeMove(2, 0) //PlayerO
	_ = game.MakeMove(1, 1) //PlayerX
	_ = game.MakeMove(2, 2) //PlayerO

	expectedBoardString :=
		`[X][O][X]
[X][X][O]
[O][ ][O]`

	actualBoardString := game.GetBoard()
	assert.Equal(t, expectedBoardString, actualBoardString)

}
