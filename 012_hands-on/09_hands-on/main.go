package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
)

type stock struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("Error opening file table.csv", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	stocks := []stock{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		stk := stock{
			Date:     record[0],
			Open:     record[1],
			High:     record[2],
			Low:      record[3],
			Close:    record[4],
			Volume:   record[5],
			AdjClose: record[6],
		}
		stocks = append(stocks, stk)
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", stocks)
}
