package main

import (
	"net/http"
	"portfolio/database"
	"unicode/utf8"
)

type LoginErrorMsg struct {
	UserNameMessage string
	PasswordMessage string
}

var loginErrorMsg LoginErrorMsg

/*
ログインボタン押下時の処理
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginErrorMsg = LoginErrorMsg{}
		ReturnPage(w, loginErrorMsg, "login")
	}
	if r.Method == "POST" {
		loginErrorMsg = LoginErrorMsg{}

		// POSTデータの取得
		r.ParseForm()
		userName := r.Form.Get("userName")
		password := r.Form.Get("password")

		// ユーザーネーム入力チェック
		userNameLen := utf8.RuneCountInString(userName)
		if userNameLen == 0 {
			loginErrorMsg.UserNameMessage = "ユーザー名が入力されていません"
		} else if userNameLen > 10 {
			loginErrorMsg.UserNameMessage = "10文字以内で入力してください"
		} else {
			// ユーザーネーム存在チェック
			if database.CheckUserName(GetConnection(), userName) {
				loginErrorMsg.UserNameMessage = "そのユーザーネームは登録されていません"
			}
		}

		// パスワード入力チェック
		passwordLen := utf8.RuneCountInString(password)
		if passwordLen == 0 {
			loginErrorMsg.PasswordMessage = "パスワードが入力されていません"
		} else if passwordLen < 6 {
			loginErrorMsg.PasswordMessage = "パスワードは6文字以上です"
		} else {
			if !database.CheckPassword(GetConnection(), userName, password) {
				loginErrorMsg.PasswordMessage = "パスワードが違います"
			}
		}

		// 入力内容によってページ遷移
		if loginErrorMsg.UserNameMessage == "" && loginErrorMsg.PasswordMessage == "" {
			http.Redirect(w, r, "/search/", 301)
		} else {
			ReturnPage(w, loginErrorMsg, "login")
		}
	}
}
