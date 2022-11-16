package main

import (
	"net/http"
	"portfolio/database"
	"unicode/utf8"
)

type SignUpErrorMsg struct {
	UserName        string
	UserNameMessage string
	Password        string
	PasswordMessage string
}

var signUpErrorMsg SignUpErrorMsg

/*
会員登録ボタン押下時の処理
*/
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		signUpErrorMsg = SignUpErrorMsg{}
		ReturnPage(w, signUpErrorMsg, "signUp")
	}
	if r.Method == "POST" {

		signUpErrorMsg = SignUpErrorMsg{}

		// POSTデータの登録
		r.ParseForm()
		database.SetUserInfo(r.Form.Get("userName"), r.Form.Get("password"))

		// ユーザーネーム入力チェック
		stringLen := utf8.RuneCountInString(database.User.Name)
		if stringLen == 0 {
			signUpErrorMsg.UserNameMessage = "ユーザー名が入力されていません"
		} else if stringLen > 10 {
			signUpErrorMsg.UserNameMessage = "10文字以内で入力してください"
		} else {
			// ユーザーネーム既登録チェック
			if !database.CheckUserName(GetConnection(), database.User.Name) {
				signUpErrorMsg.UserNameMessage = "ユーザーネームが既に使用されています"
			}
		}

		// パスワード入力チェック
		if len(database.User.Password) < 6 {
			signUpErrorMsg.PasswordMessage = "パスワードが短すぎます"
		}

		// エラーがない場合に処理を実行
		if signUpErrorMsg.UserNameMessage == "" && signUpErrorMsg.PasswordMessage == "" {

			// ユーザー情報をDBに登録
			database.AddUserInfo(GetConnection())

			// ページ遷移
			http.Redirect(w, r, "/search/", 301)
		} else {
			signUpErrorMsg.UserName = r.Form.Get("userName")
			signUpErrorMsg.Password = r.Form.Get("password")
			ReturnPage(w, signUpErrorMsg, "signUp")
		}
	}
}
