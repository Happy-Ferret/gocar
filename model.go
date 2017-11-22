package main

import (
	"math/rand"
	"time"
)

//const about size of game
const (
	FIRSTLINE      = 0
	GAMEY          = 11
	GAMEX          = 50
	STARTPOSITIONY = 5
	STARTPOSITIONX = 49 //==GAMEX - 1
)

//const about time
const (
	startTime         = 400
	deltaTime         = 1
	minTime           = 100
	timeGoldGenerator = 500
)

type SellContains int

const (
	Nothing SellContains = iota
	Car
	Block
	Gold
)

type GameStatus int

const (
	Started GameStatus = iota
	Ended
)

type PositionSetter int

const (
	Left PositionSetter = iota
	Right
	Up
	Down
)

type Game struct {
	goldCount int
	time      time.Duration
	deltaTime time.Duration
	allTime   time.Duration
	board     [GAMEY][GAMEX]SellContains
	status    GameStatus
	paused    bool
	carx      int
	cary      int
}

func getNewGame() *Game {
	game := new(Game)
	game.board[STARTPOSITIONY][STARTPOSITIONX] = Car
	game.carx = STARTPOSITIONX
	game.cary = STARTPOSITIONY
	game.time = startTime * time.Millisecond
	game.deltaTime = deltaTime * time.Millisecond
	game.paused = false
	game.goldCount = 0
	return game
}

func (game *Game) addBlock(row int) {
	if game.board[row][FIRSTLINE] == Car {
		game.status = Ended
		return
	}
	game.board[row][FIRSTLINE] = Block
}

func (game *Game) setCarPosition(setter PositionSetter) {
	switch setter {
	case Left:
		if game.carx != 0 {
			game.board[game.cary][game.carx] = Nothing
			if game.board[game.cary][game.carx-1] == Block {
				game.status = Ended
				return
			}
			if game.board[game.cary][game.carx-1] == Gold {
				game.goldCount++
			}
			game.board[game.cary][game.carx-1] = Car
			game.carx--
		}
	case Right:
		if game.carx != GAMEX-1 {
			game.board[game.cary][game.carx] = Nothing
			if game.board[game.cary][game.carx+1] == Block {
				game.status = Ended
				return
			}
			if game.board[game.cary][game.carx+1] == Gold {
				game.goldCount++
			}
			game.board[game.cary][game.carx+1] = Car
			game.carx++
		}
	case Up:
		if game.cary != 0 {
			game.board[game.cary][game.carx] = Nothing
			if game.board[game.cary-1][game.carx] == Block {
				game.status = Ended
				return
			}
			if game.board[game.cary-1][game.carx] == Gold {
				game.goldCount++
			}
			game.board[game.cary-1][game.carx] = Car
			game.cary--
		}
	case Down:
		if game.cary != GAMEY-1 {
			game.board[game.cary][game.carx] = Nothing
			if game.board[game.cary+1][game.carx] == Block {
				game.status = Ended
				return
			}
			if game.board[game.cary+1][game.carx] == Gold {
				game.goldCount++
			}
			game.board[game.cary+1][game.carx] = Car
			game.cary++
		}
	}
}

//all bocks go from left to right
func (game *Game) nextStep() {
	for i := 0; i < GAMEY; i++ {
		for j := GAMEX - 2; j > -1; j-- {
			if game.board[i][j] == Block {
				if game.board[i][j+1] == Car {
					game.status = Ended
				} else {
					game.board[i][j] = Nothing
					game.board[i][j+1] = Block
				}
			}
		}
		if game.board[i][GAMEX-1] == Block {
			game.board[i][GAMEX-1] = Nothing
		}
	}
}

func (game *Game) doNothingInPause() {
	for game.paused {

	}
}

func (game *Game) doSteps() {
	for {
		game.doNothingInPause()
		row := rand.Intn(GAMEY)
		game.nextStep()
		if game.status == Ended {
			return
		}
		time.Sleep(game.time)
		game.allTime += game.time
		game.addBlock(row)
		if game.time > minTime*time.Millisecond {
			game.time -= game.deltaTime
		}
	}
}

func (game *Game) goldGenerator() {
	for game.status != Ended {
		game.doNothingInPause()
		time.Sleep(time.Millisecond * timeGoldGenerator)
		row := rand.Intn(GAMEY)
		col := rand.Intn(GAMEX)

		if game.board[row][col] == Nothing {
			game.board[row][col] = Gold
		}

	}
}
