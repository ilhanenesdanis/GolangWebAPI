package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	PageDescription string
	Header          string
	Content         string
	URI             string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var builder bytes.Buffer

	builder.WriteString("Golang dersleri\n")
	builder.WriteString("İlhan enes daniş")

	uri := "https://github.com/ilhanenesdanis"

	page := Page{
		Title:           "Golang Dersleri",
		Author:          "İlhan Enes Daniş",
		PageDescription: "GolangDersleri",
		Header:          "İlhan Enes Daniş Golang dersleri",
		Content:         builder.String(),
		URI:             uri,
	}
	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
