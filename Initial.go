package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name  string
	Score int
}

type ScoreBoard struct {
	Players []Player
}

func (s *ScoreBoard) AddPlayer(name string, score int) {
	s.Players = append(s.Players, Player{
		Name:  name,
		Score: score,
	})
}

func (s *ScoreBoard) Print() {
	sort.Slice(s.Players, func(i, j int) bool {
		return s.Players[i].Score > s.Players[j].Score
	})

	fmt.Println("Score Board")
	fmt.Println("===========")

	for index, player := range s.Players {
		fmt.Printf("%d. %s - %d\n", index+1, player.Name, player.Score)
	}
}

func main() {
	board := ScoreBoard{}

	board.AddPlayer("Alice", 950)
	board.AddPlayer("Brian", 870)
	board.AddPlayer("Clara", 1120)
	board.AddPlayer("David", 990)

	board.Print()
}
