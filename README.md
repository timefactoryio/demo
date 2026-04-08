# demo

A demonstration application for the [frame](https://github.com/timefactoryio/frame) system that showcases dynamic content rendering in the [pathless](https://github.com/timefactoryio/pathless) viewport allocator.


```go
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
```

## Overview

This app serves three types of frames:
- **Home**: A custom landing page with logo and heading
- **Text**: Markdown content rendered from a URL or file
- **Slides**: An image slideshow from a directory

### Environment Variables

Configure the app using environment variables:

| Variable   | Description                           | Default                 |
| ---------- | ------------------------------------- | ----------------------- |
| `LOGO`     | Path to SVG logo file                 | -                       |
| `TITLE`    | Text displayed on home frame          | -                       |
| `TEXT`     | URL or file path for markdown content | -                       |
| `SLIDES`   | Directory containing slide images     | -                       |
| `PATHLESS` | URL of pathless instance              | `http://localhost:1000` |

### Example Configuration

```yaml
environment:
	- PATHLESS=https://timefactory.io
	- LOGO=https://zero.s3.timefactory.io/timefactory.svg
	- TITLE=the perpetual motion machine
	- TEXT=https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md
	- SLIDES=/slides
```
## Quick Start

### Using Docker Compose

The easiest way to run the demo:

```bash
docker compose up
```

This starts both services:
- **pathless** on `http://localhost:1000` (the viewport interface)
- **demo** on `http://localhost:1001` (the frame server)
