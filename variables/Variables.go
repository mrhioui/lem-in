package variables

type Room struct {
	X, Y      int
	Relations []*Room
	Visited   bool
}

type Ant struct {
	Path        []string
	CurrentRoom int
	Previous    string
	IsInactive  bool
}

var (
	Rooms  = make(map[string]*Room)
	NAnt   int
	Start  string
	End    string
	Paths  [][]Room
	Errors = map[string]string{
		"Args":       "ERROR: There more argements not inportant OR no argement file",
		"Input":      "ERROR: unable to read file: no such file or directory",
		"Ants":       "ERROR: The number of ants is not correct or missing",
		"Repeat":     "ERROR: This line is repeated : ",
		"Room":       "ERROR: This is not a room : ",
		"XError":     "ERROR: The X of room not a number : ",
		"YError":     "ERROR: The Y of room not a number : ",
		"Tunnul":     "ERROR: Invalid Tunnul format, expected format : 'RoomA-RoomB'",
		"TunnulRoom": "ERROR: One or both rooms do not exist : ",
		"Invalid":    "ERROR: Invalid file for Lem-In",
	}
)
