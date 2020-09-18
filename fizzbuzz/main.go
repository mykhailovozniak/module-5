package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	FIZZ          ="fizz"
	BUZZ          ="buzz"
	MaxLowerBound =-1000
	MaxUpperBound =1000
	BASE          =10
	BitSize       =64
)

func parseInputArgs() (int64, int64) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: fizzbuzz [from - to] 1 15")
		os.Exit(1)
	}

	lowerBound := os.Args[1]
	upperBound := os.Args[2]

	from, lowerBoundErr := strconv.ParseInt(lowerBound, BASE, BitSize)
	to, upperBoundErr := strconv.ParseInt(upperBound, BASE, BitSize)

	if lowerBoundErr != nil || upperBoundErr != nil {
		fmt.Println("Usage: fizzbuzz [from - to] 1 15 from to should be a valid integers")
		os.Exit(1)
	}

	return from, to
}

func checkInputArgs(from int64 ,to int64) {
	if from > to {
		fmt.Println("Usage: fizzbuzz [from-to] 1 15 - from can not be bigger than to")
		os.Exit(1)
	}

	if from < MaxLowerBound {
		fmt.Printf("Usage: fizzbuzz [from-to] 1 15 - from can not be lower than %d" , MaxLowerBound)
		os.Exit(1)
	}

	if to > MaxUpperBound {
		fmt.Printf("Usage: fizzbuzz [from-to] 1 15 - to can not be bigger than %d" , MaxUpperBound)
		os.Exit(1)
	}
}

func printFIZZBUZZNumber(n int64) {
	if n % 3 == 0 {
		fmt.Printf(FIZZ)
	}

	if n % 5 == 0 {
		fmt.Printf(BUZZ)
	}

	if n % 3 != 0 && n % 5 != 0 {
		fmt.Printf("%d", n)
	}

	fmt.Printf("\n")
}

func printFIZZBUZZRange(from int64, to int64, print func(int64)) {
	for n := from; n <= to; n++ {
		print(n)
	}
}

func main() {
	from, to := parseInputArgs()
	checkInputArgs(from, to)
	printFIZZBUZZRange(from, to, printFIZZBUZZNumber)

	defer func() {
		fmt.Println("End of application...\nPerform some cleaning...")
	}()
}
