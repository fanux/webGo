package gochess

import (
	"fmt"
	"testing"
)

func TestLazi(t *testing.T) {
	var chessManual *ChessManual
	var chessboard Chessboard
	//var dropPoint DropPoint
	//var dropPoints []DropPoint

	//Test Ko
	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 1, 2, 0, 0},
		{1, 0, 0, 2, 0},
		{0, 1, 2, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	chessManual = &ChessManual{&chessboard, [2]*Player{nil, nil}, 0, 0, []DropPoint{{PointStateWhite, 0, 0, ""}}}
	Lazi(chessManual, DropPoint{PointStateBlack, 1, 1, ""})
	Lazi(chessManual, DropPoint{PointStateWhite, 1, 2, ""})
	Lazi(chessManual, DropPoint{PointStateBlack, 1, 1, ""})

	fmt.Println("=======test eat two return back one=====")
	//Test daerhuanyi
	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 1, 1, 2, 0},
		{1, 2, 0, 0, 2},
		{0, 1, 1, 2, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	chessManual = &ChessManual{&chessboard, [2]*Player{nil, nil}, 0, 0, []DropPoint{{PointStateWhite, 0, 0, ""}}}
	Lazi(chessManual, DropPoint{PointStateBlack, 1, 2, ""})
	Lazi(chessManual, DropPoint{PointStateWhite, 1, 3, ""})
	Lazi(chessManual, DropPoint{PointStateBlack, 1, 2, ""})

	fmt.Println("=======eat multipart=====")
	//Test daerhuanyi
	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 1, 1, 2, 1},
		{1, 2, 0, 0, 2},
		{0, 1, 1, 2, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0}}}
	chessManual = &ChessManual{&chessboard, [2]*Player{nil, nil}, 0, 0, []DropPoint{{PointStateWhite, 0, 0, ""}}}
	Lazi(chessManual, DropPoint{PointStateWhite, 1, 3, ""})
}

func TestIsDead(t *testing.T) {
	var chessboard Chessboard
	var dropPoint DropPoint
	var isDead bool
	var dropPoints []DropPoint

	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	dropPoint = DropPoint{PointStateBlack, 0, 0, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 2, 0, 0, 0},
		{2, 1, 2, 0, 0},
		{0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	dropPoint = DropPoint{PointStateWhite, 1, 1, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 2, 0, 0, 0},
		{2, 1, 2, 0, 0},
		{1, 1, 1, 2, 0},
		{2, 2, 1, 2, 0},
		{0, 2, 1, 1, 2}}}
	dropPoint = DropPoint{PointStateWhite, 1, 1, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

	chessboard = Chessboard{Size: 5, Points: [][]int{
		{0, 2, 0, 0, 0},
		{2, 1, 2, 0, 0},
		{1, 1, 1, 2, 0},
		{2, 2, 1, 2, 0},
		{0, 2, 1, 1, 0}}}
	dropPoint = DropPoint{PointStateWhite, 1, 1, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

}
