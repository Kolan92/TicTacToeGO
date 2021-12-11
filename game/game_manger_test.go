package game

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

type FakeOutput struct {
	mock.Mock
}

func (m *FakeOutput) Write(text string) {
	_ = m.Called(text)
}

func (m *FakeOutput) Read() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

type FakeGame struct {
	mock.Mock
}

func (m *FakeGame) GetCurrentPlayer() Player {
	args := m.Called()
	return args.Get(0).(Player)
}

func (m *FakeGame) MakeMove(coordinateX, coordinateY int) error {
	args := m.Called(coordinateX, coordinateY)
	return args.Error(0)
}

func (m *FakeGame) GetStatus() (bool, Player) {
	args := m.Called()
	return args.Bool(0), args.Get(1).(Player)
}

func (m *FakeGame) GetBoard() string {
	args := m.Called()
	return args.String(0)
}

func TestPrintsGreetings(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(true, None)

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "Hello to tic tac toe game")
}

func TestGameMangerDraw(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(true, None)

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "Draw")
}

func TestGameMangerPlayerXWins(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(true, PlayerX)

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "PlayerX has won!")
}

func TestGameMangerPlayerOWins(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(true, PlayerO)

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "PlayerO has won!")
}

func TestPrintsFirstPlayer(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(true, None).Once()
	fakeGame.On("GetCurrentPlayer").Return(PlayerX).Once()
	output.On("Read").Return("1", nil)
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	fakeGame.On("GetBoard").Return("test game board")

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "PlayerX move")
}

func TestFirstPlayerMakesMove(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(true, None).Once()
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	fakeGame.On("GetCurrentPlayer").Return(PlayerX)
	output.On("Read").Return("1", nil).Once()
	output.On("Read").Return("2", nil).Once()
	fakeGame.On("GetBoard").Return("test game board")

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	fakeGame.AssertCalled(t, "MakeMove", 1, 2)
}

func TestPlayerInputsIncorrectValue(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(true, None).Once()
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	fakeGame.On("GetCurrentPlayer").Return(PlayerX)
	output.On("Read").Return("not number", nil).Once()
	output.On("Read").Return("1", nil)
	fakeGame.On("GetBoard").Return("test game board")

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "Incorrect number format, please try again")
}

func TestPlayerInputsIndexOutOfRange(t *testing.T) {
	t.Run("3", func(t *testing.T) {

		fakeGame := new(FakeGame)
		output := new(FakeOutput)
		output.On("Write", mock.AnythingOfType("string")).Return()
		fakeGame.On("GetStatus").Return(false, None).Once()
		fakeGame.On("GetStatus").Return(true, None).Once()
		fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		fakeGame.On("GetCurrentPlayer").Return(PlayerX)
		output.On("Read").Return("1", nil)
		fakeGame.On("GetBoard").Return("test game board")
		output.On("Read").Return("3", nil).Once()

		gameManger := NewGameManger(fakeGame, output)
		gameManger.Run()

		output.AssertCalled(t, "Write", "Incorrect number format, please try again")
	})

	t.Run("-1", func(t *testing.T) {

		fakeGame := new(FakeGame)
		output := new(FakeOutput)
		output.On("Write", mock.AnythingOfType("string")).Return()
		fakeGame.On("GetStatus").Return(false, None).Once()
		fakeGame.On("GetStatus").Return(true, None).Once()
		fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		fakeGame.On("GetCurrentPlayer").Return(PlayerX)
		output.On("Read").Return("1", nil)
		fakeGame.On("GetBoard").Return("test game board")
		output.On("Read").Return("-1", nil).Once()

		gameManger := NewGameManger(fakeGame, output)
		gameManger.Run()

		output.AssertCalled(t, "Write", "Value outside of correct range, please try again")
	})
}

func TestPlayerInputsOccupiedField(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(true, None).Once()
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("Test error")).Once()
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	fakeGame.On("GetCurrentPlayer").Return(PlayerX)
	output.On("Read").Return("1", nil).Once()
	output.On("Read").Return("1", nil).Once()
	output.On("Read").Return("2", nil).Once()
	output.On("Read").Return("2", nil).Once()
	fakeGame.On("GetBoard").Return("test game board")

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	output.AssertCalled(t, "Write", "Filed already occupied, please try again")
}

func TestPrintsBoardAfterEveryMove(t *testing.T) {
	fakeGame := new(FakeGame)
	output := new(FakeOutput)
	output.On("Write", mock.AnythingOfType("string")).Return()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(false, None).Once()
	fakeGame.On("GetStatus").Return(true, None).Once()
	fakeGame.On("MakeMove", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	output.On("Read").Return("1", nil)
	fakeGame.On("GetBoard").Return("test game board")
	fakeGame.On("GetCurrentPlayer").Return(PlayerX)

	gameManger := NewGameManger(fakeGame, output)
	gameManger.Run()

	fakeGame.AssertNumberOfCalls(t, "GetBoard", 2)
}
