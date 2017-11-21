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

	game = rungame()

	for {
		if game.status != Ended {
			listenGame(game, playerEvent)
		} else {
			printGameEnded()
			select {
			case event := <-playerEvent:
				if event.Type == termbox.EventKey {
					if event.Key == termbox.KeyEsc {
						return
					}
					if event.Key == termbox.KeyEnter {
						game = rungame()
					}
				}
			}
		}
	}
}

func rungame() *Game {
	game := getNewGame()
	go game.doSteps()
	return game
}

func listenGame(game *Game, playerEvent chan termbox.Event) {
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
			case event.Ch == 'n':
				game.status = Ended
			case event.Ch == 'p':
				game.paused = !game.paused
			}
		}

	default:
		printGame(game)
	}
}
