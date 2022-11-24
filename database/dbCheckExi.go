package database

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
ユーザーネームが既にDBに登録されているかの確認を行う
*/
func CheckUserName(db *sql.DB, userName string) bool {

	query := "SELECT COUNT(*) FROM user_info WHERE name = '" + userName + "'"

	rows, err := db.QueryContext(context.Background(), query)

	if err != nil {
		fmt.Println("ユーザー情報の問い合わせに失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("ユーザー情報の問い合わせに成功しました")
	}

	count := 0
	if rows.Next() {
		rows.Scan(&count)
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}

/*
ユーザーネームに対するパスワードの確認を行う
*/
func CheckPassword(db *sql.DB, userName string, password string) bool {

	query := "SELECT password FROM user_info WHERE name = '" + userName + "'"

	rows, err := db.QueryContext(context.Background(), query)

	if err != nil {
		fmt.Println("パスワードの問い合わせに失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("パスワードの問い合わせに成功しました")
	}

	pass := ""
	if rows.Next() {
		rows.Scan(&pass)
	}

	// ハッシュ化したパスワードとの照合
	compErr := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if compErr == nil {
		fmt.Println("パスワードの照合に成功しました")
		return true
	} else {
		fmt.Println("パスワードの照合に失敗しました")
		return false
	}
}
