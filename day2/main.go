// https://www.hackerrank.com/challenges/30-data-types/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var i uint64 = 4
	// var d float64 = 4.0
	// var s string = "HackerRank "
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Declare second integer, double, and String variables.
	i2, _ := strconv.Atoi(lines[0])

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + uint64(i2))
}
