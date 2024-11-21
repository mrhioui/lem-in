package lemin

type Room struct {
	X, Y int
	// Parents   []Room
	Relations []*Room
	Visited   bool
}

type Shema struct {
	NAnt  int
	Start string
	End   string
	Paths [][]Room
}
