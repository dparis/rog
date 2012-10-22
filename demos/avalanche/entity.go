package main

//import "fmt"

import (
//    "runtime"
    "github.com/ajhager/rog"
)

type IEntity interface {
    Updatable
    Renderable
}

type Entity struct {
    IEntity,
    Movable,

    x, y int

    min_x, min_y int
    max_x, max_y int

    fg, bg rog.RGB
    glyph rune
}

func (self *Entity) X() int { return self.x }
func (self *Entity) Y() int { return self.y }

func (self *Entity) SetX(v int) { self.x = v }
func (self *Entity) SetY(v int) { self.y = v }

func (self *Entity) MinX() int { return self.min_x }
func (self *Entity) MinY() int { return self.min_y }
func (self *Entity) MaxX() int { return self.max_x }
func (self *Entity) MaxY() int { return self.max_y }

func (self *Entity) Fg() rog.RGB { return self.fg }
func (self *Entity) Bg() rog.RGB { return self.bg }
func (self *Entity) Glyph() rune { return self.glyph }

func (self *Entity) Render() {
    rog.Set(
        self.X(), self.Y(),
        self.Fg(), self.Bg(),
        string(self.Glyph()),
    )
}

