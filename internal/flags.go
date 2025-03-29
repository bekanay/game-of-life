package internal

import (
	"flag"
	"fmt"
	"os"
)

func InitFlags() {
	helpFlag := flag.Bool("help", false, "Show the help message and exit")
	verboseFlag := flag.Bool("verbose", false, "Display detailed information about the simulation")
	delayMsFlag := flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 milliseconds")

	flag.Parse()
	if *helpFlag {
		fmt.Println("Usage: go run main.go [options]\n")
		fmt.Println("Options:")
		fmt.Println("  --help        : Show the help message and exit")
		fmt.Println("  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
		fmt.Println("  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds")
		os.Exit(0)
	}

	if *verboseFlag {
		fmt.Println("Verbose mode is ON. Displaying detailed simulation information...")
	}
	fmt.Printf("Animation speed is set to %d milliseconds.\n", *delayMsFlag)
}
