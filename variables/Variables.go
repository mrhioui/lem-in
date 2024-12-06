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
	Paths  [][]Room
	Errors = map[string]string{
		"Args":       "ERROR: there are more or less arguments than expected",
		"Input":      "ERROR: unable to read file, no such file or directory",
		"Ants":       "ERROR: the number of ants is not correct or missing",
		"Repeat":     "ERROR: this line is repeated: ",
		"Room":       "ERROR: this cannot be a room: ",
		"XError":     "ERROR: the X of room not a number: ",
		"YError":     "ERROR: the Y of room not a number: ",
		"Tunnul":     "ERROR: invalid Tunnul format, expected format: 'RoomA-RoomB'",
		"TunnulRoom": "ERROR: one or both rooms do not exist:",
		"Invalid":    "ERROR: invalid data format",
	}
)
