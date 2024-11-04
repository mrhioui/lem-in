# Lem-in Project Overview

## 1. Understand the Requirements
- **Input Structure**: Read from a file describing rooms, tunnels, and the number of ants. The file includes specific sections (`##start`, `##end`, etc.) and comment lines.
- **Output Structure**: Display ant movements from `##start` to `##end`, showing only valid movements.
- **Error Handling**: Implement checks for various errors, such as missing start/end rooms, duplicate room names, invalid formats, etc.

## 2. Design Your Program
- **Data Structures**: 
  - **Rooms**: A struct that contains the name and coordinates.
  - **Tunnels**: A struct that connects two rooms.
  - **Colony**: A struct that contains all rooms and tunnels.
  - **Ants**: A struct that represents the ants and their current positions.

## 3. Parse Input File
- Read the file line by line.
- Identify sections (e.g., `##start`, `##end`) and store the corresponding data.
- Handle comments and skip them.

## 4. Validate Input
- Implement checks for:
  - Presence of `##start` and `##end`.
  - Duplicates in room names.
  - Valid links between rooms (i.e., they must exist).

## 5. Implement Pathfinding
- Use an algorithm like BFS or DFS to find the shortest path from `##start` to `##end`.
- Consider multiple paths since ants might need to avoid congestion.

## 6. Simulate Ant Movements
- Implement a simulation loop where you move the ants along the paths.
- Ensure that ants do not occupy the same room (except at start/end).
- Print the movements according to the specified format.

## 7. Error Handling
- Ensure your program doesn’t crash due to unexpected input.
- Use meaningful error messages when encountering invalid data.

## 8. Test Your Program
- Create multiple test files covering various scenarios (valid input, missing elements, malformed lines).
- Write unit tests if applicable to ensure each component behaves as expected.

## 9. Bonus: Visualizer
- If implementing the bonus, consider how to capture the output and display it visually.
- You could use ASCII art to represent the rooms and tunnels or integrate with a graphical library for a more sophisticated approach.

## Example Structure in Go
Here’s a basic outline of how your Go code structure might look:

```go
package main

import (
    "fmt"
    "os"
)

type Room struct {
    Name string
    X, Y int
}

type Tunnel struct {
    From, To *Room
}

type Colony struct {
    Rooms   map[string]*Room
    Tunnels []*Tunnel
}

type Ant struct {
    ID       int
    Position *Room
}

// Function prototypes
func parseInput(filename string) (*Colony, int, error) {
    // Implementation for parsing input
}

func validateData(colony *Colony) error {
    // Implementation for validation
}

func findPaths(colony *Colony) [][]*Room {
    // Implementation for pathfinding
}

func moveAnts(colony *Colony, paths [][]*Room, numAnts int) {
    // Implementation for simulating ant movements
}

func main() {
    // Read file, parse input, validate, find paths, simulate movements
}
