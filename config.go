package main

import termbox "github.com/nsf/termbox-go"

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

//color and size
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
