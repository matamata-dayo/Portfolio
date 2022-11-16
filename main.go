package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"portfolio/database"
	"portfolio/session"
)

var db *sql.DB

func main() {

	// コネクションを開き、アプリ終了時に閉じる
	db = database.Connect()
	defer db.Close()

	// ハンドラの登録
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", handleMenu)
	http.HandleFunc("/contact/", handleContact)
	http.HandleFunc("/contactResult/", handleContactResult)
	http.HandleFunc("/signUp/", handleSignUp)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/search/", handleSearch)
	http.HandleFunc("/searchResult/", handleSearchResult)
	http.ListenAndServe(":8080", nil)
}

/*
menu画面のテンプレートを表示
*/
func handleMenu(w http.ResponseWriter, r *http.Request) {

	// セッションの開始
	session.SessionStart(w, r)

	ReturnPage(w, "", "menu")
}

/*
各画面を表示
*/
func ReturnPage(w http.ResponseWriter, templateMsg interface{}, fileName string) {
	t, err := template.ParseFiles("./assets/html/" + fileName + ".html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, templateMsg); err != nil {
		panic(err.Error())
	}
}

func GetConnection() *sql.DB {
	return db
}
