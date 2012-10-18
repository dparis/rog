package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
    "github.com/ajhager/rog"
)

const (
    screen_width = 40
    screen_height = 20

)

type AvalancheApp struct {
    AppController
    entities []IEntity
    player *Player
    delay_time float64
    do_delay bool
    spawn_time float64
    do_spawn bool
}

const (
    _ = iota
    UPDATE_REMOVE
    UPDATE_COLLIDE
)

func (self *AvalancheApp) Run() {
    rog.Open(self.Width(), self.Height(), 1, false, self.Title(), nil)

    self.player = NewPlayer(
        self,
        self.Width() / 2, 
        self.Height() - 1,
    )
  
    self.StartDelayTimer()
    self.StartSpawnTimer()

    for rog.Running() {
        self.Update()
        self.Render()
        rog.Flush()
    }
}

func (self *AvalancheApp) Render() {
    for _, e := range self.entities {
        var R = e
        Render(R)
    }

    P := self.player
    Render(P)

    // rog.Set(0, 2, rog.White, rog.Black, strconv.Itoa(self.player.points))
    // rog.Set(0, 4, rog.White, rog.Black, fmt.Sprintf("%f", self.SpawnTime()))

    for i := self.player.life; i > 0; i-- {
        rog.Set(2 * i, 1, rog.Green, rog.Green, " ")
    }
}

func (self *AvalancheApp) Update() {
    removed := make([]int, 0)
    for idx, e := range self.entities {
        var U Updatable = e
        remove := U.Update()
        if e.X() == self.player.X() && e.Y() == self.player.Y() {
            self.HandleCollision()
            remove = UPDATE_REMOVE
        }
        
        if remove == UPDATE_REMOVE {
            removed = append(removed, idx)
        }

    }

    points := len(removed)

    l := self.entities
    for _, eidx := range removed {
        l = l[:eidx+copy(l[eidx:], l[eidx+1:])]
    }

    self.entities = l

    if self.do_spawn {
        pos := rand.Intn(self.Width() - 1)
        new_icicle := NewIcicle(self, pos, 0)
        self.entities = append(self.entities, new_icicle)
        self.do_spawn = false
    }

    if self.do_delay {
        self.spawn_time -= 10.0
        self.do_delay = false
    }

    if self.spawn_time < -100.0 {
        self.spawn_time = 600 - (float64(self.player.points) * .1)
        self.delay_time *= 1.5
        self.player.life += 1
    }

    self.HandleKeys()
    self.HandlePoints(points)
}

func (self *AvalancheApp) HandleKeys() {
    switch rog.Key() {
        case rog.Esc:
            rog.Close()
        case rog.Left:
            MoveLeft(self.player)
        case rog.Right:
            MoveRight(self.player)
    }        
}

func (self *AvalancheApp) HandlePoints(points int) {
    self.player.points += points
}

func (self *AvalancheApp) HandleCollision() {
    rog.Set(2 * self.player.life, 1, rog.Black, rog.Black, " ")
    self.player.life -= 1

    if self.player.life == 0 {
        fmt.Println("Your score:", self.player.points)
        rog.Close()
    }
}

func (self *AvalancheApp) GetDelayCallback() func() {
    cb_delay := func () {
        self.do_delay = true
        self.StartDelayTimer()
    }
    return cb_delay
}

func (self *AvalancheApp) StartDelayTimer() {
    time.AfterFunc(time.Duration(self.delay_time) * time.Millisecond, self.GetDelayCallback())
}

func (self *AvalancheApp) GetSpawnCallback() func() {
    cb_spawn := func () {
        self.do_spawn = true
        self.StartSpawnTimer()
    }
    return cb_spawn
}

func (self *AvalancheApp) SpawnTime() float64 {
    return math.Max(self.spawn_time, float64(self.player.points) * .006)
}

func (self *AvalancheApp) StartSpawnTimer() {
    time.AfterFunc(time.Duration(self.SpawnTime()) * time.Millisecond, self.GetSpawnCallback())
}


func main() {
    app := AvalancheApp {
        AppController {
            screen_width, screen_height,
            "rog demo: Avalanche",
        },
        make([]IEntity, 0),
        nil,
        10.0,
        false,
        800.0,
        false,
    }
    app.Run()
}
