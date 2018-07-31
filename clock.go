package main

import (
	"fmt"
	"time"
)

type Clock struct {
	Min     BallQueue
	FiveMin BallQueue
	Hour    BallQueue
	HalfDay int
	Main    BallQueue
}

func (c *Clock) CycleDays() {
	start := time.Now()
	c.runMinute()
	for !c.Main.InOrder() {
		c.runMinute()
	}
	end := time.Now()
	fmt.Printf("%d balls cycle after %d days.\n", len(c.Main.Array), c.HalfDay/2)
	elapsed := end.Sub(start).Nanoseconds() / 1000000
	// elapsed := (end.UnixNano() - start.UnixNano()) / 1000
	floatElapsed := float64(elapsed)
	fmt.Printf("Completed in %d milliseconds (%f seconds)\n", elapsed, floatElapsed/1000)
}

func (c *Clock) ClockState(minutes int) {
	for i := 0; i < minutes; i++ {
		c.runMinute()
	}
	fmt.Printf("%s\n", c)
}

func (c *Clock) runMinute() {
	nextMin := c.Main.Pop(0)
	c.Min.Append(nextMin)
	if c.Min.IsFull() {
		lastVal := c.Min.Empty(&c.Main)

		c.FiveMin.Append(lastVal)
		if c.FiveMin.IsFull() {
			lastVal := c.FiveMin.Empty(&c.Main)

			c.Hour.Append(lastVal)
			if c.Hour.IsFull() {
				lastVal = c.Hour.Empty(&c.Main)

				c.Main.Append(lastVal)
				c.HalfDay++
			}
		}
	}

}

func (c Clock) String() string {
	return fmt.Sprintf("Min: %s, FiveMin: %s, Hour: %s, Main: %s", c.Min, c.FiveMin, c.Hour, c.Main)
}

func CreateClock(balls int) Clock {
	c := Clock{
		Min:     BallQueue{Size: 0, Array: make([]int, 5), Name: "Min"},
		FiveMin: BallQueue{Size: 0, Array: make([]int, 12), Name: "FiveMin"},
		Hour:    BallQueue{Size: 0, Array: make([]int, 12), Name: "Hour"},
		HalfDay: 0,
		Main:    BallQueue{Size: 0, Array: make([]int, balls), Name: "Main"},
	}
	for i := 0; i < balls; i++ {
		c.Main.Append(i + 1)
	}
	return c
}
