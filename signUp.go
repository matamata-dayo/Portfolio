package main

import (
	"net/http"
	"unicode/utf8"
)

type UserInformation struct {
	UserName string
}

type SignUpErrorMsg struct {
	UserNameMessage string
	PasswordMessage string
}

var userInformation UserInformation = UserInformation{}
var signUpErrorMsg SignUpErrorMsg = SignUpErrorMsg{}

/*
会員登録ボタン押下時の処理
*/
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		signUpErrorMsg = SignUpErrorMsg{}
		ReturnPage(w, signUpErrorMsg, "signUp")
	}
	if r.Method == "POST" {
		signUpErrorMsg.UserNameMessage = ""
		signUpErrorMsg.PasswordMessage = ""
		r.ParseForm()

		// ユーザーネーム入力チェック
		stringLen := utf8.RuneCountInString(r.Form.Get("userName"))
		if stringLen == 0 {
			signUpErrorMsg.UserNameMessage = "ユーザー名が入力されていません"
		} else if stringLen > 10 {
			signUpErrorMsg.UserNameMessage = "10文字以内で入力してください"
		}

		// パスワード入力チェック
		if len(r.Form.Get("password")) < 6 {
			signUpErrorMsg.PasswordMessage = "パスワードが短すぎます"
		}

		// 入力内容によってページ遷移
		if signUpErrorMsg.UserNameMessage == "" && signUpErrorMsg.PasswordMessage == "" {
			userInformation = UserInformation{
				UserName: r.Form.Get("userName"),
			}
			http.Redirect(w, r, "/search/", 301)
		} else {
			ReturnPage(w, signUpErrorMsg, "signUp")
		}
	}
}
