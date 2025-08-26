package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Control flow statements in Go are similar to other languages.

	// --- if/else ---

	// Basic if/else statement.
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// You can have an initialization statement in the if block.
	// The variable 'y' is only available within the scope of the if/else block.
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	} else {
		fmt.Println("y is not greater than 10")
	}

	// --- for ---

	// Go has only one looping construct, the 'for' loop.

	// 1. The classic for loop.
	for i := 0; i < 5; i++ {
		fmt.Println("Classic for loop:", i)
	}

	// 2. The "while" loop.
	// You can omit the initialization and post statements to create a "while" loop.
	j := 0
	for j < 5 {
		fmt.Println("While loop:", j)
		j++
	}

	// 3. The infinite loop.
	// You can omit the loop condition to create an infinite loop.
	// We use 'break' to exit the loop.
	k := 0
	for {
		fmt.Println("Infinite loop with break:", k)
		k++
		if k == 5 {
			break
		}
	}

	// --- switch ---

	// The switch statement is a shorter way to write a sequence of if/else statements.
	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, but Go only runs the selected case, not all the cases that follow.
	// In effect, the 'break' statement that is needed at the end of each case in those languages is provided automatically in Go.

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch with no condition.
	// This is a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
