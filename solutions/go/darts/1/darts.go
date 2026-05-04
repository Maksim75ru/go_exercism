package darts

import "math"

func Score(x, y float64) int {
	sum := math.Pow(x, 2) + math.Pow(y, 2)
    gipotenuza := math.Sqrt(sum) 

    if gipotenuza <= 1.0 {
        return 10
    } else if gipotenuza > 1.0 && gipotenuza <= 5.0 {
        return 5
    } else if gipotenuza > 5.0 && gipotenuza <= 10.0 {
        return 1
    } else {
        return 0
    }
}
