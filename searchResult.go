package main

import (
	"fmt"
	"net/http"
	"portfolio/database"
	"portfolio/session"
)

/*
検索結果画面遷移時の処理
*/
func handleSearchResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// ログインしていない状態からURL直打ちでアクセスした場合メニュー画面に遷移
		if session.CheckSession(w, r) {
			ReturnPage(w, database.SearchResult, "searchResult")
		} else {
			fmt.Println("セッションが存在しないため、メニュー画面に戻します")
			ReturnPage(w, "", "menu")
		}

	} else if r.Method == "POST" {

		database.GetPlayerInfo(r, GetConnection())
		ReturnPage(w, database.SearchResult, "searchResult")

	}
}
