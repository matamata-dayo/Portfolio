package main

import (
	"html/template"
	"net/http"
)

/*
検索結果画面遷移時の処理
*/
func handleContactResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/html/contactResult.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, ""); err != nil {
			panic(err.Error())
		}
	}
	if r.Method == "POST" {
		//
	}
}
