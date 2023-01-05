package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Name string
}

type Hero struct {
	Name    string
	Emails  []string
	Friends []Friend
}

const HERO = `
Hero Name: {{.Name}}
{{range .Emails}}
Email: {{.}}
{{end}}
{{with .Friends}}
{{range .}}
Friend Name: {{.Name}}
{{end}}
{{end}}

`

func main() {
	f1 := Friend{"Thor"}
	f2 := Friend{"Hulk"}

	t := template.New("Hero")
	t, err := t.Parse(HERO)
	if err != nil {
		panic(err)
	}

	hero := Hero{
		Name:    "Ironman",
		Emails:  []string{"ironman@gmail.com", "ironman@ironman.com"},
		Friends: []Friend{f1, f2},
	}

	err = t.Execute(os.Stdout, hero)
	if err != nil {
		panic(err)
	}

}
