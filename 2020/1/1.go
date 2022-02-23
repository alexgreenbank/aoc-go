package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(nos map[int]bool) {
	for k, _ := range nos {
		n := 2020 - k
		if _, ok := nos[n]; ok {
			if n > k {
				fmt.Printf("%d and %d sum to %d, product = %d\n",
					k, n, k+n, k*n)
			}
		}
	}
}

func part2(nos map[int]bool) {
	for k1, _ := range nos {
		for k2, _ := range nos {
			n := 2020-k1-k2
			if _, ok := nos[n]; ok {
				if n > k2 && k2 > k1 {
					fmt.Printf("%d and %d and %d sum to %d, product = %d\n",
						k1, k2, n, k1+k2+n, k1*k2*n)
				}
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need to supply filename as first arg")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vals := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n int
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		vals[n] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(vals)

	part2(vals)
}
