package main

import (
    "fmt"
    "math"
    "math/rand"
    "github.com/ajhager/rog"
)

const (
    screen_width = 40
    screen_height = 20
    thirty_frames = 1000.0 / 120.0

    avalanche_depth = -100
    max_spawn_msecs = 300
    min_spawn_msecs = 45
    base_delay_decay = 90
    delay_decay_factor = 1.9
)

var (
    max = 0
)

type Clear struct { x, y int }

type AvalancheApp struct {
    AppController
    entities []IEntity
    player *Player

    frame_time float64
    delay_time float64
    ice_spawn_time float64

    frame_timer Timer
    delay_ticker *Ticker
    ice_spawn_ticker *Ticker
    boulder_spawn_ticker *Ticker

    clears []Clear

    debug bool
}

func (self *AvalancheApp) SpawnIcicle() {
    pos := rand.Intn(self.Width() - 1)
    new_icicle := NewIcicle(self, pos, 0)
    self.entities = append(self.entities, new_icicle)
}

func (self *AvalancheApp) SpawnBoulder() {
    new_boulder := NewBoulder(self)
    self.entities = append(self.entities, new_boulder)
}

func (self *AvalancheApp) AddClear(x, y int) {
    self.clears = append(self.clears, Clear{x, y})
}


 func (self *AvalancheApp) Run() {
    rog.Open(self.Width(), self.Height(), 1, false, self.Title(), nil)

    self.player = NewPlayer(
        self,
        self.Width() / 2, 
        self.Height() - 1,
    )
  
    self.frame_timer = NewTimer(self.frame_time)
            
    self.delay_ticker = NewTicker(self.delay_time, func() {
        // increase icicle spawn rate
        self.ice_spawn_time -= 10.0
        // after avalanche is over
        self.delay_ticker.SetRate(self.delay_time)
    })

    self.ice_spawn_ticker = NewTicker(self.ice_spawn_time, func() {
        self.SpawnIcicle()
        self.ice_spawn_ticker.SetRate(math.Max(18, self.ice_spawn_time))
    })

    self.boulder_spawn_ticker = NewTicker(RandRangeFloat(1500, 2000), func() {
        if self.ice_spawn_time > 40 {
            self.SpawnBoulder()    
        }
        self.boulder_spawn_ticker.SetRate(RandRangeFloat(1500, 2000))
    })

    
    for rog.Running() {
        if CheckTimer(self.frame_timer) {
            self.Update()
            self.HandleKeys()
            self.Render()
            rog.Flush()
        }
    }
}

func (self *AvalancheApp) Render() {
    // clears
    for _, c := range self.clears {
        rog.Set(c.x, c.y, rog.Black, rog.Black, " ")
    }
    self.clears = make([]Clear, 0)
    // player
    self.player.Render()
    // entities (icicles)
    for _, e := range self.entities {
        e.Render()
    }
    // life bar
    for i := self.player.life; i > 0; i-- {
        rog.Set(2 * i, 1, rog.Green, rog.Green, " ")
    }
    // debug info
    if self.debug {
        rog.Set(0, 2, rog.White, rog.Black, fmt.Sprintf("Pts:%d", self.player.points))
        rog.Set(0, 4, rog.White, rog.Black, fmt.Sprintf("Pos:%d, %d", self.player.x, self.player.y))
        rog.Set(0, 6, rog.White, rog.Black, fmt.Sprintf("IceSpawn:%f.0", math.Max(self.ice_spawn_time, 5)))

    }
}

func (self *AvalancheApp) Update() {
    // reference current entities
    entities := self.entities
    // clear tracked entities
    self.entities = make([]IEntity, 0)
    // for each entity
    for _, e := range entities {
        keep := true
        msgs := e.Update(self)
        for _, msg := range msgs {
            switch t := msg.(type) {
                // award a point for removed entities
                case MSG_REMOVE:
                    self.player.points += 1
                    keep = false
                // track cleared tiles
                case MSG_CLEAR:
                    self.AddClear(t.x, t.y)
            }
        }

        if e.X() == self.player.X() && e.Y() == self.player.Y() {
            self.HandleCollision()
            keep = false
        }

        if keep {
            // keep tracking entity
            self.entities = append(self.entities, e) 
        }

    }

    if self.ice_spawn_time < avalanche_depth {
        // reset spawn time
        self.ice_spawn_time = max_spawn_msecs - (float64(self.player.points) * .1)
        self.ice_spawn_ticker.SetRate(self.ice_spawn_time)
        // increase dynamics rate
        self.delay_time *= delay_decay_factor
        self.delay_ticker.SetRate(self.delay_time)
        // award player life
        self.player.life += 1
    }

}

func (self *AvalancheApp) HandleKeys() {
    switch rog.Key() {
        case rog.Esc:
            rog.Close()
        case rog.Left, 'a':
            x, y := MoveLeft(self.player)
            self.AddClear(x, y)
        case rog.Right,'s':
            x, y := MoveRight(self.player)
            self.AddClear(x, y)
        case rog.Tab:
            self.debug = !self.debug
    }
}

func (self *AvalancheApp) HandleCollision() {
    // clear the last hp icon
    rog.Set(2 * self.player.life, 1, rog.Black, rog.Black, " ")
    // remove a life
    self.player.life -= 1
    // print score and quit if player dies
    if self.player.life == 0 {
        fmt.Println("Your score:", self.player.points)
        rog.Close()
    }
}

func main() {
    app := AvalancheApp {
        AppController {
            screen_width, screen_height,
            "rog demo: Avalanche",
        },
        make([]IEntity, 0),
        nil,

        thirty_frames,
        base_delay_decay,
        max_spawn_msecs,

        nil,
        nil,
        nil,
        nil,

        make([]Clear, 0),

        false,
    }
    app.Run()
}
