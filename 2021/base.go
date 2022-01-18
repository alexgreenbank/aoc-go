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

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                var n int
                n, err = strconv.Atoi(scanner.Text())
                if err != nil {
                        log.Fatal(err)
                }
		// Do something with n
		fmt.Printf("Got [%d]\n", n)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

	fmt.Printf("Output here\n")
}
