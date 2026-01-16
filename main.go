package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	f := frame.NewFrame()
	f.Home("/timefactory.svg", "the perpetual motion machine")
	f.Text("https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md")
	f.Slides("/slides")
	f.Start(os.Getenv("PATHLESS"))
	select {}
}
