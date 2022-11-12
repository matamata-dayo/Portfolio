package database

import (
	"context"
	"database/sql"
	"fmt"
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

	query := "INSERT INTO user_info(name, password) VALUES('" + User.Name + "', '" + User.Password + "')"

	_, err := db.QueryContext(context.Background(), query)

	if err != nil {
		fmt.Println("ユーザー情報の登録に失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("ユーザー情報の登録に成功しました")
	}

}
