package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock int = iota + 1
	Paper
	Scissors
)

const (
	Loss int = 0
	Draw int = 3
	Win  int = 6
)

// Y == draw, X == lose, Z == win
var table = map[string]int{
	"A Y": Draw + Rock,
	"B Y": Draw + Paper,
	"C Y": Draw + Scissors,

	"A X": Loss + Scissors,
	"B X": Loss + Rock,
	"C X": Loss + Paper,

	"A Z": Win + Paper,
	"B Z": Win + Scissors,
	"C Z": Win + Rock,
}

func str2rps(encoded string) int {
	switch encoded {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	return 0
}

func main() {
	day2()
}

func day2() {

	totalScore := 0
	cheatScore := 0
	data, err := os.ReadFile("./day2.txt")

	if nil != err {
		panic(err)
	}

	rps := string(data)

	for _, line := range strings.Split(rps, "\n") {
		hand := strings.Split(line, " ")
		lhs := str2rps(hand[0]) //opponent
		rhs := str2rps(hand[1])

		score := rpsScore(lhs, rhs)

		totalScore += score + rhs
		cscore := table[line]

		cheatScore += cscore
	}

	fmt.Println("Total Score is:", totalScore)
	fmt.Println("Cheat score is:", cheatScore)
}

func rpsScore(lhs, rhs int) int {
	if lhs == rhs {
		return Draw
	} else if rhs == Rock && lhs == Scissors {
		return Win
	} else if rhs == Scissors && lhs == Paper {
		return Win
	} else if rhs == Paper && lhs == Rock {
		return Win
	}

	return Loss
}
