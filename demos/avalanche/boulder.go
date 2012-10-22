package main

import (
    "github.com/ajhager/rog"
)

const (
    min_size = 2
    max_size = 4
    boulder_ = 30
    min_boulder_show_delay = 500
    max_boulder_show_delay = 1000
    min_boulder_fall_delay = 30
    max_boulder_fall_delay = 60
)

const (
    _ = iota
    STATE_SHOW
    STATE_FALL
)

type Boulder struct {
    Entity
    show_delay float64
    show_timer Timer
    fall_delay float64
    fall_timer Timer
    size int
    state int
}

func NewBoulder(app App) *Boulder {
    size := RandRangeInt(min_size, max_size)
    show_delay := RandRangeFloat(min_boulder_show_delay, max_boulder_show_delay)
    fall_delay := RandRangeFloat(min_boulder_fall_delay, max_boulder_fall_delay)
    b := Boulder {
        Entity: Entity {
            x: RandRangeInt(0, app.Width()),
            y: -(size - 2),

            min_x: 0, min_y: 0,
            max_x: app.Width(),
            max_y: app.Height(),

            fg:    rog.White.Alpha(rog.Blue, RangeScaleFloat(min_boulder_fall_delay, max_boulder_fall_delay, fall_delay)),
            bg:    rog.Black,
            glyph: '#',
        },
        show_delay: show_delay,
        show_timer: NewTimer(show_delay),
        fall_delay: fall_delay,
        fall_timer: NewTimer(fall_delay),
        size: size,
        state: STATE_SHOW,
    }

    return &b
}

func (self *Boulder) GetClears(x, y int) Messages {
    clears := make(Messages, 0)
    for i:= 0; i < self.size; i++ {
        for ii:= 0; ii < self.size; ii++ {
            clear := MSG_CLEAR { x + i, y + ii }
            clears = append(clears, clear)
        }
    }
    return clears
}

func (self *Boulder) Render() {
    x := self.X()
    y := self.Y()
    for i := 0; i < self.size; i++ {
        for ii := 0; ii < self.size; ii++ {
            rog.Set(x + i, y + ii,
                self.Fg(), self.Bg(),
                string(self.Glyph()),
            )
        }
    }
}

func (self *Boulder) Update(app App) Messages {
    switch self.state {
        case STATE_SHOW:
            return self.DoShow(app)
        case STATE_FALL:
            return self.DoFall(app)
    }; panic(nil)
}

func (self *Boulder) DoShow(app App) Messages {
    switch CheckTimer(self.show_timer) {
        case false:
            return Messages { MSG_NIL{} }
        case true:
            self.state = STATE_FALL
            return Messages { MSG_NIL{} }
    }; panic(nil)
}

func (self *Boulder) DoFall(app App) Messages {
    if CheckTimer(self.fall_timer) {
        if self.y >= self.max_y {
            return Messages { MSG_REMOVE{} }
        } else {
            x, y := MoveDown(self)
            return self.GetClears(x, y)
        }
    }
    return Messages { MSG_NIL{} }
}