package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL")
	api := os.Getenv("API_URL")

	f := frame.NewFrame(pathless, api)
	f.Landing("the perpetual motion machine", "timefactoryio", "")
	f.README(f.ToBytes("https://raw.githubusercontent.com/timefactoryio/pathless/main/README.md"))
	f.BuildSlides("./slides")
	f.AddPath("./img")
	f.Serve()
	select {}
}
