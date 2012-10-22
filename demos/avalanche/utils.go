package main

import (
    "math/rand"
)

func RandRangeFloat(min, max float64) float64 {
    return (rand.Float64() * (max - min)) + min
}

func RandRangeInt(min, max int) int {
    return int(RandRangeFloat(float64(min), float64(max)))
}

func RangeScaleFloat(min, max, val float64) float64 {
    return 1-(val/(min+max))
}
