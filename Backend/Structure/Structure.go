package structure

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

type Teacher struct {
	Name string `json:"name"`
}

type Attendence struct {
	AID     uint   `json:"aid"`cd Docu	
	cd 
	Day     string `json:"day"`
	Month   string `json:"month"`
	Year    string `json:"year"`
	Punchin  string `json:"punchin"`
	Puchout string `json:"puchout"`
	Type    string `json:"type"`
	Class   string `json:"class"`
}
