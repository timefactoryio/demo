package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL")
	api := os.Getenv("API_URL")

	f := frame.NewFrame(pathless, api)
	f.AddPath("./img")
	f.Home("the perpetual motion machine", "timefactoryio", "")
	f.Text(f.ToBytes("./README.md"))
	f.Slides("slides")
	f.Serve()
	select {}
}
