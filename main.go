package main

import (
	"github.com/timefactoryio/frame"
)

// func main() {
// 	f := frame.NewFrame()
// 	f.Home(os.Getenv("LOGO_PATH"), os.Getenv("HEADING"))
// 	f.Text(os.Getenv("TEXT_URL"))
// 	f.Slides(os.Getenv("SLIDES_PATH"))
// 	f.Start(os.Getenv("PATHLESS"))
// 	select {}
// }

func main() {
	f := frame.NewFrame()
	f.Home("/timefactory.svg", "the perpetual motion machine")
	f.Text("https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md")
	f.Slides("/slides")
	f.Start("")
	select {}
}
