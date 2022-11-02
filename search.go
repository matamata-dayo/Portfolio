package main

import (
	"html/template"
	"net/http"
)

/*
会員登録ボタン押下時の処理
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/html/search.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, signUpErrorMsg); err != nil {
			panic(err.Error())
		}
	}
	if r.Method == "POST" {
		//
	}
}
