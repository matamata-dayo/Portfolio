package main

import (
	"net/http"
	"unicode/utf8"
)

type InruiryErrorMsg struct {
	NameError      string
	MailError      string
	MailCheckError string
	QuestionError  string
}

var inruiryFormMessage InruiryErrorMsg = InruiryErrorMsg{}

/*
お問合せ画面のテンプレートを表示
*/
func handleInquiry(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		inruiryFormMessage = InruiryErrorMsg{}
		ReturnPage(w, inruiryFormMessage, "inquiry")
	}
	if r.Method == "POST" {
		inruiryFormMessage = InruiryErrorMsg{}
		r.ParseForm()

		// 未入力チェック
		nameLen := utf8.RuneCountInString(r.Form.Get("name"))
		mailLen := utf8.RuneCountInString(r.Form.Get("mail"))
		mailCheckLen := utf8.RuneCountInString(r.Form.Get("mailCheck"))
		questionLen := utf8.RuneCountInString(r.Form.Get("question"))
		if nameLen == 0 {
			inruiryFormMessage.NameError = inruiryFormMessage.NameError + "「お名前」が入力されていません。"
		} else if mailLen == 0 {
			inruiryFormMessage.MailError = inruiryFormMessage.MailError + "「メールアドレス」が入力されていません。"
		} else if mailCheckLen == 0 {
			inruiryFormMessage.MailCheckError = inruiryFormMessage.MailCheckError + "「メールアドレス確認」が入力されていません。"
		} else if questionLen == 0 {
			inruiryFormMessage.QuestionError = inruiryFormMessage.QuestionError + "「お問い合わせ内容」が入力されていません。"
		}

		// メールアドレス一致チェック
		if r.Form.Get("mail") != r.Form.Get("mailCheck") {
			inruiryFormMessage.MailError = inruiryFormMessage.MailError + "メールアドレスが一致していません。"
		}

		// 入力内容によってページ遷移
		if inruiryFormMessage.NameError == "" &&
			inruiryFormMessage.MailError == "" &&
			inruiryFormMessage.MailCheckError == "" &&
			inruiryFormMessage.QuestionError == "" {
			http.Redirect(w, r, "/inquiryResult/", 301)
		} else {
			ReturnPage(w, inruiryFormMessage, "inquiry")
		}
	}
}
