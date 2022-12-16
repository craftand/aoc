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

func main() {
	day2()
}

func day2() {

	totalScore := 0
	cheatScore := 0
	data, err := os.ReadFile("./day2.dat")

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
		cheatScore += rpsScore2(rhs) + rhs
	}

	fmt.Println("Total Score is:", totalScore)
	fmt.Println("Cheat score is:", cheatScore)
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

func rpsScore2(rhs int) int {
	if rhs == Rock {
		return Loss
	} else if rhs == Paper {
		return Draw
	} else if rhs == Scissors {
		return Win
	}

	return Loss
}
