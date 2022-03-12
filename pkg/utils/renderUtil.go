package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is for renerdering html files..
func RenderTemplate(w http.ResponseWriter, templ string) {

	parsedTempl, pe := template.ParseFiles("./templates/" + templ)
	if pe != nil {
		fmt.Println("parse error ", pe)
		return
	}

	err := parsedTempl.Execute(w, nil)
	if err != nil {
		fmt.Printf("error occured while parsing file : %s", err)
		return
	}
}
