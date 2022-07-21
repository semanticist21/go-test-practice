package sort_example

import (
	"fmt"
	"sort"
)

type player struct {
	name  string
	age   int
	score int
	rate  float32
}

type players []player

func (p players) Less(i, j int) bool { return p[i].score < p[j].score }
func (p players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p players) Len() int           { return len(p) }

func ExecuteExample() {
	var pls players
	pls = addPlayer(pls, makePlayer("a", 13, 45, 78.4))
	pls = addPlayer(pls, makePlayer("c", 18, 54, 50.8))
	pls = addPlayer(pls, makePlayer("d", 16, 36, 89.7))
	pls = addPlayer(pls, makePlayer("b", 16, 24, 67.4))

	sort.Sort(sort.Reverse(pls))
	fmt.Println(pls)
}

func addPlayer(list players, newM player) players {
	return append(list, newM)
}

func makePlayer(name string, age int, score int, rate float32) player {
	var newPlayer player
	newPlayer.name = name
	newPlayer.age = age
	newPlayer.score = score
	newPlayer.rate = rate
	return newPlayer
}
