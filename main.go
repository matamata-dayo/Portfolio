package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"portfolio/database"
	"portfolio/session"

	"golang.org/x/sync/errgroup"
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

	if err := run(context.Background()); err != nil {
		fmt.Println("サーバーの正常終了に失敗しました")
	}

}

/*
外部から終了を指示されたときにサーバーを終了するrun関数
*/
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":8080",
	}
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("サーバーのクローズに失敗しました")
			return err
		}
		return nil
	})

	// チャネルからの通知（終了通知）を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		fmt.Println("サーバーのシャットダウンに失敗しました")
	}
	return eg.Wait()
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
