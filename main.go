package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS")
	api := os.Getenv("FRAME")

	f := frame.NewFrame(pathless, api)
	f.Landing("the perpetual motion machine", "timefactoryio", "")
	f.README("./README.md")
	f.BuildSlides("./slides")
	f.AddPath("./img")
	f.Serve()
	select {}
}
