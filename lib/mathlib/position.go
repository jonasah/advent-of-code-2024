package mathlib

type Direction complex128

const (
	Right Direction = 1
	Down  Direction = 1i
	Left  Direction = -1
	Up    Direction = -1i
)

func (d Direction) TurnRight() Direction {
	return d * 1i
}

type Position complex128

func NewPosition(x, y int) Position {
	return Position(complex(float64(x), float64(y)))
}

func (p Position) X() int {
	return int(real(p))
}

func (p Position) Y() int {
	return int(imag(p))
}

func (p Position) Move(d Direction) Position {
	return p + Position(d)
}

func (p Position) IsWithin(p0, p1 Position) bool {
	return p0.X() <= p.X() && p.X() <= p1.X() && p0.Y() <= p.Y() && p.Y() <= p1.Y()
}
