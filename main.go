package main

import (
	"os"

	"github.com/timefactoryio/frame"
)

func main() {
	f := frame.NewFrame()
	f.Home(os.Getenv("LOGO_PATH"), os.Getenv("HEADING"))
	f.Text(os.Getenv("TEXT_URL"))
	f.Slides(os.Getenv("SLIDES_PATH"))
	f.Start(os.Getenv("PATHLESS"))
	select {}
}
