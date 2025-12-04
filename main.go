package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL") // (eg. https://timefactory.io)
	api := os.Getenv("API_URL")           // (eg. https://demo.timefactory.io)

	f := frame.NewFrame(pathless, api)
	text := f.ToBytes("https://raw.githubusercontent.com/timefactoryio/pathless/main/README.md")
	f.README(text)
	f.Landing("the perpetual motion machine", "timefactoryio", "")
	f.BuildSlides("./slides")
	f.AddPath("./img")
	f.Serve()
	select {}
}
