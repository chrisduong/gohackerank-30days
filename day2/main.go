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
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// *END* of constraint block //

	// Read and save an integer, double, and String to your variables.
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Declare second integer, double, and String variables.
	i2, _ := strconv.ParseUint(lines[0], 10, 64)
	d2, _ := strconv.ParseFloat(lines[1], 64)
	s2 := lines[2]

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + i2)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+d2)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + s2)

}
