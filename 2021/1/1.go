package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need to supply filename as first arg")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cache [4]int

	idx := 0
	part1 := 0
	part2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n int
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		idx++
		// Use a circular queue to store the 4 most recent values
		cache[idx%4] = n
		// If we have at least one previous value we can compare it
		if idx > 1 && n > cache[(idx-1)%4] {
			part1++
		}
		// If we have at least three previous values we can compare
		// A+B+C > B+C+D
		// B+C are common to both sides so this can be simplified to:-
		// A > D
		if idx > 3 && n > cache[(idx-3)%4] {
			part2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}
