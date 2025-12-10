package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/timefactoryio/frame"
	"github.com/timefactoryio/frame/zero"
)

func main() {
	pathless := os.Getenv("PATHLESS_URL")
	api := os.Getenv("API_URL")

	f := frame.NewFrame(pathless, api)
	f.Home("the perpetual motion machine", "timefactoryio", "")
	f.README(f.ToBytes("./README.md"))
	f.BuildSlides("slides")
	f.AddPath("./img")
	// BuildVideo(f, "./video/beef.mp4")
	// BuildVideo(f, "./video/hungary.mp4")
	f.Keyboard(true)
	f.Serve()
	select {}
}

func BuildVideo(f *frame.Frame, filePath string) *zero.One {
	f.AddFile(filePath, "video")
	name := filepath.Base(filePath)
	name = name[:len(name)-len(filepath.Ext(name))]

	video := f.Video("")
	css := f.CSS(f.VideoCSS())
	js := f.JS(fmt.Sprintf(`
(async function() {
    const el = pathless.frame().querySelector('video');
    if (!el) return;

    const { data: src } = await pathless.fetch(apiUrl + '/video/%s', { key: 'video-%s' });
    el.src = src;

    el.addEventListener('loadedmetadata', () => {
        el.currentTime = pathless.state().t || 0;
    }, { once: true });

    pathless.keybind(
        k => { if (k === ' ') el.paused ? el.play() : el.pause(); },
        () => { pathless.update('t', el.currentTime || 0); el.pause(); }
    );
})();
    `, name, name))
	return f.Build("video", true, video, &css, &js)
}
