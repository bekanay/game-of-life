package internal

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Width       int
	Height      int
	Random      bool
	Verbose     bool
	Delay       time.Duration
	EdgePortals bool
	Fullscreen  bool
	Footprints  bool
	Colored     bool
	flags       map[string]interface{}
	File        *os.File
	CustomCells []rune
}

func InitConfig(flags map[string]interface{}) (*Config, error) {
	var config Config
	for key, val := range flags {
		switch key {
		case "verbose":
			if v, ok := val.(bool); ok {
				config.Verbose = v
			}
		case "edges-portal":
			if v, ok := val.(bool); ok {
				config.EdgePortals = v
			}
		case "fullscreen":
			if v, ok := val.(bool); ok {
				config.Fullscreen = v
			}
		case "footprints":
			if v, ok := val.(bool); ok {
				config.Footprints = v
			}
		case "colored":
			if v, ok := val.(bool); ok {
				config.Colored = v
			}
		case "delay-ms":
			if v, ok := val.(int); ok {
				config.Delay = time.Millisecond * time.Duration(v)
			}
		case "file":
			if v, ok := val.(string); ok {
				fileInfo, err := os.Stat(v)
				if err != nil {
					return nil, err
				}

				if fileInfo.Size() == 0 {
					fmt.Println("The file is empty.")
				}

				config.File, err = os.Open(v)
				if err != nil {
					return nil, err
				}
			}
		case "random":
			if values, ok := val.([]int); ok && len(values) == 2 {
				config.Width = values[0]
				config.Height = values[1]
				config.Random = true
			}
		case "custom-cells":
			if values, ok := val.([]rune); ok {
				customCells := make([]rune, 0)
				customCells = append(customCells, values...)
				config.CustomCells = append(config.CustomCells, customCells...)
			}
		default:
			return nil, errors.New("Warning: Unknown flag " + key)
		}
	}
	return &config, nil
}
