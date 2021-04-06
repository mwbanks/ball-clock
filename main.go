package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type AC interface {
	ClockState(int) string
	CycleDays() int
	Init(int8)
}

func main() {
	var numVar int
	var minutes int
	var clockType int
	flag.IntVar(&numVar, "balls", 27, "The number of balls to use. Must be 27 <= and <= 127")
	flag.IntVar(&minutes, "minutes", -1, "Set to get clock state at the specified minute")
	flag.IntVar(&clockType, "clock-type", 0, "Set to use a different clock")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	flag.Parse()

	if numVar < 27 || numVar > 127 {
		fmt.Println("Can't have a value less than 27 or greater than 127")
		return
	}

	clock := CreateClock(numVar, clockType)
	if minutes > 0 {
		fmt.Println(clock.ClockState(minutes))
	} else {
		clock.CycleDays()
	}

}
