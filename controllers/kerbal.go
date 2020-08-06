package controllers

import "fmt"

// KerbalItems is a list of avatar items
type KerbalItems struct {
	Color      string `json:"color"`
	Eyes       string `json:"eyes"`
	Mouth      string `json:"mouth"`
	Hair       string `json:"hair"`
	FacialHair string `json:"facial-hair"`
	Glasses    string `json:"glasses"`
	Suit       string `json:"suit"`
	Extras     string `json:"extras"`
}

// CreateKerbal takes a list of items and generates avatar
func CreateKerbal(items KerbalItems) {
	fmt.Println(items)
}
