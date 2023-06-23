package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./template/03-slice_nested_struct.gohtml"))
}

type exclusive_games struct {
	Game     string
	Platform string
}

func main() {
	LOL := exclusive_games{"league of legends", "Pc"}
	Dokkan := exclusive_games{"Dokkan battle", "Mobile"}
	FGO := exclusive_games{"Fate Grand Order", "Mobile"}
	Halo := exclusive_games{"Halo", "Pc&xbox"}
	Forza := exclusive_games{"Forza Horizon", "Pc&xbox"}
	exclusives := []exclusive_games{LOL, Dokkan, FGO, Halo, Forza}
	data := struct {
		Exclusive []exclusive_games
	}{exclusives}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
