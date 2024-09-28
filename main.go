package main

import (
	"furqonProduct/Config"
	"furqonProduct/Controllers/categorycontroller"
	"furqonProduct/Controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	Config.ConnectDB()

	//1. homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
