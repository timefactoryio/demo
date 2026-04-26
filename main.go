package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	f := frame.NewFrame()
	f.Home(os.Getenv("LOGO"), os.Getenv("TITLE"))
	f.Text(os.Getenv("TEXT"))
	f.Slides(os.Getenv("SLIDES"))
	f.App(os.Getenv("APP_NAME"), os.Getenv("APP_URL"))
	f.Start(os.Getenv("PATHLESS"))
	select {}
}

// func main() {
// 	f := frame.NewFrame()
// 	f.Home("https://zero.s3.timefactory.io/timefactory.svg", "the perpetual motion machine")
// 	// f.Home("./timefactory.svg", "the perpetual motion machine")
// 	f.Text("https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md")
// 	f.Slides("./slides")
// 	f.App("traefik", "traefik.timefactory.io")
// 	f.Start(os.Getenv("PATHLESS"))
// 	select {}
// }
