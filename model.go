package main

import (
	"math/rand"
	"time"
)

type SellContains int
type GameStatus int
type PositionSetter int

const (
	Nothing SellContains = iota
	Car
	Block
	Gold

	Started GameStatus = iota
	Ended

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
			if game.carMovie(game.cary, game.carx-1) {
				return
			}
			game.board[game.cary][game.carx-1] = Car
			game.carx--
		}
	case Right:
		if game.carx != GAMEX-1 {
			game.board[game.cary][game.carx] = Nothing
			if game.carMovie(game.cary, game.carx+1) {
				return
			}
			game.board[game.cary][game.carx+1] = Car
			game.carx++
		}
	case Up:
		if game.cary != 0 {
			game.board[game.cary][game.carx] = Nothing
			if game.carMovie(game.cary-1, game.carx) {
				return
			}
			game.board[game.cary-1][game.carx] = Car
			game.cary--
		}
	case Down:
		if game.cary != GAMEY-1 {
			game.board[game.cary][game.carx] = Nothing
			if game.carMovie(game.cary+1, game.carx) {
				return
			}
			game.board[game.cary+1][game.carx] = Car
			game.cary++
		}
	}
}

//if game ended then true
func (game *Game) carMovie(i, j int) bool {
	if game.board[i][j] == Block {
		game.status = Ended
		return true
	}
	if game.board[i][j] == Gold {
		game.goldCount++
	}
	return false
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

func (game *Game) doSteps() {
	for {
		row := rand.Intn(GAMEY)
		game.nextStep()
		if game.status == Ended {
			return
		}
		time.Sleep(game.time)
		if game.status == Ended {
			return
		}
		game.allTime += game.time

		game.addBlock(row)
		if game.time > minTime*time.Millisecond {
			game.time -= game.deltaTime
		}
		if game.status == Ended {
			return
		}
	}
}

func (game *Game) goldGenerator() {
	for game.status != Ended {

		time.Sleep(time.Millisecond * timeGoldGenerator)
		row := rand.Intn(GAMEY)
		col := rand.Intn(GAMEX)

		if game.board[row][col] == Nothing {
			game.board[row][col] = Gold
		}

	}
}
