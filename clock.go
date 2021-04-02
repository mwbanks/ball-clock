package main

import (
	"fmt"
	"time"
)

type Clock struct {
	Min     *BallQueue
	FiveMin *BallQueue
	Hour    *BallQueue
	HalfDay int
	Main    *BallQueue
}

type AClock struct {
	Min     *ABallQueue
	FiveMin *ABallQueue
	Hour    *ABallQueue
	HalfDay int
	Main    *ABallQueue
}

func (c *Clock) Init(balls int) {
	for i := 0; i < balls; i++ {
		c.Main.Append(i + 1)
	}
	// c.Main.Init(balls)
}

func (c *AClock) Init(balls int) {
	c.Main.Init(balls)
}

func (c *Clock) CycleDays() int {
	start := time.Now()
	fmt.Printf("Prerun Main: %#v\n", c.Main)
	c.runMinute()
	for !c.Main.InOrder() {
		c.runMinute()
		// days := c.HalfDay / 2
		// if days%500 == 0 {
		// 	fmt.Printf("%d days\r", days)
		// }
	}
	end := time.Now()
	fmt.Printf("Main: %#v\n", c.Main)
	fmt.Printf("Min: %#v\n", c.Min)
	fmt.Printf("FiveMin: %#v\n", c.FiveMin)
	fmt.Printf("Hour: %#v\n", c.Hour)
	fmt.Printf("%d balls cycle after %d days.\n", c.Main.MaxSize, c.HalfDay/2)
	elapsed := end.Sub(start).Nanoseconds() / 1000000
	// elapsed := (end.UnixNano() - start.UnixNano()) / 1000
	fmt.Printf("Completed in %d milliseconds (%f seconds)\n", elapsed, float64(elapsed)/1000)
	return c.HalfDay / 2
}

func (c *Clock) ClockState(minutes int) string {
	for i := 0; i < minutes; i++ {
		c.runMinute()
	}
	return fmt.Sprintf("%s", c)
}

func (c *Clock) runMinute() {
	nextMin := c.Main.Pop(0)
	c.Min.Append(nextMin)
	if c.Min.IsFull() {
		lastVal := c.Min.Empty(c.Main)

		c.FiveMin.Append(lastVal)
		if c.FiveMin.IsFull() {
			lastVal := c.FiveMin.Empty(c.Main)

			c.Hour.Append(lastVal)
			if c.Hour.IsFull() {
				lastVal = c.Hour.Empty(c.Main)

				c.Main.Append(lastVal)
				c.HalfDay++
			}
		}
	}

}

func (c Clock) String() string {
	return fmt.Sprintf(`{"Min": %s, "FiveMin": %s, "Hour": %s, "Main": %s`, c.Min, c.FiveMin, c.Hour, c.Main)
}

func (c *AClock) CycleDays() int {
	start := time.Now()
	fmt.Printf("Prerun Main: %#v\n", c.Main)
	c.runMinute()
	for !c.Main.InOrder() {
		c.runMinute()
	}
	end := time.Now()
	fmt.Printf("Main: %#v\n", c.Main)
	// fmt.Printf("Min: %#v\n", c.Min)
	fmt.Printf("FiveMin: %#v\n", c.FiveMin)
	fmt.Printf("Hour: %#v\n", c.Hour)
	fmt.Printf("%d balls cycle after %d days.\n", c.Main.MaxSize, c.HalfDay/2)
	elapsed := end.Sub(start).Nanoseconds() / 1000000
	// elapsed := (end.UnixNano() - start.UnixNano()) / 1000
	fmt.Printf("Completed in %d milliseconds (%f seconds)\n", elapsed, float64(elapsed)/1000)
	return c.HalfDay / 2
}

func (c *AClock) ClockState(minutes int) string {
	for i := 0; i < minutes; i++ {
		c.runMinute()
	}
	return fmt.Sprintf("%s\n", c)
}

func (c *AClock) runMinute() {
	nextMin := c.Main.PopFront()
	c.Min.Append(nextMin)
	if c.Min.IsFull() {
		lastVal := c.Min.Empty(c.Main)

		c.FiveMin.Append(lastVal)
		if c.FiveMin.IsFull() {
			lastVal := c.FiveMin.Empty(c.Main)

			c.Hour.Append(lastVal)
			if c.Hour.IsFull() {
				lastVal = c.Hour.Empty(c.Main)

				c.Main.Append(lastVal)
				c.HalfDay++
			}
		}
	}

}

func (c AClock) String() string {
	return fmt.Sprintf(`{"Min": %s, 
	"FiveMin": %s, 
	"Hour": %s, 
	"Main": %s
}`, c.Min, c.FiveMin, c.Hour, c.Main)
}

func CreateClock(balls int, clockType int) AC {
	arr := [154]int{}
	var c AC
	if clockType == 0 {
		c = &Clock{
			Min:     NewBallQueue(arr[:5], "Min", 5),
			FiveMin: NewBallQueue(arr[5:5+12], "FiveMin", 12),
			Hour:    NewBallQueue(arr[5+12:5+12+12], "Hour", 12),
			HalfDay: 0,
			Main:    NewBallQueue(arr[5+12+12:], "Main", balls),
		}

	} else {
		c = &AClock{
			Min:     NewABallQueue(arr[:5], "Min", 5),
			FiveMin: NewABallQueue(arr[5:5+12], "FiveMin", 12),
			Hour:    NewABallQueue(arr[5+12:5+12+12], "Hour", 12),
			HalfDay: 0,
			Main:    NewABallQueue(arr[5+12+12:], "Main", balls),
		}

	}
	c.Init(balls)
	return c
}
