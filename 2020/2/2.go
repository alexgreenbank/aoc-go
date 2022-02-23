package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	p1 := 0
	p2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var min, max int
		var c byte
		var s string
		inp := scanner.Text()
		// 2-9 c: cccccccc
		n, err := fmt.Sscanf(inp, "%d-%d %c: %s", &min, &max, &c, &s)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("Didn't get 4 values")
		}
		cfreq := make(map[byte]int)
		for _, ch := range s {
			cfreq[byte(ch)]++
		}
		if cfreq[c] >= min && cfreq[c] <= max {
			p1++
		}
		if (s[min-1] == c) != (s[max-1] == c) {
			p2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: %d\n", p1)
	fmt.Printf("part2: %d\n", p2)
}
