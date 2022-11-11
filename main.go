package main

import (
	"html/template"
	"net/http"
)

/*
localhostの接続のみ受け付けるサーバー
*/
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", handleMenu)
	http.HandleFunc("/inquiry/", handleInquiry)
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
