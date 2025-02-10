package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal("Could not start window:", err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	message := "Hello world of WASM"
	for {
		switch e := window.Event().(type) {
		case app.ConfigEvent:
			size := e.Config.Size
			message = fmt.Sprint(size.X, size.Y)
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ctx := app.NewContext(&ops, e)

			title := material.H1(theme, "Hello world of WASM"+message)

			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			title.Alignment = text.Middle
			title.Layout(ctx)
			e.Frame(ctx.Ops)
		}
	}
}
