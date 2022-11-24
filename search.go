package main

import (
	"fmt"
	"net/http"
	"portfolio/session"
)

/*
検索画面遷移時の処理
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		// ログインしていない状態からURL直打ちでアクセスした場合メニュー画面に遷移
		if session.CheckSession(w, r) {
			ReturnPage(w, "", "search")
		} else {
			fmt.Println("セッションが存在しないため、メニュー画面に戻します")
			ReturnPage(w, "", "menu")
		}

	} else if r.Method == "POST" {
		//
	}
}
