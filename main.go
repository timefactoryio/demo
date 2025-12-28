package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL")
	api := os.Getenv("API_URL")

	f := frame.NewFrame(pathless, api, "./keyboard.html")
	f.Reader("./img")
	f.Home("the perpetual motion machine", "timefactoryio", "")
	f.Text("./README.md")
	f.Slides("slides")
	f.Start()
	select {}
}
