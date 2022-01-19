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

	var cache int

	idx := 0
	part1 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n int
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		idx++
		if idx > 1 && n > cache {
			part1++
		}
		cache = n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1: %d\n", part1)
}
