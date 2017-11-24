package main

import termbox "github.com/nsf/termbox-go"

var int2color = map[int]termbox.Attribute{
	0: termbox.ColorBlack,
	1: termbox.ColorBlue,
	2: termbox.ColorCyan,
	3: termbox.ColorGreen,
	4: termbox.ColorMagenta,
	5: termbox.ColorRed,
	6: termbox.ColorWhite,
	7: termbox.ColorYellow,
}

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

type Config struct {
	GAMEMARGINX      int
	GAMEMARGINY      int
	TEXTBLOCKMARGINY int
	TEXTBLOCKMARGINX int

	CARCOLOR            int
	NOTHINGCOLOR        int
	BLOCKCOLOR          int
	TEXTCOLOR           int
	TEXTBACKGROUNDCOLOR int
	GOLDCOLOR           int
}

func writeDefultJsonConfig() {

}
