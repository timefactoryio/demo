package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	logo := os.Getenv("logo")
	heading := os.Getenv("heading")
	text := os.Getenv("text")
	slides := os.Getenv("slides")

	f := frame.NewFrame()
	f.Home(logo, heading)
	f.Text(text)
	f.Slides(slides)
	f.Start("")
	select {}
}
