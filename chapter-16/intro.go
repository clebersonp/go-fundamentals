package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	dataMap := make(map[string]string)
	dataMap["name"] = "My new"
	dataMap["birthday"] = "1980-10-22"
	dataMap["age"] = "38"
	dataMap["first_name"] = "Mark"
	dataMap["last_name"] = "Down"

	dataSlice := make([]string, 0)
	dataSlice = append(dataSlice, "John")
	dataSlice = append(dataSlice, "1985-05-18")
	dataSlice = append(dataSlice, "42")
	dataSlice = append(dataSlice, "Dude")

	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
	executeTemplate("Dot is: {{.}}!\n", dataMap)
	executeTemplate("Dot is: {{.name}}!\n", dataMap)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is false!{{end}} finish\n", false)
	executeTemplate("Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n", dataMap)
	executeTemplate("Before loop: {{.}}\n{{range $k, $v := .}}In loop: '{{$k}}': {{$v}}\n{{end}}After loop: {{.}}\n", dataMap)
	executeTemplate("Before loop: {{.}}\n{{range $i, $v := .}}In loop: '{{$i}}': {{$v}}\n{{end}}After loop: {{.}}\n", dataSlice)
}
