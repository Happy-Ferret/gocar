package main

import (
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	CARCOLOR            = termbox.ColorGreen
	NOTHINGCOLOR        = termbox.ColorBlue
	BLOCKCOLOR          = termbox.ColorRed
	TEXTCOLOR           = termbox.ColorWhite
	TEXTBACKGROUNDCOLOR = termbox.ColorBlack

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
	printString("n start/pause game", startx, correntLine)
	correntLine += 2 //one empty line
	printString("a ←", startx, correntLine)
	correntLine++
	printString("s ↓", startx, correntLine)
	correntLine++
	printString("d →", startx, correntLine)
	correntLine++
	printString("w ↑", startx, correntLine)
	correntLine++
	printString("o exit", startx, correntLine)
	correntLine += 2
	if game.paused {
		printString("Pause", startx, correntLine)
		correntLine++
	}
	stringTime := strings.Join([]string{"Step by", game.time.String()}, " ")
	printString(stringTime, startx, correntLine)
}

func printString(line string, startX, startY int) {
	for _, char := range line {
		termbox.SetCell(startX, startY, char, TEXTCOLOR, TEXTBACKGROUNDCOLOR)
		startX++ //go to the next cell
	}
}
