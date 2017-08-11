package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mrmiguu/Golfer"
)

const (
	width, height = 450, 800
	host, port    = "localhost", "4200"
)

var (
	alert             func(interface{})
	document          *js.Object
	documentBody      *js.Object
	documentBodyStyle *js.Object
	window            *js.Object

	phaser    *js.Object
	game      *js.Object
	gameLoad  *js.Object
	gameAdd   *js.Object
	gameWorld *js.Object

	loading *js.Object
	// ws                *js.Object
	// t                 *js.Object
	// button            *js.Object
)

func main() {
	alert = func(x interface{}) { js.Global.Call("alert", x) }
	document = js.Global.Get("document")
	documentBody = document.Get("body")
	documentBodyStyle = documentBody.Get("style")
	window = js.Global.Get("window")

	documentBodyStyle.Set("background", "#000000")
	documentBodyStyle.Set("margin", 0)

	// load libraries
	<-golfer.Lib("https://unpkg.com/notie")
	<-golfer.Lib("lib/phaser.min.js")

	phaser = js.Global.Get("Phaser")
	// iW := window.Get("innerWidth").Float()
	// iH := window.Get("innerHeight").Float()
	// devPx := window.Get("devicePixelRatio").Float()
	game = phaser.Get("Game").New(width, height, phaser.Get("AUTO"), "phaser-example", js.M{"preload": preload, "create": create})

}

func preload() {
	game.Get("canvas").Set("oncontextmenu", func(e *js.Object) { e.Call("preventDefault") })
	scale := game.Get("scale")
	showAll := phaser.Get("ScaleManager").Get("SHOW_ALL")
	scale.Set("scaleMode", showAll)
	scale.Set("fullScreenScaleMode", showAll)
	scale.Set("pageAlignHorizontally", true)
	scale.Set("pageAlignVertically", true)

	gameLoad = game.Get("load")
	gameLoad.Call("spritesheet", "loading", "res/loading.png", width, height)
}

func create() {
	gameAdd = game.Get("add")
	gameWorld = game.Get("world")

	loading = newSprite("loading")
	loading.Set("alpha", 0)
	fadeIn := newTween(loading, js.M{"alpha": 1}, 1333)
	fadeIn.Call("start")

	taptostart := "res/taptostart.png"

	onLoad, loaded := golfer.Callback()
	gameLoad.Get("onLoadComplete").Call("add", onLoad)

	spin := loading.Get("animations").Call("add", "spin")
	spin.Call("play", 9, true)
	spin.Get("onLoop").Call("add", func() {
		select {
		case <-loaded:
			spin.Call("stop")
			loading.Set("visible", false)

			goBtn, goHit := newButton(taptostart)
			go func() {
				<-goHit
				goBtn.Set("visible", false)
				// msg := js.Global.Call("prompt", "Username")
				// alert(msg)

				field := document.Call("getElementById", "field")
				field.Call("focus")
				field.Call("click")

				// js.Global.Get("notie").Call("input", js.M{"text": "Username"})
				start()
			}()
		default:
			return
		}
	})

	gameLoad.Call("spritesheet", taptostart, taptostart, width, height)

	gameLoad.Call("start")

	// button = gameAdd.Call("button", game.Get("world").Get("centerX"), game.Get("world").Get("centerY"), "button", func() {
	// 	ws.Call("send", "Hello!")
	// }, nil, 1, 0, 2)
	// button.Get("anchor").Call("setTo", 0.5, 0.5)

	// var text = "- phaser -\n with a sprinkle of \n pixi dust."
	// var style = js.M{"font": "65px Arial", "fill": "#ff0044", "align": "center"}
	// t = gameAdd.Call("text", 0, 0, text, style)

	// ws = js.Global.Get("WebSocket").New("ws://" + host + ":" + port + "/connected")
	// ws.Set("onopen", onConnectionOpen)
	// ws.Set("onclose", onConnectionClose)
	// ws.Set("onmessage", onConnectionMessage)
	// ws.Set("onerror", onConnectionError)
}

func newTween(o *js.Object, params js.M, ms int) *js.Object {
	twn := gameAdd.Call("tween", o).Call("to", params, ms)
	twn.Set("frameBased", true)
	return twn
}

func newSprite(id string) *js.Object {
	spr := gameAdd.Call("sprite", gameWorld.Get("centerX"), gameWorld.Get("centerY"), id)
	spr.Get("anchor").Call("setTo", 0.5, 0.5)
	return spr
}

func newButton(url string) (*js.Object, <-chan bool) {
	x, y := gameWorld.Get("centerX").Int(), gameWorld.Get("centerY").Int()
	hit := make(chan bool)
	btn := game.Get("add").Call("button", x, y, url, func() { hit <- true }, nil, 0, 0, 1, 0)
	btn.Get("anchor").Call("setTo", 0.5, 0.5)
	btn.Get("onInputDown").Call("add", func() { btn.Set("y", y+min(height-btn.Get("height").Int(), 3)) })
	btn.Get("onInputOver").Call("add", func() { btn.Set("y", y) })
	btn.Get("onInputOut").Call("add", func() { btn.Set("y", y) })
	btn.Get("onInputUp").Call("add", func() { btn.Set("y", y) })
	return btn, hit
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// func onConnectionOpen(evt *js.Object) {
// 	print("Connected")
// }

// func onConnectionClose(evt *js.Object) {
// 	print("Disconnected")
// 	ws = nil
// }

// func onConnectionMessage(evt *js.Object) {
// 	print("Server: " + evt.Get("data").String())
// 	button.Call("setFrames", 4, 3, 5)
// }

// func onConnectionError(evt *js.Object) {
// 	print("Error: " + evt.Get("data").String())
// }

// func print(message interface{}) {
// 	t.Set("text", message)
// }

func start() {}
