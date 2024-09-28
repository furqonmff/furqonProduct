package categorycontroller

import (
	"furqonProduct/models/categorymodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("Views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("Views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
