package main

import (
	"fmt"
	"syscall/js"
)

// documentオブジェクト取得用
var document = js.Global().Get("document")

// windowオブジェクトを取得
var window = js.Global()

// bodyのDOM取得
var body = document.Get("body")

// 対象のDOMのIDを取得する関数
func getElementByID(targetID string) js.Value {
	return document.Call("getElementById", targetID)
}

// 対象のDOMを作成する関数
func createElement(elementName string) js.Value {
	return document.Call("createElement", elementName)
}

func main() {
	// goからbuttonのDOMを作成する
	cLogBtn := createElement("button")
	// cLogBtnボタンのテキストを設定
	cLogBtn.Set("textContent", "console log!!")
	// buttonをbodyへ追加
	body.Call("appendChild", cLogBtn)
	// cLogBtnにclickのEventLisnerを設定
	cLogBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		fmt.Println("Hello Webassembly!")
		return nil
	}))

	// buttonDOMを作成する
	textChangeBtn := createElement("button")
	// textChangeBtnのテキストを設定
	textChangeBtn.Set("textContent", "text change!!")
	// buttonをbodyへ追加
	body.Call("appendChild", textChangeBtn)
	// textChangeにclick時のEventLisnerを設定
	textChangeBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		message := getElementByID("message")
		message.Set("innerHTML", "Hello, WebAssembry!!")
		return nil
	}))

	// buttonのDOMを生成
	alertBtn := createElement("button")
	// alertBtnのテキストを設定
	alertBtn.Set("textContent", "alert!!")
	// buttonをbodyへ追加
	body.Call("appendChild", alertBtn)
	// alertBtnにclick時のEventLisnerを設定
	alertBtn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		window.Call("alert", "Hello!!")
		return nil
	}))

	// 既にhtmlに書かれているボタンのidを取得して、goからclick時のEventListenerを設定する
	document.Call("getElementById", "init-text").Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		message := getElementByID("message")
		message.Set("innerHTML", "Hello, World")
		return nil
	}))

	// textエリアの入力値を取得
	getElementByID("in").Call("addEventListener", "keyup", js.FuncOf(func(js.Value, []js.Value) interface{} {
		getElementByID("out").Set("value", getElementByID("in").Get("value"))
		return nil
	}))

	// プログラムが終了しないように待機
	select {}
}
