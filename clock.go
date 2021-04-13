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
	FiveMin *A12Queue
	Hour    *A12Queue
	HalfDay int
	Main    *ABallQueue
}

func (c *Clock) Init(balls int8) {
	for i := 0; i < int(balls); i++ {
		c.Main.Append(i + 1)
	}
	// c.Main.Init(balls)
}

func (c *AClock) Init(balls int8) {
	c.Main.Init(sizenum(balls))
}

func (c *Clock) CycleDays() int {
	start := time.Now()
	fmt.Printf("Prerun Main: %s\n", c.Main)
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
	startup := true
	start := time.Now()
	for !c.Main.InOrder() || startup {
		// c.run5Min()
		// c.runHour()
		c.runHalfDay()
		// c.runHalfDay()
		startup = false
	}
	elapsed := time.Since(start).Nanoseconds() / 1000000
	// fmt.Printf("Main: %#v\n", c.Main)
	// fmt.Printf("Min: %#v\n", c.Min)
	// fmt.Printf("FiveMin: %#v\n", c.FiveMin)
	// fmt.Printf("Hour: %#v\n", c.Hour)
	fmt.Printf("%d balls cycle after %d days.\n", c.Main.MaxSize, c.HalfDay/2)
	fmt.Printf("Completed in %d milliseconds (%f seconds)\n", elapsed, float64(elapsed)/1000)
	return c.HalfDay / 2
}

func (c *AClock) ClockState(minutes int) string {
	for i := 0; i < minutes; i++ {
		c.run5Min()
	}
	return fmt.Sprintf("%s\n", c)
}

func (c *AClock) runHalfDay() {
	// var val ballnum
	var lastVal ballnum
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			val := c.Main.FastReverseReturn()
			c.FiveMin.Append(val)
		}
		lastVal = c.FiveMin.Empty(c.Main)

		c.Hour.Append(lastVal)
	}
	lastVal = c.Hour.Empty(c.Main)

	c.Main.Append(lastVal)
	c.HalfDay++

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			val := c.Main.FastReverseReturn()
			c.FiveMin.Append(val)
		}
		lastVal = c.FiveMin.Empty(c.Main)

		c.Hour.Append(lastVal)
	}
	lastVal = c.Hour.Empty(c.Main)

	c.Main.Append(lastVal)
	c.HalfDay++
}

func (c *AClock) runHour() {
	// var val ballnum
	for i := 0; i < 12; i++ {
		val := c.Main.FastReverseReturn()
		c.FiveMin.Append(val)
	}
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

func (c *AClock) run5Min() {
	// var val ballnum
	val := c.Main.FastReverseReturn()
	fmt.Println("val:", val)
	c.FiveMin.Append(val)
	if c.FiveMin.IsFull() {
		lastVal := c.FiveMin.Empty(c.Main)
		fmt.Println("lastVal:", lastVal)

		c.Hour.Append(lastVal)
		if c.Hour.IsFull() {
			lastVal = c.Hour.Empty(c.Main)
			fmt.Println("lastVal:", lastVal)

			c.Main.Append(lastVal)
			c.HalfDay++
		}
	}

}

// func (c *AClock) runMinute() {
// 	nextMin := c.Main.PopFront()
// 	c.Min.Append(nextMin)
// 	if c.Min.IsFull() {
// 		lastVal := c.Min.Empty(c.Main)

// 		c.FiveMin.Append(lastVal)
// 		if c.FiveMin.IsFull() {
// 			lastVal := c.FiveMin.Empty(c.Main)

// 			c.Hour.Append(lastVal)
// 			if c.Hour.IsFull() {
// 				lastVal = c.Hour.Empty(c.Main)

// 				c.Main.Append(lastVal)
// 				c.HalfDay++
// 			}
// 		}
// 	}

// }

func (c AClock) String() string {
	return fmt.Sprintf(`{
	"FiveMin": %s, 
	"Hour": %s, 
	"Main": %s
}`, c.FiveMin, c.Hour, c.Main)
}

func CreateClock(balls int, clockType int) AC {
	var c AC
	if clockType == 0 {
		arr := [155]int{}
		c = &Clock{
			Min:     NewBallQueue(arr[:5], "Min", 5),
			FiveMin: NewBallQueue(arr[5:5+12], "FiveMin", 12),
			Hour:    NewBallQueue(arr[5+12:5+12+12], "Hour", 12),
			HalfDay: 0,
			Main:    NewBallQueue(arr[5+12+12:], "Main", balls),
		}

	} else {
		arr := [280]ballnum{}
		// arr := make([]ballnum, 29+balls)
		c = &AClock{
			FiveMin: New12Queue(arr[:12], "FiveMin"),
			Hour:    New12Queue(arr[12:12+12], "Hour"),
			HalfDay: 0,
			Main:    NewABallQueue(arr[12+12:], "Main", sizenum(balls)),
		}

	}
	c.Init(int8(balls))
	return c
}
