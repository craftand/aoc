package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day1()
}

func day1() {

	data, err := os.ReadFile("./day1.dat")

	if nil != err {
		panic(err)
	}

	index := 0
	calories := make([]int, 0)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		fmt.Println("Got line", line)

		var count int
		if len(line) > 0 {
			calories, count = findCalorie(calories, index)
			lineCount, err := strconv.Atoi(line)

			fmt.Println("Line count", lineCount, " for index", index)

			if nil != err {
				fmt.Println("Unable to convert ", line, " to an integer.")
				continue
			}

			calories[index] = lineCount + count

			fmt.Println("Total count", calories[index], " for index", index)
		} else {
			fmt.Println("Elf at index", index, "has a total of", calories[index], "calories")
			index++
		}
	}

	len := len(calories)
	sort.Ints(calories)
	fmt.Println("Day 1: Max calories is:", calories[len-1])

	top3 := calories[len-3:]
	sum := 0
	for _, i := range top3 {
		sum += i
	}
	fmt.Println("Day 2: Top 3 calories:", sum)
}

func findCalorie(calories []int, index int) ([]int, int) {
	count := 0
	len := len(calories)
	if len > 0 && (len-1) == index {
		count = calories[index]
	} else {
		calories = append(calories, 0)
	}

	return calories, count
}
