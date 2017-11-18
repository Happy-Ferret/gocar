package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"math/rand"
)

var (
	game *Game
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	playerEvent := make(chan termbox.Event)
	go func() {
		for {
			playerEvent <- termbox.PollEvent()
		}
	}()

	game = getNewGame()

	go game.timeBlock()

	for game.status != Ended {
		select {
		case event := <-playerEvent:
			if event.Type == termbox.EventKey {
				switch {
				case event.Ch == 's':
					game.setCarPosition(Down)
				case event.Ch == 'w':
					game.setCarPosition(Up)
				case event.Ch == 'a':
					game.setCarPosition(Left)
				case event.Ch == 'd':
					game.setCarPosition(Right)
				case event.Ch == 'o':
					return
				}
			}

		default:
			printGame(game)
		}
	}
}
