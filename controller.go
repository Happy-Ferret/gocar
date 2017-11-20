package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
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
	termbox.HideCursor()
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
				case event.Ch == 's' || event.Key == termbox.KeyArrowDown:
					game.setCarPosition(Down)
				case event.Ch == 'w' || event.Key == termbox.KeyArrowUp:
					game.setCarPosition(Up)
				case event.Ch == 'a' || event.Key == termbox.KeyArrowLeft:
					game.setCarPosition(Left)
				case event.Ch == 'd' || event.Key == termbox.KeyArrowRight:
					game.setCarPosition(Right)
				case event.Ch == 'o':
					return
				case event.Ch == 'n':
					game.paused = !game.paused
				}
			}

		default:
			printGame(game)
		}
	}
}
