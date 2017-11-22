package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	CARCOLOR            = termbox.ColorGreen
	NOTHINGCOLOR        = termbox.ColorBlue
	BLOCKCOLOR          = termbox.ColorRed
	TEXTCOLOR           = termbox.ColorWhite
	TEXTBACKGROUNDCOLOR = termbox.ColorBlack
	GOLDCOLOR           = termbox.ColorYellow

	GAMEMARGINX      = 2
	GAMEMARGINY      = 1
	TEXTBLOCKMARGINY = 1
	TEXTBLOCKMARGINX = 3
)

func printGame(game *Game) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			if game.board[i][j] == Car {
				termbox.SetCell(GAMEMARGINX+j, GAMEMARGINY+i, 1, CARCOLOR, CARCOLOR)
			} else if game.board[i][j] == Block {
				termbox.SetCell(GAMEMARGINX+j, GAMEMARGINY+i, 1, BLOCKCOLOR, BLOCKCOLOR)
			} else if game.board[i][j] == Gold {
				termbox.SetCell(GAMEMARGINX+j, GAMEMARGINY+i, 1, GOLDCOLOR, GOLDCOLOR)
			} else {
				termbox.SetCell(GAMEMARGINX+j, GAMEMARGINY+i, 1, NOTHINGCOLOR, NOTHINGCOLOR)
			}
		}
	}
	printAboutTextBlock()
	termbox.Flush()
}

func printAboutTextBlock() {
	correntLine := TEXTBLOCKMARGINY
	startx := GAMEMARGINX + GAMEX + TEXTBLOCKMARGINX
	printString("p start/pause game", startx, correntLine)
	correntLine += 2 //one empty line
	printString("a ←", startx, correntLine)
	correntLine++
	printString("s ↓", startx, correntLine)
	correntLine++
	printString("d →", startx, correntLine)
	correntLine++
	printString("w ↑", startx, correntLine)
	correntLine++
	printString("n end this game", startx, correntLine)
	correntLine += 2
	if game.paused {
		printString("Pause", startx, correntLine)
		correntLine++
	}
	stringTime := strings.Join([]string{"Step by", game.time.String()}, " ")
	printString(stringTime, startx, correntLine)
	correntLine++
	stringGold := strings.Join([]string{"Gold taken", strconv.Itoa(game.goldCount)}, " ")
	printString(stringGold, startx, correntLine)
	correntLine++
	stringAllTime := strings.Join([]string{"Game's time", game.allTime.String()}, " ")
	printString(stringAllTime, startx, correntLine)
	correntLine++
	printString(totalScore(game), startx, correntLine)
}

func printString(line string, startX, startY int) {
	for _, char := range line {
		termbox.SetCell(startX, startY, char, TEXTCOLOR, TEXTBACKGROUNDCOLOR)
		startX++ //go to the next cell
	}
}

func printGameEnded() {
	//waiting gorutins with time
	for i := 1; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		printString("#", 1, i)
	}

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	stringGold := strings.Join([]string{"You have take", strconv.Itoa(game.goldCount), "gold!"}, " ")
	printString(stringGold, 1, 2)
	stringAllTime := strings.Join([]string{"You played", game.allTime.String()}, " ")
	printString(stringAllTime, 1, 3)
	printString(totalScore(game), 1, 7)
	printString("Press [esc] to exit", 1, 4)
	printString("Press [enter] new game", 1, 5)
	termbox.Flush()
}

func totalScore(game *Game) string {
	min := int64(game.allTime.Minutes())
	goldCount := (int64)(game.goldCount)
	//some f(min, goldCount)
	score := (int)(min*2 + goldCount)
	stringScore := strings.Join([]string{"Total score", strconv.Itoa(score)}, " ")
	return stringScore
}
