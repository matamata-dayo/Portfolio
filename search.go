package main

import (
	"html/template"
	"net/http"
	"portfolio/database"
)

/*
検索画面遷移時の処理
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/html/search.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
	}
	if r.Method == "POST" {
		database.GetPlayerInfo(r, GetConnection())
		http.Redirect(w, r, "/searchResult/", 301)
	}
}
