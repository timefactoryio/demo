package main

import (
	"github.com/timefactoryio/frame"
)

func main() {
	f := frame.NewFrame("timefactory.io", "demo.timefactory.io")

	f.Zero("the perpetual motion machine", "timefactoryio", "timefactoryio")
	f.README("./README.md", "")
	f.BuildSlides("./slides")
	f.AddPath("./img")
	f.Serve()
	select {}
}
