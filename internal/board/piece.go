package board

import "log"

type Piece interface {
	Value() float32                           // returns the value of a piece
	GetMoves(board Board, pos Square) []Board // get legal moves
	Name() string
}

type PieceColor string

const (
	WHITE PieceColor = "white"
	BLACK PieceColor = "black"
	EMPTY PieceColor = ""
)

// Outputs 1 if the color is white, -1 if the color is black and 0 if the color is EMPTY
func (p *PieceColor) toInt() int {

	value_sign := 0

	if *p == WHITE {
		value_sign = -1
	} else if *p == BLACK {
		value_sign = 1
	}

	return value_sign
}

type PieceType uint8

const (
	PAWN PieceType = iota
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
)

func GetPieceColor(piece Piece) (is_occupied bool, color PieceColor) {
	if piece.Value() > 0 {
		return true, WHITE
	} else if piece.Value() < 0 {
		return true, BLACK
	} else {
		return false, EMPTY
	}
}

// Empty chess piece
type Empty struct {
}

func NewEmpty() Empty {
	return Empty{}
}

func (Empty) Value() float32 {
	return 0
}

func (Empty) GetMoves(Board, Square) []Board {
	return nil
}

func (Empty) Name() string {
	return ""
}

type Pawn struct {
	value float32
}

func (p Pawn) Value() float32 {
	return p.value
}

func (Pawn) Name() string {
	return "Pawn"
}

func (pawn Pawn) GetMoves(board Board, sq Square) (moves []Board) {
	moves = []Board{}

	board.SetPiece(sq, NewEmpty()) // since the pawn will be moved, I remove it from the initial place

	_, color := GetPieceColor(pawn)

	direction := color.toInt()

	// Check above
	_, one_above := NewSquareByInt(0, 1)
	var err error
	var square_above *Square

	if direction == 1 {
		err, square_above = sq.Added(one_above)
	} else {
		err, square_above = sq.Subtracted(one_above)
	}

	if err != nil {
		log.Fatal(err)
	}

	promote := func(b Board, sq *Square) []Board {
		log.Fatal("TODO")
	}

	if empty, _ := board.GetPiece(*square_above); empty {
		b := board
		b.SetPiece(*square_above, &pawn)
		moves = append(moves, b)
		// Promotion
		if y := square_above.GetY(); y == 7 || y == 0 {
			moves = append(moves, promote(b, square_above)...)
		}
	}

	// Check takes
	_, right_direction := NewSquareByInt(1, 0)
	if err, right := square_above.Added(right_direction); err == nil {
		_, p := board.GetPiece(*right)
		if is_occupied, c := GetPieceColor(p); is_occupied && c != color {
			b := board
			b.SetPiece(*right, &pawn)
			moves = append(moves, b)
			// Promotion
			if y := right.GetY(); y == 7 || y == 0 {
				moves = append(moves, promote(b, right)...)
			}
		}
	}

	if err, left := square_above.Subtracted(right_direction); err == nil {
		_, p := board.GetPiece(*left)
		if is_occupied, c := GetPieceColor(p); is_occupied && c != color {
			b := board
			b.SetPiece(*left, &pawn)
			moves = append(moves, b)
			// Promotion
			if y := left.GetY(); y == 7 || y == 0 {
				moves = append(moves, promote(b, left)...)
			}
		}
	}

	return
}

func NewPawn(color PieceColor) Pawn {
	p := Pawn{
		value: float32(color.toInt()) * 1.0,
	}

	if p.value == 0 {
		log.Fatal("ERROR! Piece color must either be WHITE or BLACK, not EMPTY")
	}

	return p
}

type Knight struct {
}

func (Knight) Value() float32 {
	return 3
}

func (Knight) Name() string {
	return "Knight"
}

func (Knight) GetMoves(board Board, sq Square) (moves []Board) {

}

type Bishop struct {
}

type Rook struct {
}

type Queen struct {
}

type King struct {
}
