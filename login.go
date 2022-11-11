package main

import (
	"net/http"
	"unicode/utf8"
)

type LoginErrorMsg struct {
	UserNameMessage string
	PasswordMessage string
}

var loginErrorMsg LoginErrorMsg = LoginErrorMsg{}

/*
ログインボタン押下時の処理
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginErrorMsg = LoginErrorMsg{}
		ReturnPage(w, loginErrorMsg, "login")
	}
	if r.Method == "POST" {
		loginErrorMsg.UserNameMessage = ""
		loginErrorMsg.PasswordMessage = ""
		r.ParseForm()

		// ユーザーネーム入力チェック
		stringLen := utf8.RuneCountInString(r.Form.Get("userName"))
		if stringLen == 0 {
			loginErrorMsg.UserNameMessage = "ユーザー名が入力されていません"
		} else if stringLen > 10 {
			loginErrorMsg.UserNameMessage = "10文字以内で入力してください"
		}

		// パスワード入力チェック

		// 入力内容によってページ遷移
		if loginErrorMsg.UserNameMessage == "" && loginErrorMsg.PasswordMessage == "" {
			userInformation = UserInformation{
				UserName: r.Form.Get("userName"),
			}
			http.Redirect(w, r, "/search/", 301)
		} else {
			ReturnPage(w, loginErrorMsg, "login")
		}
	}
}
