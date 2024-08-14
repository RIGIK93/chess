package board

import "log"

type Board struct {
	white_move       bool        // is white to move?
	white_king_moved bool        // for determining castling rules
	black_king_moved bool        // for determining castling rules
	no_capture_count uint8       // for 50-move-rule. The amount of moves passed without captures
	pieces           [8][8]Piece // board[row][column] or board[y][x]
}

func (board *Board) IsWhiteToMove() bool {
	return board.white_move
}

// Outputs piece color that is to move
func (board *Board) ColorToMove() PieceColor {
	if board.IsWhiteToMove() {
		return WHITE
	} else {
		return BLACK
	}
}

func (board *Board) WhiteKingMoved() bool {
	return board.white_king_moved
}

func (board *Board) BlackKingMoved() bool {
	return board.black_king_moved
}

// Get legal moves
func (board *Board) GetMoves() (moves []Board) {
	moves := []Board{}

	// check for checks
	log.Fatalf("TODO")

	// check for mates

	for y, row := range board.pieces {
		for x, piece := range row {
			_, color := GetPieceColor(piece)

			if color == EMPTY || color != board.ColorToMove() {
				continue
			}

			moves = append(moves, piece.GetMoves(*board, NewSquareByInt(uint8(x), uint8(y)))...)
		}
	}

	return
}

func (board *Board) GetPiece(sq Square) (empty bool, piece Piece) {
	piece = board.pieces[sq.GetY()][sq.GetX()]
	return piece.Value() == 0, piece
}

func (board *Board) SetPiece(sq Square, p Piece) {
	board.pieces[sq.GetY()][sq.GetX()] = p
}

func (board *Board) Move() error {
	return nil
}

// func (board *Board) Clone() (clone Board) {
// 	return Board {
// 		white_move: board.white_move,
// 		white_king_moved: board.white_king_moved,
// 		black_king_moved: board.black_king_moved,
// 		no_capture_count: board.no_capture_count,
// 		pieces: ,
// 	}
// }

func NewBoard() Board {
	// return Board{
	// 	is_white_turn: true,
	// 	pieces: ,
	// }
}
