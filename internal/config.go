package internal

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Config struct {
	width       int
	height      int
	random      bool
	verbose     bool
	delay       time.Duration
	edgePortals bool
	fullscreen  bool
	footprints  bool
	colored     bool
	flags       map[string]interface{}
	file        *os.File
}

func InitConfig(flags map[string]interface{}) (*Config, error) {
	var config Config
	for key, val := range flags {
		switch key {
		case "verbose":
			if v, ok := val.(bool); ok {
				config.verbose = v
			}
		case "edges-portal":
			if v, ok := val.(bool); ok {
				config.edgePortals = v
			}
		case "fullscreen":
			if v, ok := val.(bool); ok {
				config.fullscreen = v
			}
		case "footprints":
			if v, ok := val.(bool); ok {
				config.footprints = v
			}
		case "colored":
			if v, ok := val.(bool); ok {
				config.colored = v
			}
		case "delay-ms":
			if v, ok := val.(int); ok {
				config.delay = time.Millisecond * time.Duration(v)
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

				config.file, err = os.Open(v)
				if err != nil {
					return nil, err
				}
			}
		case "random":
			if values, ok := val.([]int); ok && len(values) == 2 {
				config.width = values[0]
				config.height = values[1]
				config.random = true
			}
		default:
			return nil, errors.New("Warning: Unknown flag " + key)
		}
	}
	return &config, nil
}
