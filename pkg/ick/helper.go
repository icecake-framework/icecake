package ick

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func mini(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
