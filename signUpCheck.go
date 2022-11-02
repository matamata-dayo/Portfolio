package main

import (
	"html/template"
	"net/http"
)

type UserInformation struct {
	UserName string
}

type Msg struct {
	Message string
}

var userInformation UserInformation = UserInformation{}
var msg Msg = Msg{}

/*
会員登録ボタン押下時の処理
*/
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		returnMenu(w, msg)
	}
	if r.Method == "POST" {
		msg.Message = ""
		r.ParseForm()

		// 入力必須チェック
		if len(r.Form.Get("userName")) == 0 {
			msg.Message = msg.Message + "ニックネームが入力されていません"
		}

		// 入力内容によってページ遷移
		if msg.Message == "" {
			userInformation = UserInformation{
				UserName: r.Form.Get("userName"),
			}
			http.Redirect(w, r, "./assets/html/search.html", 301)
		} else {
			returnMenu(w, msg)
		}
	}
}

/*
メニュー画面に跳ね返す
*/
func returnMenu(w http.ResponseWriter, msg Msg) {
	t, err := template.ParseFiles("./assets/html/menu.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, msg); err != nil {
		panic(err.Error())
	}
}
