package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestMainFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	fmt.Println("*---メニュー画面の接続確認---*")
	serverMenu := httptest.NewServer(http.HandlerFunc(handleMenu))
	defer serverMenu.Close()
	testServer(t, serverMenu)

	fmt.Println("*---お問い合わせ画面の接続確認---*")
	serverContact := httptest.NewServer(http.HandlerFunc(handleContact))
	defer serverContact.Close()
	testServer(t, serverContact)

	fmt.Println("*---お問い合わせ結果画面の接続確認---*")
	serverContactResult := httptest.NewServer(http.HandlerFunc(handleContactResult))
	defer serverContactResult.Close()
	testServer(t, serverContactResult)

	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}

func testServer(t *testing.T, s *httptest.Server) {
	res, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}
}
