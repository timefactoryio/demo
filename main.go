package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL")
	api := os.Getenv("API_URL")

	f := frame.NewFrame(pathless, api)
	f.BuildFromHtml("./keyboard.html")
	f.Home("the perpetual motion machine", "timefactoryio", "")
	f.README(f.ToBytes("./README.md"))
	f.BuildSlides("slides")
	f.AddPath("./img")
	// BuildVideo(f, "./video/beef.mp4")
	// BuildVideo(f, "./video/hungary.mp4")
	f.Serve()
	select {}
}
