package main

import (
    "github.com/ajhager/rog"
)

type Renderable interface {
    X() int
    Y() int
    SetX(int)
    SetY(int)

    Fg() rog.RGB
    Bg() rog.RGB
    Glyph() rune
}

func Render(self Renderable) { 
    rog.Set(
        self.X(), self.Y(),
        self.Fg(), self.Bg(),
        string(self.Glyph()),
    )
}

