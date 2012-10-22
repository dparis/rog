package main

import (
	"github.com/ajhager/rog"
    "math/rand"
)

const (
	min_icicle_fall_delay = 80.0
	max_icicle_fall_delay = 400.0
)

type Icicle struct {
	Entity
	id         float64
	fall_delay float64
    fall_timer Timer
}

func NewIcicle(app App, x, y int) *Icicle {
    fall_delay := RandRangeFloat(min_icicle_fall_delay, max_icicle_fall_delay)
	new_icicle := Icicle{
		Entity: Entity{
			x: x, y: y,

			min_x: 0, min_y: 0,
			max_x: app.Width(),
            max_y: app.Height(),

            fg:    rog.White.Alpha(rog.Blue, RangeScaleFloat(min_icicle_fall_delay, max_icicle_fall_delay, fall_delay)),
            bg:    rog.Black,
            glyph: 'V',
        },
        id:         rand.Float64(),
        fall_delay: fall_delay,
        fall_timer: NewTimer(fall_delay),
    }

	return &new_icicle
}


func (self *Icicle) Update(app App) Messages {
    if CheckTimer(self.fall_timer) {

        if self.y >= self.max_y {
            return Messages { MSG_REMOVE{} }
        } else {
            x, y := MoveDown(self)
            return Messages { MSG_CLEAR{x, y} }
        }

    }
    return Messages { MSG_NIL{} }
}