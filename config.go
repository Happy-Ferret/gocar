package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/nsf/termbox-go"
)

//int to color
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
var (
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

//do it if err!=nil
func writeDefultJsonConfig(err error) {
	if err == nil {
		return
	}
	//default values
	config := &Config{2, 1, 1, 3, 3, 1, 5, 6, 0, 7}
	configJSON, _ := json.Marshal(config)
	ioutil.WriteFile("carconfig.json", configJSON, 0644)
}

func doConfig() {
	jsonFile, err := ioutil.ReadFile("carconfig.json")
	writeDefultJsonConfig(err)
	var conf Config
	err = json.Unmarshal(jsonFile, &conf)
	writeDefultJsonConfig(err)
	if err == nil {
		config2var(&conf)
	}
}

func config2var(conf *Config) {
	CARCOLOR = int2color[conf.CARCOLOR]
	NOTHINGCOLOR = int2color[conf.NOTHINGCOLOR]
	BLOCKCOLOR = int2color[conf.BLOCKCOLOR]
	TEXTCOLOR = int2color[conf.TEXTCOLOR]
	TEXTBACKGROUNDCOLOR = int2color[conf.TEXTBACKGROUNDCOLOR]
	GOLDCOLOR = int2color[conf.GOLDCOLOR]

	GAMEMARGINX = conf.GAMEMARGINX
	GAMEMARGINY = conf.GAMEMARGINY
	TEXTBLOCKMARGINY = conf.TEXTBLOCKMARGINY
	TEXTBLOCKMARGINX = conf.TEXTBLOCKMARGINX
}
