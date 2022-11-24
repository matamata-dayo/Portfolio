package database

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	Name     string
	Password string
}

var User UserInfo

func SetUserInfo(a string, b string) {
	User = UserInfo{
		Name:     a,
		Password: b,
	}
}

func AddUserInfo(db *sql.DB) {

	password := []byte(User.Password)

	// パスワードをハッシュ化
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)

	query := "INSERT INTO user_info(name, password) VALUES('" + User.Name + "', '" + string(hashed) + "')"

	_, err := db.QueryContext(context.Background(), query)

	if err != nil {
		fmt.Println("ユーザー情報の登録に失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("ユーザー情報の登録に成功しました")
	}

}
