package main

import (
    "time"
    "math/rand"
    "github.com/ajhager/rog"
)

const (
    min_fall_delay = 50.0
    max_fall_delay = 300.0
)

type Icicle struct {
    Entity
    id float64
    fall_delay time.Duration
    timer *time.Timer
}

func NewIcicle(app App, x, y int) *Icicle {
    fall_delay := (rand.Float64() * (max_fall_delay - min_fall_delay)) + min_fall_delay
    new_icicle := Icicle {
        Entity: Entity {
            x: x, y: y,

            min_x: 0, min_y: 0,
            max_x: app.Width(),
            max_y: app.Height(),

            fg: rog.White.Alpha(rog.Blue, 1 - (fall_delay / (min_fall_delay + max_fall_delay))),
            bg: rog.Black,
            glyph: 'V',
        },
        id: rand.Float64(),
        fall_delay: time.Duration(fall_delay) * time.Millisecond,
        timer: nil,
    }

    return &new_icicle
}

func (self *Icicle) Update() int {
    if self.y >= self.max_y {
        return UPDATE_REMOVE
    }

    if self.timer == nil {
        cb_fall := func() {
            MoveDown(self)
            self.timer = nil
        }
        self.timer = time.AfterFunc(self.fall_delay, cb_fall)

    }

    return 0
}
