package main

import (
	"html/template"
	"net/http"
	"portfolio/database"
	"portfolio/session"
)

/*
検索画面遷移時の処理
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		url := "./assets/html/search.html"

		// ログインしていない状態からURL直打ちでアクセスした場合
		if !session.CheckSession(w, r) {
			url = "./assets/html/menu.html"
		}

		t, err := template.ParseFiles(url)
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
	} else if r.Method == "POST" {

		database.GetPlayerInfo(r, GetConnection())
		ReturnPage(w, "", "searchResult")

	}
}
