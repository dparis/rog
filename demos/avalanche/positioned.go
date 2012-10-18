package main

type IPositioned interface {
    X() int
    Y() int
    SetX() int
    SetY() int
}

