package main

import (
	"github.com/nsf/termbox-go"
)

const (
	CARCOLOR     = termbox.ColorGreen
	NOTHINGCOLOR = termbox.ColorBlue
	BLOCKCOLOR   = termbox.ColorRed

	GAMEMARGINX = 2
	GAMEMARGINY = 1
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
	termbox.Flush()
}
