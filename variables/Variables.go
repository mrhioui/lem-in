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
	Paths  [][]*Room
	Errors = map[string]string{
		"Args":        "ERROR: There more argements not inportant OR no argement file.",
		"Input":       "ERROR: unable to read file: no such file or directory.",
		"AntsMissing": "ERROR: The number of ants is missing",
		"Ants":        "ERROR: The number of ants is not correct or equal '0'.",
		"Start-End" : "ERROR: The start or end flag is not associated with a room.",
		"Repeat":      "ERROR: This line is repeated : ",
		"Room":        "ERROR: Invalid Room format, expected format : 'Name x y'.",
		"XError":      "ERROR: The X of room not a number : ",
		"YError":      "ERROR: The Y of room not a number : ",
		"Tunnul":      "ERROR: Invalid Tunnul format, expected format : 'RoomA-RoomB'",
		"TunnulRoom":  "ERROR: One or both rooms do not exist : ",
		"Invalid!":     "ERROR: Invalid file for Lem-In",
	}
)
