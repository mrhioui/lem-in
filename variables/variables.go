package variables

type Room struct {
	X, Y      int
	Relations []*Room
	Visited   bool
}

var (
	Rooms  = make(map[string]*Room)
	NAnt   int
	Start  string
	End    string
	Paths  [][]string
	Errors = map[string]string{
		"Args":       "ERROR: There are more or less arguments than expected",
		"Input":      "ERROR: Unable to read file, no such file or directory",
		"Ants":       "ERROR: The number of ants is not correct or missing",
		"Repeat":     "ERROR: This line is repeated: ",
		"Room":       "ERROR: This cannot be a room: ",
		"XError":     "ERROR: The X of room not a number: ",
		"YError":     "ERROR: The Y of room not a number: ",
		"XYError":    "ERROR: Both rooms' X and Y coordinates are equal:",
		"Tunnul":     "ERROR: Invalid Tunnul format, expected format: 'RoomA-RoomB'",
		"TunnulRoom": "ERROR: One or both rooms do not exist:",
		"Invalid":    "ERROR: Invalid data format",
	}
)
