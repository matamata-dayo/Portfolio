package main

import (
	"html/template"
	"net/http"
)

/*
menu画面のテンプレートを表示
*/
func menu(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/html/menu.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

/*
localhostの接続のみ受け付けるサーバー
*/
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", menu)
	http.ListenAndServe(":8080", nil)
}
