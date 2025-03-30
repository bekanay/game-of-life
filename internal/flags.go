package internal

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func InitFlags() (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)
	result["delay-ms"] = 2500
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "--help":
			if _, ok := result["help"]; !ok {
				helpFlag()
				result["help"] = true
			}
		case "--verbose":
			if _, ok := result["verbose"]; !ok {
				result["verbose"] = true
			}
		case "--edges-portal":
			if _, ok := result["edges-portal"]; !ok {
				result["edges-portal"] = true
			}
		case "--fullscreen":
			if _, ok := result["fullscreen"]; !ok {
				result["fullscreen"] = true
			}
		case "--footprints":
			if _, ok := result["footprints"]; !ok {
				result["footprints"] = true
			}
		case "--colored":
			if _, ok := result["colored"]; !ok {
				result["colored"] = true
			}
		default:
			if len(arg) > 11 {
				if arg[:11] == "--delay-ms=" {
					if result["delay-ms"] == 2500 {
						val, err := strconv.Atoi(arg[11:])
						if err != nil {
							return nil, err
						}
						result["delay-ms"] = val
					}
					continue
				}
			}

			if len(arg) > (7) {
				if arg[:7] == "--file=" {
					_, ok := result["random"]
					_, ok1 := result["file"]

					if !ok && !ok1 {
						result["file"] = arg[7:]
						_, err := os.Open(arg[7:])
						if err != nil {
							return nil, err
						}
					}
					continue
				}
			}

			if len(arg) > (9) {
				if arg[:9] == "--random=" {
					_, ok := result["random"]
					_, ok1 := result["file"]
					if !ok && !ok1 {
						arr := make([]int, 0)
						num := ""
						for _, ch := range arg[9:] {
							if ch == 'x' {
								val1, err := strconv.Atoi(num)
								if val1 < 3 {
									return nil, errors.New("Specified value:" + strconv.Itoa(val1) + " is too low")
								}
								if err != nil {
									return nil, err
								}
								arr = append(arr, val1)
								num = ""
								continue
							}
							if !(ch >= 48 && ch <= 57) {
								return nil, errors.New("Incorrected random value")
							}
							num += string(ch)
						}
						val2, err := strconv.Atoi(num)
						if err != nil {
							return nil, err
						}
						if val2 < 3 {
							return nil, errors.New("Specified value:" + strconv.Itoa(val2) + " is too low")
						}
						arr = append(arr, val2)
						result["random"] = arr
					}
					continue
				}
			}
			return map[string]interface{}{}, errors.New("Non-existent flag is entered: " + arg)
		}
	}

	return result, nil
}

func helpFlag() {
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help        : Show the help message and exit")
	fmt.Println("  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println("  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println("  --file=X      : Load the initial grid from a specified file")
	fmt.Println("  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println("  --random=WxH  : Generate a random grid of the specified width (W) and height (H)")
	fmt.Println("  --fullscreen  : Adjust the grid to fit the terminal size with empty cells")
	fmt.Println("  --footprints  : Add traces of visited cells, displayed as '∘'")
	fmt.Println("  --colored     : Add color to live cells and traces if footprints are enabled")
	os.Exit(0)
}
