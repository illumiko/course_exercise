package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./template/02-slice(var).gohtml"))
}
func main() {
	languages := map[string]string{
		"js":     "webdev",
		"python": "a lot of fields",
		"java":   "legacy software + andirod apps",
		"Swift":  "ios app development",
	}
	err := tpl.Execute(os.Stdout, languages)
	if err != nil {
		log.Fatal(err)
	}
}
