package lemin

type Room struct {
	Name string
	X, Y int
}

type Tunnel struct {
	From Room
	To   Room
}

type Shema struct {
	NAnt    int
	Start   Room
	End     Room
	Rooms   []Room
	Tunnuls []Tunnel
}
