package models

type Status struct {
	Water uint `json:"water"`
	Wind  uint `json:"wind"`
}

var Data = []*Status{}

func Init() {
	Data = append(Data, &Status{Water: 6, Wind: 8})
}
