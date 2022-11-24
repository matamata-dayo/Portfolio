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

	// セッションオブジェクトを取得
	session, _ := store.Get(r, session_name)

	// 認証状態にする
	session.Values["auth"] = true

	// 保存
	err := session.Save(r, w)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("セッションを開始しました")
}

func SessionEnd(w http.ResponseWriter, r *http.Request) {

	// セッションオブジェクトを取得
	session, _ := store.Get(r, session_name)

	// 未認証状態にする
	session.Values["auth"] = false

	// 保存
	err := session.Save(r, w)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("セッションを終了しました")
}

// セッションの存在チェック
func CheckSession(w http.ResponseWriter, r *http.Request) bool {

	// セッションオブジェクトを取得
	session, _ := store.Get(r, session_name)

	login, _ := session.Values["auth"].(bool)

	return login
}

// セッション用の初期処理
func SessionInit() {

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
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
	}

	fmt.Println("セッション情報を初期化しました")
}
