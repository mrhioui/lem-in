package lemin

type Room struct {
	Name      string
	X, Y      int
	Relations []Room
}

type Shema struct {
	NAnt  int
	Start Room
	End   Room
	Rooms []Room
}
