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
	mi1 := menuItem{
		"Breakfast",
		"Idli [2 Pieces] with Vada [1 Piece]",
		75,
	}
	mi2 := menuItem{
		"Breakfast",
		"5 Puri Bhaji",
		90,
	}
	mi3 := menuItem{
		"Breakfast",
		"Onion with Tomato Mixed Uthappam",
		100,
	}
	mi4 := menuItem{
		"Lunch",
		"Masala Dosa",
		90,
	}
	mi5 := menuItem{
		"Lunch",
		"Mysore Masala Dosa",
		110,
	}
	mi6 := menuItem{
		"Dinner",
		"Poori Chole",
		110,
	}
	mi7 := menuItem{
		"Breakfast",
		"Upma",
		45,
	}
	mi8 := menuItem{
		"Breakfast",
		"Vada",
		65,
	}
	menuItems := []menuItem{mi1, mi2, mi3, mi4, mi5, mi6, mi7, mi8}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", menuItems)
}
