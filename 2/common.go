package main

type Gesture struct {
	Name   string
	Beats  []string
	Points int
}

var Rock = Gesture{
	Name:   "Rock",
	Beats:  []string{"Scissors", "C"},
	Points: 1,
}

var Paper = Gesture{
	Name:   "Paper",
	Beats:  []string{"Rock", "A"},
	Points: 2,
}

var Scissors = Gesture{
	Name:   "Scissors",
	Beats:  []string{"Paper", "B"},
	Points: 3,
}

const lossPoints = 0
const drawPoints = 3
const winPoints = 6

var aliases = map[string]Gesture{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var outcomes = map[string]int{
	"X": lossPoints,
	"Y": drawPoints,
	"Z": winPoints,
}

var gestures = [3]Gesture{Rock, Paper, Scissors}
