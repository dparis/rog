rog A roguelike game library written in go
===
![Rog Screenshot](http://hagerbot.com/img/screenshot_rog_fov.png)

* 24bit, scaling console with custom font support
* Cross platform rendering and input
* Field of view, lighting, and pathfinding algorithms
* Procedural color and palette manipulation

[Documentation](http://hagerbot.com/rog/docs.html "Documentation")

```go
package main

import (
    "hagerbot.com/rog"
)

func main() {
    rog.Open(20, 11, 2, false, "rog", nil)
    for rog.Running() {
        rog.Set(5, 5, nil, nil, "Hello, 世界!")
        if rog.Key() == rog.Esc {
            rog.Close()
        }
        rog.Flush()
    }
}
```

Notes
-----
* You will need glfw dynamic libs and development headers installed for now.
* On Windows you can build your project with `go build -ldflags -Hwindowsgui` to inhibit the console window that pops up by default.

Plans
-----
* Website, tutorial, and more demos
* Audio generation and output
* Image (subcell) blitting
* Custom drawing callback
* Noise generators
* Merge in lighting
* World creation
* More fov algorithms
