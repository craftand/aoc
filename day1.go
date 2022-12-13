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

		sort.So
		fmt.Println("Elf with max calories is:", max)
	}
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
