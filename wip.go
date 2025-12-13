package main

import (
	"fmt"
	"path/filepath"

	"github.com/timefactoryio/frame"
	"github.com/timefactoryio/frame/zero"
)

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

// func Keyboard(f *frame.Frame, asFrame bool) *zero.One {
// 	css := f.CSS(f.ToString("./keyboard.css"))
// 	js := f.JS(`
// (function(){
//   const keys = [
//     ['tab', '', ''],
//     ['1', '2', '3'],
//     ['q', 'w', 'e'],
//     ['a', 's', 'd']
//   ];

//   const space = pathless.space();
//   const grid = space.querySelector('.grid');
//   if (!grid) return;

// function render() {
//     grid.innerHTML = '';
//     const keybinds = pathless.keyboard();

//     keys.forEach((row, rowIndex) => {
//       row.forEach((k, colIndex) => {
//         if (!k) return; // Skip empty cells

//         const entry = keybinds[k];
//         const pressed = entry?.pressed;
//         const label = Array.isArray(entry?.label) ? entry.label[1] : '';

//         const keyEl = document.createElement('div');
//         keyEl.className = 'key';
//         keyEl.dataset.key = k;
//         keyEl.style.gridRow = rowIndex + 1;
//         keyEl.style.gridColumn = colIndex + 1;

//         const top = document.createElement('div');
//         top.className = 'top';
//         top.textContent = k.toUpperCase();
//         keyEl.appendChild(top);

//         const bottom = document.createElement('div');
//         bottom.className = 'bottom';
//         bottom.textContent = label || '';
//         keyEl.appendChild(bottom);

//         if (pressed) keyEl.classList.add('pressed');
//         if (k === 'tab') keyEl.classList.add('tabkey');
//         grid.appendChild(keyEl);
//       });
//     });
// }

//   render();
//   window.addEventListener('keyboardchange', render);
// })();
// `)
// 	html := zero.One(template.HTML(`<div class="grid"></div>`))
// 	final := f.Build("keyboard", asFrame, &html, &css, &js)
// 	return final
// }
