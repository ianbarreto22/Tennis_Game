package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//P Number of points to a player win the game.
	P int = 4
)

// TennisMatch keeps the two players of the tennis game.
type TennisMatch struct {
	Player1 TennisPlayer
	Player2 TennisPlayer
}

// TennisPlayer keeps the name and points of a player
type TennisPlayer struct {
	Name   string
	Points int
}

// playerTwo simulates a shot of the first player, if it misses the shot, it will increase the opponent's points
func playerOne(player TennisPlayer, match chan TennisMatch, end chan bool) {
	for true {
		m := <-match

		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 10
		play := rand.Intn(max-min+1) + min

		if play <= 3 {
			fmt.Println(player.Name, "errou!")
			fmt.Println("Ponto do jogador:", m.Player2.Name+".")
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

// playerTwo simulates a shot o the second player, if it misses the shot, it will increase the opponent's points
func playerTwo(player TennisPlayer, match chan TennisMatch, end chan bool) {
	for true {
		m := <-match

		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 10
		play := rand.Intn(max-min+1) + min

		if play <= 3 {
			fmt.Println(player.Name, "errou!")
			fmt.Println("Ponto do jogador:", m.Player1.Name+".")
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

// printResult prints the result of a tennis match
func printResult(m TennisMatch) {
	fmt.Println()
	fmt.Println("Placar:")
	fmt.Println(m.Player1.Name, m.Player1.Points, "x", m.Player2.Points, m.Player2.Name)
	fmt.Println()
}

// isFinished returns true if one of the players have won the game, false otherwise
func isFinished(match TennisMatch) bool {
	pointsPlayer1, pointsPlayer2 := match.Player1.Points, match.Player2.Points
	if pointsPlayer1 >= P && pointsPlayer1 >= pointsPlayer2+2 {
		fmt.Println(match.Player1.Name, "venceu a partida!")
		return true
	} else if pointsPlayer2 >= P && pointsPlayer2 >= pointsPlayer1+2 {
		fmt.Println(match.Player2.Name, "venceu a partida!")
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

	fmt.Println("-------------------------------------------")
	fmt.Println()
	fmt.Println("Começou a partida!")
	fmt.Println()

	go playerOne(player1, match, end)
	go playerTwo(player2, match, end)

	finished := false
	for !finished {
		finished = <-end
	}

	fmt.Println()
	fmt.Println("Final da partida!")
	fmt.Println()
	fmt.Println("-------------------------------------------")

	close(match)
}
