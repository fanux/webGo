package gochess

//PointState is
const (
	PointStateBlank int = iota
	PointStateWhite
	PointStateBlack
)

//DefaultChessboardSize is
const DefaultChessboardSize int = 19 * 19

// Level
const (
	Level18K int = -18
	Level1D  int = 0
	Level2D  int = 1
	Level3D  int = 2
)

//DropPoint is
type DropPoint struct {
	State   int //may be white or black, if player pass, it blank
	X       int
	Y       int
	Comment string
}

//Chessboard is
type Chessboard struct {
	Points [][]int
	Size   int // default is 19*19
}

//Player is
type Player struct {
	Name    string
	Level   int
	Comment string
}

//ChessManual is
type ChessManual struct {
	*Chessboard
	Players    [2]*Player // Player 0 using Black, Player 1 using White
	StartFlag  int        // "rang zi", if has 9 "rang zi" StartFlag is set to 9
	Komi       float32    // 5 Komi and a half
	DropPoints []DropPoint
}
