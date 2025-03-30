# Crunch03 - A Cellular Automaton Simulator

## Overview
Crunch03 is a simple cellular automaton simulation inspired by Conway's Game of Life, where cells on a grid evolve over time based on their neighbors. This implementation allows for various configuration options such as random grid generation, custom cells, and the ability to load an initial grid from a file. It includes support for visual enhancements like colored cells and fullscreen mode.

## Features
- **Grid Size**: Customize the grid's width and height.
- **Random Grid Generation**: Generate a random grid with living and empty cells.
- **File Input**: Load the grid from a specified file.
- **Custom Cells**: Define your own symbols for living, empty, and footprint cells.
- **Edge Portals**: Wrap the grid edges, making cells that exit one side appear on the opposite side.
- **Verbose Mode**: Display detailed simulation information.
- **Animation Speed**: Control the delay between each tick of the simulation.
- **Fullscreen Mode**: Adjust the grid to fit the terminal size.

## Usage

### Command-Line Flags
Crunch03 supports several command-line options for configuring the simulation:

- `--help`: Display the help message.
- `--verbose`: Display detailed simulation information (grid size, live cells count, tick number).
- `--delay-ms=X`: Set the animation speed in milliseconds (default is 2500 ms).
- `--file=FILE`: Load the initial grid from the specified file.
- `--edges-portal`: Enable portal edges where cells exiting one edge appear on the opposite edge.
- `--random=WxH`: Generate a random grid with specified width (W) and height (H).
- `--fullscreen`: Adjust the grid to fit the terminal size.
- `--footprints`: Display traces of visited cells as '∘'.
- `--colored`: Add color to the live cells and footprints.
- `--custom-cells=XXX`: Define custom characters for living, empty, and footprint cells.

### Example Commands
```bash
go run main.go --random=20x10 --verbose --delay-ms=1000
go run main.go --file=grid.txt --fullscreen --colored
go run main.go --random=30x15 --edges-portal --footprints --custom-cells=██·
```
## Configuration File Format

If you choose to load the grid from a file, the format must adhere to the following rules:

- The first line should contain two integers: the grid's height and width, separated by a space.
- The subsequent lines should contain strings of `.` (empty) and `#` (living cells), corresponding to the grid's rows.


## How It Works

Crunch03 simulates cellular automaton behavior by updating the grid at each tick. Each cell's state is determined by the number of living neighbors around it, and cells evolve based on the following rules:

- A live cell with fewer than two live neighbors dies (underpopulation).
- A live cell with more than three live neighbors dies (overpopulation).
- A dead cell with exactly three live neighbors becomes alive (reproduction).
- All other cells remain in their current state.

## File Structure

- **`main.go`**: The entry point of the program, responsible for initializing and starting the simulation.
- **`game/`**: Contains logic related to initializing the game, handling grid updates, and printing the grid.
- **`internal/`**: Holds the configuration, grid setup, and helper functions.
- **`utils/`**: Provides utility functions like determining if a cell is alive.

