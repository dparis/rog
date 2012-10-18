package main

import (
    "github.com/ajhager/rog"
)

type Player struct {
    Entity
    points int
    life int
}

func NewPlayer(app App, x, y int) *Player {
    player := Player {
        Entity: Entity {
            x:x, y:y,
            min_x: 0, min_y: 0,
            max_x: app.Width(),
            max_y: app.Height(),

            fg: rog.Green,
            bg: rog.Black,
            glyph: '웃',
        },
        points: 35,
        life: 1,
    }


    return &player
}

