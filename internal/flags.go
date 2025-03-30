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
			if _, exist := result["help"]; !exist {
				helpFlag()
			}
		case "--verbose":
			if _, exist := result["verbose"]; !exist {
				result["verbose"] = true
			}
		case "--edges-portal":
			if _, exist := result["edges-portal"]; !exist {
				result["edges-portal"] = true
			}
		case "--fullscreen":
			if _, exist := result["fullscreen"]; !exist {
				result["fullscreen"] = true
			}
		case "--footprints":
			if _, exist := result["footprints"]; !exist {
				result["footprints"] = true
			}
		case "--colored":
			if _, exist := result["colored"]; !exist {
				result["colored"] = true
			}
		default:
			if len(arg) > 15 {
				if arg[:15] == "--custom-cells=" {
					if _, exist := result["custom-cells"]; !exist {
						chars := arg[15:]
						if len(chars) != 3 {
							return map[string]interface{}{}, errors.New("incorrect number of chars, expected 3")
						}
						charsRunes := []rune(chars)

						result["custom-cells"] = charsRunes
					}
					continue
				}
			}

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

			if len(arg) > (9) {
				if arg[:9] == "--random=" {
					_, exist := result["random"]
					_, exist1 := result["file"]
					if !exist && !exist1 {
						arr := make([]int, 0)
						num := ""
						for _, ch := range arg[9:] {
							if ch == 'x' && num != "" {
								val1, err := strconv.Atoi(num)
								if err != nil {
									return nil, err
								}
								if val1 < 3 {
									return nil, errors.New("specified value:" + strconv.Itoa(val1) + " is too low")
								}

								arr = append(arr, val1)
								num = ""
								continue
							}
							if !(ch >= 48 && ch <= 57) {
								return nil, errors.New("incorrected random value")
							}
							num += string(ch)
						}
						val2, err := strconv.Atoi(num)
						if err != nil {
							return nil, err
						}
						if val2 < 3 {
							return nil, errors.New("specified value:" + strconv.Itoa(val2) + " is too low")
						}
						arr = append(arr, val2)
						result["random"] = arr
					}
					continue
				}
			}

			if len(arg) > (7) {
				if arg[:7] == "--file=" {
					_, exist := result["random"]
					_, exist1 := result["file"]

					if !exist && !exist1 {
						result["file"] = arg[7:]
						file, err := os.Open(arg[7:])
						if err != nil {
							return nil, err
						}
						defer file.Close()
					}
					continue
				}
			}

			return map[string]interface{}{}, errors.New("non-existent flag is entered: " + arg)
		}
	}

	return result, nil
}

func helpFlag() {
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help        	: Show the help message and exit")
	fmt.Println("  --verbose     	: Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println("  --delay-ms=X  	: Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println("  --file=X      	: Load the initial grid from a specified file")
	fmt.Println("  --edges-portal	: Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println("  --random=WxH  	: Generate a random grid of the specified width (W) and height (H)")
	fmt.Println("  --fullscreen  	: Adjust the grid to fit the terminal size with empty cells")
	fmt.Println("  --footprints  	: Add traces of visited cells, displayed as '∘'")
	fmt.Println("  --colored     	: Add color to live cells and traces if footprints are enabled")
	fmt.Println("  --custom-cells	: Add custom cells to living and empty cells")
	os.Exit(0)
}
