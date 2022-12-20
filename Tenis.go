package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayerStatus int64

const (
	Waiting PlayerStatus = iota
	Hitting
)

const (
	P int = 4
)

type TennisMatch struct {
	Player1     TennisPlayer
	Player2     TennisPlayer
	PointsToWin int
}

type TennisPlayer struct {
	Name   string
	Points int
}

func playerOne(player TennisPlayer, match chan TennisMatch, end chan bool) {
	for true {
		m := <-match

		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 10
		play := rand.Intn(max-min+1) + min

		if play <= 3 {
			fmt.Println(player.Name, "errou!")
			m.Player2.Points += 1
			printResult(m)
		} else {
			fmt.Println(player.Name, "acertou a bola!")
		}
		time.Sleep(time.Second * 1)

		if isFinished(m) {
			end <- true
		} else {
			match <- m
		}
	}
}

func playerTwo(player TennisPlayer, match chan TennisMatch, end chan bool) {
	for true {
		m := <-match

		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 10
		play := rand.Intn(max-min+1) + min

		if play <= 3 {
			fmt.Println(player.Name, "errou!")
			m.Player1.Points += 1
			printResult(m)
		} else {
			fmt.Println(player.Name, "acertou a bola!")

		}
		time.Sleep(time.Second * 1)

		if isFinished(m) {
			end <- true
		} else {
			match <- m
		}
	}
}

func printResult(m TennisMatch) {
	fmt.Println(m.Player1.Name, m.Player1.Points, "x", m.Player2.Points, m.Player2.Name)
}

func isFinished(match TennisMatch) bool {
	if match.Player1.Points == P || match.Player2.Points == P {
		return true
	} else {
		return false
	}
}

func main() {

	player1 := TennisPlayer{
		Name:   "Player 1",
		Points: 0,
	}

	player2 := TennisPlayer{
		Name:   "Player 2",
		Points: 0,
	}

	match := make(chan TennisMatch)
	go func() {
		match <- TennisMatch{
			Player1: player1,
			Player2: player2,
		}
	}()

	end := make(chan bool)

	go playerOne(player1, match, end)
	go playerTwo(player2, match, end)

	finished := false
	for !finished {
		finished = <-end
	}
}
