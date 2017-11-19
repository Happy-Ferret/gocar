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
	START_TIME = 500
	DELTA_TIME = 1
)

type SellContains int

const (
	Nothing SellContains = iota
	Car
	Block
)

type GameStatus int

const (
	Started GameStatus = iota
	Paused
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
	time      time.Duration
	deltaTime time.Duration
	board     [GAMEY][GAMEX]SellContains
	status    GameStatus
	carx      int
	cary      int
}

func getNewGame() *Game {
	game := new(Game)
	game.board[STARTPOSITIONY][STARTPOSITIONX] = Car
	game.carx = STARTPOSITIONX
	game.cary = STARTPOSITIONY
	game.time = START_TIME * time.Millisecond
	game.deltaTime = DELTA_TIME * time.Millisecond
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

func (game *Game) timeBlock() {
	for {
		if game.status == Paused {
			continue
		}
		row := rand.Intn(GAMEY)
		game.nextStep()
		time.Sleep(game.time)
		game.time -= game.deltaTime
		game.addBlock(row)
		if game.status == Ended {
			return
		}
	}
}
