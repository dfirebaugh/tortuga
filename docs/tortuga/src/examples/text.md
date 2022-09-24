# Text

<wasm-view height=400 width=530 src="font.wasm"></wasm-view>

```golang
package main

import "tortuga/pkg/tortuga"

var console tortuga.Console

type cart struct{}

func (c cart) Update() {}
func (c cart) Render() {
	console.PrintAt("hello, world!", 10, 100, 4)
}

func main() {
	console = tortuga.New()
	console.SetTitle("font example")
	console.SetScaleFactor(3)
	console.Run(cart{})
}
```
