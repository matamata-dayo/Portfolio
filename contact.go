package main

import (
	"net/http"
	"portfolio/database"
	"unicode/utf8"
)

type ContactErrorMsg struct {
	Name           string
	NameError      string
	Mail           string
	MailError      string
	MailCheck      string
	MailCheckError string
	Content        string
	ContentError   string
}

var contactErrorMsg ContactErrorMsg

/*
お問合せ画面のテンプレートを表示
*/
func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		contactErrorMsg = ContactErrorMsg{}
		ReturnPage(w, contactErrorMsg, "contact")
	}
	if r.Method == "POST" {

		// 初期化
		contactErrorMsg = ContactErrorMsg{}

		// POSTデータを格納
		r.ParseForm()
		database.SetContactInfo(r.FormValue("name"), r.FormValue("mail"), r.FormValue("mailCheck"), r.FormValue("content"))

		// 未入力チェック
		NameLen := utf8.RuneCountInString(database.Contact.Name)
		mailLen := utf8.RuneCountInString(database.Contact.Mail)
		mailCheckLen := utf8.RuneCountInString(database.Contact.MailCheck)
		ContentLen := utf8.RuneCountInString(database.Contact.Content)
		if NameLen == 0 {
			contactErrorMsg.NameError = contactErrorMsg.NameError + "お名前が入力されていません。"
		}
		if mailLen == 0 {
			contactErrorMsg.MailError = contactErrorMsg.MailError + "メールアドレスが入力されていません。"
		}
		if mailCheckLen == 0 {
			contactErrorMsg.MailCheckError = contactErrorMsg.MailCheckError + "メールアドレス確認が入力されていません。"
		}
		if ContentLen == 0 {
			contactErrorMsg.ContentError = contactErrorMsg.ContentError + "お問い合わせ内容が入力されていません。"
		}

		// メールアドレス一致チェック
		if contactErrorMsg.MailError == "" && database.Contact.Mail != database.Contact.MailCheck {
			contactErrorMsg.MailError = contactErrorMsg.MailError + "メールアドレスが一致していません。"
		}

		// エラーメッセージがない場合にページ遷移を行う
		if contactErrorMsg.NameError == "" &&
			contactErrorMsg.MailError == "" &&
			contactErrorMsg.MailCheckError == "" &&
			contactErrorMsg.ContentError == "" {

			// お問い合わせ内容をDBに格納
			database.AddCustomerInfo(GetConnection())

			// 結果画面を表示
			http.Redirect(w, r, "/contactResult/", 301)
		} else {
			contactErrorMsg.Name = r.FormValue("name")
			contactErrorMsg.Mail = r.FormValue("mail")
			contactErrorMsg.MailCheck = r.FormValue("mailCheck")
			contactErrorMsg.Content = r.FormValue("content")
			ReturnPage(w, contactErrorMsg, "contact")
		}
	}
}
