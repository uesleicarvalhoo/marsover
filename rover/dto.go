package rover

type Params struct {
	Name        string      `json:"name"`
	Coordinates Coordinates `json:"coordinates"`
	Direction   Direction   `json:"direction"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}
