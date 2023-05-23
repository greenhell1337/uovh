package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/mainpage.html")
	if err != nil{
		log.Fatal(err.Error())
	}

	tmpl.Execute(w, nil)
}

func FirstVideoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/firstvideo.html")
	if err != nil{
		log.Fatal(err.Error())
	}

	tmpl.Execute(w, nil)
}

func SecondVideoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/secondvideo.html")
	if err != nil{
		log.Fatal(err.Error())
	}

	tmpl.Execute(w, nil)
}

func ThirdVideoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/thirdvideo.html")
	if err != nil{
		log.Fatal(err.Error())
	}

	tmpl.Execute(w, nil)
}