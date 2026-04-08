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
	f.Start(os.Getenv("PATHLESS"))
	select {}
}
