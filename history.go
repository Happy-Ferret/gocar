package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

var history [][GAMEY][GAMEX]CellContains //slice of board
var lastPosition Game

func historyWriter() {
	for game.status != Ended {
		lastPosition = *game
		time.Sleep(time.Millisecond * 10)
		if lastPosition.board == game.board {
			continue
		}
		history = append(history, game.board)
	}
	go writeHistoryToJSON(history)
}

func writeHistoryToJSON(h [][GAMEY][GAMEX]CellContains) {
	HistoryJSON, _ := json.Marshal(h)
	ioutil.WriteFile("history.json", HistoryJSON, 0644)
}
