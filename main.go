package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"portfolio/database"
	"portfolio/session"
)

var db *sql.DB

func main() {

	fmt.Println("*---プレミアリーグ選手検索を開始します---*")

	// コネクションを開き、アプリ終了時に閉じる
	db = database.Connect()
	defer db.Close()

	// セッションの初期化処理
	session.SessionInit()

	// ハンドラの登録
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", handleMenu)
	http.HandleFunc("/contact/", handleContact)
	http.HandleFunc("/contactResult/", handleContactResult)
	http.HandleFunc("/signUp/", handleSignUp)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/logout/", handleLogout)
	http.HandleFunc("/search/", handleSearch)
	http.HandleFunc("/searchResult/", handleSearchResult)
	http.ListenAndServe(":8080", nil)
}

/*
menu画面のテンプレートを表示
*/
func handleMenu(w http.ResponseWriter, r *http.Request) {
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
