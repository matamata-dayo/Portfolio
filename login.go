package main

import (
	"net/http"
	"portfolio/database"
	"portfolio/session"
	"unicode/utf8"
)

type LoginErrorMsg struct {
	UserName        string
	UserNameMessage string
	Password        string
	PasswordMessage string
}

var loginErrorMsg LoginErrorMsg

/*
ログインボタン押下時の処理
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// テンプレートメッセージの初期化
		loginErrorMsg = LoginErrorMsg{}

		// ログインセッションが存在する場合は、ログイン画面を飛ばす
		if session.CheckSession() {
			http.Redirect(w, r, "/search/", 301)
		} else {
			ReturnPage(w, loginErrorMsg, "login")
		}
	} else if r.Method == "POST" {
		// テンプレートメッセージの初期化
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
				loginErrorMsg.PasswordMessage = "ユーザーネームかパスワードが違います"
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
				loginErrorMsg.PasswordMessage = "ユーザーネームかパスワードが違います"
			}
		}

		// 入力内容によってページ遷移
		if loginErrorMsg.UserNameMessage == "" && loginErrorMsg.PasswordMessage == "" {
			http.Redirect(w, r, "/search/", 301)
		} else {
			loginErrorMsg.UserName = userName
			loginErrorMsg.Password = password
			ReturnPage(w, loginErrorMsg, "login")
		}
	}
}
