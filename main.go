package main

import (
	"log"
	"time"
	"github.com/sclevine/agouti"
)

var pageURL = "http://typingx0.net/sushida/play.html?soundless"

func main() {
	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatalf("driverの起動に失敗しました : %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("セッション作成に失敗しました : %v", err)
	}

	if err := page.Navigate(pageURL); err != nil {
		log.Fatalf("寿司打になにかあったかもしれません : %v", err)
	}

	// ロード待ちため7秒待つ
	time.Sleep(7 * time.Second)

	page.MoveMouseBy(550, 400) // WebGLを対象に実行するため、要素を取得できないので座標で対応する
	page.DoubleClick()
	time.Sleep(3 * time.Second) // アニメーション待ち


	page.MoveMouseBy(0, 20) // WebGLを対象に実行するため、要素を取得できないので座標で対応する
	page.DoubleClick()
	time.Sleep(3 * time.Second) // アニメーション待ち

	// キーストロークを送信する方法
	// https://github.com/sclevine/agouti/issues/61
	page.Session().Send("POST", "keys", map[string][]string{"value": {"\uE007"}}, nil)
	time.Sleep(2 * time.Second) // アニメーション待ち


	for i :=0; i < 2000; i++ {
		page.Session().Send("POST", "keys", map[string][]string{"value": {",-!?qwertyuiopasdfghjklzxcvbnm"}}, nil)
	}

	// スクショとる
	if err := page.Screenshot("sushi.png"); err != nil {
		log.Fatalf("スクショ取れまへん : %v", err)
	}
}
