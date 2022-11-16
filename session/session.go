package session

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

// セッション名
var session_name string = "session"

// Cookie型のstore情報
var store *sessions.CookieStore

// セッションオブジェクト
var session *sessions.Session

// セッションを開始させる
func SessionStart(w http.ResponseWriter, r *http.Request) {

	// セッション初期処理
	sessionInit()

	// セッションオブジェクトを取得
	session, _ := store.Get(r, session_name)

	// ログインしている状態
	session.Values["login"] = false

	// 保存
	err := session.Save(r, w)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("セッションを開始しました")
}

// セッションの存在チェック
func CheckSession() bool {
	// login, _ := session.Values["login"]
	return false
}

// セッション用の初期処理
func sessionInit() {

	// 乱数生成
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	str := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	// 新しいstoreを準備
	store = sessions.NewCookieStore([]byte(str))

	// セッションの有効範囲を指定
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/login/",
		MaxAge:   10,
		Secure:   false,
		HttpOnly: true,
	}
}
