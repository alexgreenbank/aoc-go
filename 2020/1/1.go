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

	for k, _ := range vals {
		n := 2020 - k
		if _, ok := vals[n]; ok {
			if n > k {
				fmt.Printf("%d and %d sum to %d, product = %d\n", k, n, k+n, k*n)
			}
		}
	}
}
