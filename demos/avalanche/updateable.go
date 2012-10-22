package main

type Updatable interface {
    Update(app App) Messages
}
