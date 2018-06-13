package poppler

import "math"

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (r Rectangle) Left() float64 {
	return math.Min(r.X1, r.X2)
}

func (r Rectangle) Right() float64 {
	return math.Max(r.X1, r.X2)
}

func (r Rectangle) Top() float64 {
	return math.Min(r.Y1, r.Y2)
}

func (r Rectangle) Bottom() float64 {
	return math.Min(r.Y1, r.Y2)
}

func (r Rectangle) Width() float64 {
	return math.Abs(r.X2 - r.X1)
}

func (r Rectangle) Height() float64 {
	return math.Abs(r.Y2 - r.Y1)
}
