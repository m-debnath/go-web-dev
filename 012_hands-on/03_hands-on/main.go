package main

import (
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Sea View Palace",
			Address: "Sea view street 16",
			City:    "Puri",
			Zip:     "610045",
			Region:  "Southern",
		},
		hotel{
			Name:    "IBIS Bhuvaneshwar",
			Address: "M K Gandhi Road",
			City:    "Bhuvaneshwar",
			Zip:     "600012",
			Region:  "Central",
		},
		hotel{
			Name:    "Radison Blu Rourkella",
			Address: "Pandit Nehru Road",
			City:    "Rourkella",
			Zip:     "611006",
			Region:  "Northern",
		},
		hotel{
			Name:    "Sea Hawk",
			Address: "Swargadwar Road",
			City:    "Puri",
			Zip:     "610001",
			Region:  "Southern",
		},
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
}
