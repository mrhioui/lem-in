package lemin

type Room struct {
	Name      string
	X, Y      int
	// Parents   []Room
	Relations []Room
	Visited   bool
}

type Shema struct {
	NAnt  int
	Start Room
	End   Room
	Paths [][]Room
}
