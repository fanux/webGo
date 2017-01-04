package chess

import (
	"fmt"
	"testing"

	"github.com/fanux/webGo/common"
)

func TestIsDead(t *testing.T) {
	var chessboard common.Chessboard
	var dropPoint common.DropPoint
	var isDead bool
	var dropPoints []common.DropPoint

	chessboard = common.Chessboard{Size: 5, Points: [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	dropPoint = common.DropPoint{common.PointStateBlack, 0, 0, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

	chessboard = common.Chessboard{Size: 5, Points: [][]int{
		{0, 2, 0, 0, 0},
		{2, 1, 2, 0, 0},
		{0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}}
	dropPoint = common.DropPoint{common.PointStateWhite, 1, 1, ""}

	isDead, dropPoints = IsDead(&chessboard, dropPoint)
	fmt.Println(isDead, dropPoints)

}
