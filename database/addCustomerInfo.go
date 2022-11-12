package database

import (
	"context"
	"database/sql"
	"fmt"
)

type ContactInfo struct {
	Name      string
	Mail      string
	MailCheck string
	Content   string
}

var Contact ContactInfo

func SetContactInfo(a string, b string, c string, d string) {
	Contact = ContactInfo{
		Name:      a,
		Mail:      b,
		MailCheck: c,
		Content:   d,
	}
}

func AddCustomerInfo(db *sql.DB) {

	query := "INSERT INTO customer_info(name, mail, content) VALUES('" + Contact.Name + "', '" + Contact.Mail + "', '" + Contact.Content + "')"

	_, err := db.QueryContext(context.Background(), query)

	if err != nil {
		fmt.Println("お問い合わせ情報の登録に失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("お問い合わせ情報の登録に成功しました")
	}

}
