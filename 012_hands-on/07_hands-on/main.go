package main

import (
	"os"
	"text/template"
)

type menuItem struct {
	TimeOfDay, Name string
	Price           int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	imi1 := menuItem{
		"Breakfast",
		"Idli [2 Pieces] with Vada [1 Piece]",
		75,
	}
	imi2 := menuItem{
		"Breakfast",
		"5 Puri Bhaji",
		90,
	}
	imi3 := menuItem{
		"Breakfast",
		"Onion with Tomato Mixed Uthappam",
		100,
	}
	imi4 := menuItem{
		"Lunch",
		"Masala Dosa",
		90,
	}
	imi5 := menuItem{
		"Lunch",
		"Mysore Masala Dosa",
		110,
	}
	imi6 := menuItem{
		"Dinner",
		"Poori Chole",
		110,
	}
	imi7 := menuItem{
		"Breakfast",
		"Upma",
		45,
	}
	imi8 := menuItem{
		"Breakfast",
		"Vada",
		65,
	}
	bmi1 := menuItem{
		"Breakfast",
		"Superfit Paneer Salad[235 Kcal]",
		167,
	}
	bmi2 := menuItem{
		"Lunch",
		"Murg Dum Biriyani",
		127,
	}
	bmi3 := menuItem{
		"Dinner",
		"Chicken Masala Meal",
		177,
	}
	iMenuItems := []menuItem{imi1, imi2, imi3, imi4, imi5, imi6, imi7, imi8}
	bMenuItems := []menuItem{bmi1, bmi2, bmi3}
	menus := map[string]([]menuItem){
		"Idlicious": iMenuItems,
		"Box-8":     bMenuItems,
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", menus)
}
