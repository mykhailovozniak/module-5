package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	fizz          ="fizz"
	buzz          ="buzz"
	maxLowerBound =-1000
	maxUpperBound =1000
	base          =10
	bitSize       =64
)

func parseInputArgs() (int64, int64) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: fizzbuzz [from - to] 1 15")
		os.Exit(1)
	}

	lowerBound := os.Args[1]
	upperBound := os.Args[2]

	from, lowerBoundErr := strconv.ParseInt(lowerBound, base, bitSize)
	if lowerBoundErr != nil {
		fmt.Println("Usage: fizzbuzz [from - to] 1 15 from should be a valid integer")
		os.Exit(1)
	}

	to, upperBoundErr := strconv.ParseInt(upperBound, base, bitSize)
	if upperBoundErr != nil {
		fmt.Println("Usage: fizzbuzz [from - to] 1 15 to to should be a valid integer")
		os.Exit(1)
	}

	return from, to
}

func checkInputArgs(from int64 ,to int64) {
	if from > to {
		fmt.Println("Usage: fizzbuzz [from-to] 1 15 - from can not be bigger than to")
		os.Exit(1)
	}

	if from < maxLowerBound {
		fmt.Printf("Usage: fizzbuzz [from-to] 1 15 - from can not be lower than %d" , maxLowerBound)
		os.Exit(1)
	}

	if to > maxUpperBound {
		fmt.Printf("Usage: fizzbuzz [from-to] 1 15 - to can not be bigger than %d" , maxUpperBound)
		os.Exit(1)
	}
}

func isFizz(n int64) bool {
	return n % 3 == 0
}

func isBuzz(n int64) bool {
	return n % 5 == 0
}

func printFIZZBUZZNumber(n int64) {
	if isFizz(n) {
		fmt.Printf(fizz)
	}

	if isBuzz(n) {
		fmt.Printf(buzz)
	}

	if !isFizz(n) && !isBuzz(n) {
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
