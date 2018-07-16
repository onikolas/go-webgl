package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

func main() {
	document := js.Global.Get("document")
	canvas := document.Call("getElementById", "cnvs")
	document.Get("body").Call("appendChild", canvas)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	gl.Viewport(0, 0, 512, 512)      // set 'window' size
	gl.ClearColor(0.9, 0.0, 0.00, 1) // background color
	gl.Clear(gl.COLOR_BUFFER_BIT)    // clear color buffer

}
