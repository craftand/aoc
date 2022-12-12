package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day1()
}

func day1() {
	index := 0
	calories := make([]int, 0)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
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
}

func findCalorie(calories []int, index int) ([]int, int) {
	count := 0
	if len(calories) > 0 {
		count = calories[index]
	} else {
		calories = append(calories, 0)
	}

	return calories, count
}
