package main

import (
	"flag"
	"fmt"
)

func main() {
	var numVar int
	var minutes int
	flag.IntVar(&numVar, "balls", 27, "The number of balls to use. Must be 27 <= and <= 127")
	flag.IntVar(&minutes, "minutes", -1, "Set to get clock state at the specified minute")
	flag.Parse()

	if numVar < 27 || numVar > 127 {
		fmt.Println("Can't have a value less than 27 or greater than 127")
		return
	}

	clock := CreateClock(numVar)
	if minutes > 0 {
		clock.ClockState(minutes)
	} else {
		clock.CycleDays()
	}

}
