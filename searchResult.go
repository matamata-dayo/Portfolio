package main

import (
	"html/template"
	"net/http"
	"portfolio/database"
	"portfolio/session"
)

/*
検索結果画面遷移時の処理
*/
func handleSearchResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		url := "./assets/html/searchResult.html"

		// ログインしていない状態からURL直打ちでアクセスした場合
		if !session.CheckSession(w, r) {
			url = "./assets/html/menu.html"
		}

		t, err := template.ParseFiles(url)
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, database.SearchResult); err != nil {
			panic(err.Error())
		}

	} else if r.Method == "POST" {
		//
	}
}
