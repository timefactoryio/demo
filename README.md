# demo

A demonstration application for the [frame](https://github.com/timefactoryio/frame) system that showcases dynamic content rendering in the [pathless](https://github.com/timefactoryio/pathless) viewport allocator.

## Overview

This app serves three types of frames:
- **Home**: A custom landing page with logo and heading
- **Text**: Markdown content rendered from a URL or file
- **Slides**: An image slideshow from a directory

## Quick Start

### Using Docker Compose

The easiest way to run the demo:

```bash
docker compose up
```

This starts both services:
- **pathless** on `http://localhost:1000` (the viewport interface)
- **demo** on `http://localhost:1001` (the frame server)

### Environment Variables

Configure the app using environment variables:

| Variable      | Description                           | Default                 |
| ------------- | ------------------------------------- | ----------------------- |
| `LOGO_PATH`   | Path to SVG logo file                 | -                       |
| `HEADING`     | Text displayed on home frame          | -                       |
| `TEXT_URL`    | URL or file path for markdown content | -                       |
| `SLIDES_PATH` | Directory containing slide images     | -                       |
| `PATHLESS`    | URL of pathless instance              | `http://localhost:1000` |

### Example Configuration

From [`docker-compose.yml`](docker-compose.yml):

```yaml
environment:
  - LOGO_PATH=/logo.svg
  - HEADING=the perpetual motion machine
  - TEXT_URL=https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md
  - SLIDES_PATH=/slides
```

## Development

### Prerequisites

- Go 1.25.6+
- Docker (optional)

### Building from Source

```bash
go build -o demo
./demo
```

### Project Structure

```
demo/
├── main.go              # Application entry point
├── slides/              # Image directory for slideshow
│   └── sort.json        # Optional: defines slide order
├── timefactory.svg      # Logo file
├── Dockerfile           # Multi-stage build
└── docker-compose.yml   # Local development setup
```

### Slides Configuration

Place images in the slides directory. Optionally create [`sort.json`](slides/sort.json) to define display order:

```json
["first", "second", "third"]
```

File names should match (without extension).

## Usage with Pathless

Once running, open `http://localhost:1000` in your browser.

### Keyboard Controls

| Key   | Action                                        |
| ----- | --------------------------------------------- |
| `q`   | Previous frame                                |
| `e`   | Next frame                                    |
| `1`   | Toggle fullscreen                             |
| `2`   | Two panel layout (toggle horizontal/vertical) |
| `3`   | Three panel layout (cycle positions)          |
| `Tab` | Cycle focus between panels                    |
| `z`   | Show keyboard overlay                         |

## Deployment

See [`docker-compose.yml`](pathless/docker-compose.yml) in the factory workspace for production deployment with Traefik reverse proxy.

## Related Projects

- [pathless](https://github.com/timefactoryio/pathless) - Viewport allocator
- [frame](https://github.com/timefactoryio/frame) - Frame generation system

## License

See individual project licenses.


```go
package main

import (
	"github.com/timefactoryio/frame"
)

func main() {
	f := frame.NewFrame()
	f.Home("/timefactory.svg", "the perpetual motion machine")
	f.Text("https://raw.githubusercontent.com/timefactoryio/pathless/refs/heads/main/README.md")
	f.Slides("/slides")
	f.Start("")
	select {}
}
```