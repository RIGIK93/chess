package board

import "errors"

type Square struct {
	x uint8
	y uint8
}

func NewSquareByInt(x uint8, y uint8) (error, *Square) {

	if x > 7 || y > 7 {
		return errors.New("Out of Bounds Error! Square position is out of bounds! Both x and y must not exceed 7 in the Square struct"), nil
	}

	return nil, &Square{
		x,
		y,
	}
}

func (raw *Square) GetX() uint8 {
	return raw.x
}

func (raw *Square) GetY() uint8 {
	return raw.y
}

func (original *Square) Added(to_add *Square) (error, *Square) {
	return NewSquareByInt(original.x+to_add.x, original.y+to_add.y)
}

func (original *Square) Subtracted(to_sub *Square) (error, *Square) {
	if to_sub.x > original.x || to_sub.y > original.y {
		return errors.New("Error! Out of bounds during subtraction error"), nil
	}

	return NewSquareByInt(original.x-to_sub.x, original.y-to_sub.y)
}

type Vec2 struct {
	x float32
	y float32
}

func NewVec2(x, y float32) *Vec2 {
	return &Vec2{
		x,
		y,
	}
}
