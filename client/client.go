package main

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	width, height = 450, 800
	host, port    = "47.148.133.120", "4200"
)

var (
	document          *js.Object
	documentBody      *js.Object
	documentBodyStyle *js.Object
	alert             func(interface{})

	phaser   *js.Object
	game     *js.Object
	gameLoad *js.Object

	loading *js.Object
	loaded  bool
	// ws                *js.Object
	// t                 *js.Object
	// button            *js.Object
)

func main() {
	loadDOM()
	loadPhaser()
}

func loadDOM() {
	document = js.Global.Get("document")
	documentBody = document.Get("body")
	documentBodyStyle = documentBody.Get("style")
	alert = func(x interface{}) { js.Global.Call("alert", x) }

	documentBodyStyle.Set("background", "#000000")
	documentBodyStyle.Set("margin", 0)
}

func loadPhaser() {
	script := document.Call("createElement", "script")
	script.Set("src", "phaser.min.js")
	script.Set("onload", onPhaserLoad)
	documentBody.Call("appendChild", script)
}

func onPhaserLoad() {
	phaser = js.Global.Get("Phaser")
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
	gameLoad.Call("spritesheet", "loading", "assets/loading.png", width, height)
	gameLoad.Get("onLoadComplete").Call("add", func() { loaded = true })
}

func create() {
	add := game.Get("add")

	loading = add.Call("sprite", 0, 0, "loading")
	loading.Set("alpha", 0)
	fadeIn := add.Call("tween", loading).Call("to", js.M{"alpha": 1}, 1333)
	fadeIn.Set("frameBased", true)
	fadeIn.Call("start")

	taptostart := "assets/taptostart.png"

	spin := loading.Get("animations").Call("add", "spin")
	spin.Call("play", 9, true)
	spin.Get("onLoop").Call("add", func() {
		if !loaded {
			return
		}
		spin.Call("stop")
		loading.Set("visible", false)

		goBtn, goHit := newBtn(taptostart, 0, 0)
		go func() {
			<-goHit
			goBtn.Set("visible", false)
			start()
		}()
	})

	gameLoad.Call("spritesheet", taptostart, taptostart, width, height)

	gameLoad.Call("start")

	// button = add.Call("button", game.Get("world").Get("centerX"), game.Get("world").Get("centerY"), "button", func() {
	// 	ws.Call("send", "Hello!")
	// }, nil, 1, 0, 2)
	// button.Get("anchor").Call("setTo", 0.5, 0.5)

	// var text = "- phaser -\n with a sprinkle of \n pixi dust."
	// var style = js.M{"font": "65px Arial", "fill": "#ff0044", "align": "center"}
	// t = add.Call("text", 0, 0, text, style)

	// ws = js.Global.Get("WebSocket").New("ws://" + host + ":" + port + "/connected")
	// ws.Set("onopen", onConnectionOpen)
	// ws.Set("onclose", onConnectionClose)
	// ws.Set("onmessage", onConnectionMessage)
	// ws.Set("onerror", onConnectionError)
}

func newBtn(url string, x, y int) (*js.Object, <-chan bool) {
	hit := make(chan bool)
	btn := game.Get("add").Call("button", x, y, url, func() { hit <- true }, nil, 0, 0, 1, 0)
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
