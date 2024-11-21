package lemin

type Room struct {
	X, Y      int
	Relations []*Room
	Visited   bool
}

type Shema struct {
	NAnt  int
	Start string
	End   string
	Paths [][]Room
}

var Rooms = make(map[string]Room)
