package main

import (
    "github.com/ajhager/rog"
)

type Movable interface {
    X() int
    Y() int
    SetX(int)
    SetY(int)
    // provided
}

func Move(self Movable, dx, dy int) {
    x := self.X()
    y := self.Y()

    rog.Set(x, y, nil, nil, " ")

    self.SetX(x + dx)
    self.SetY(y + dy)
}

func MoveLeft(self Movable) { Move(self, -1, 0) }
func MoveRight(self Movable) { Move(self, 1, 0) }
func MoveUp(self Movable) { Move(self, 0, -1) }
func MoveDown(self Movable) { Move(self, 0, 1) }
