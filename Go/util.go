package chess

import (
	"fmt"

	"github.com/fanux/webGo/common"
)

//Lazi is
func Lazi(chessManual *common.ChessManual, dropPoint common.DropPoint) (invalid bool, tizi []common.DropPoint) {
	invalid = false
	return
}

//IsDead is
func IsDead(chessboard *common.Chessboard, dropPoint common.DropPoint) (bool, []common.DropPoint) {
	already := make(map[common.DropPoint]int)
	if recursion(chessboard, already, dropPoint, dropPoint.State) {
		dead := new([]common.DropPoint)
		for k := range already {
			*dead = append(*dead, k)
		}
		return true, *dead
	}
	return false, nil
}

// return false means not dead
func recursion(chessboard *common.Chessboard, already map[common.DropPoint]int, dropPoint common.DropPoint, color int) bool {
	if dropPoint.X < 0 || dropPoint.X >= chessboard.Size || dropPoint.Y < 0 || dropPoint.Y >= chessboard.Size {
		fmt.Println("out of range")
		return true
	}
	fmt.Printf("X is:%d, Y is:%d, Point:%d\n", dropPoint.X, dropPoint.Y, chessboard.Points[dropPoint.X][dropPoint.Y])
	if chessboard.Points[dropPoint.X][dropPoint.Y] == common.PointStateBlank {
		// Not Dead
		fmt.Println("is blank, alive")
		return false
	}
	if _, ok := already[dropPoint]; ok {
		fmt.Println("already in map")
		return true
	}
	if dropPoint.State != color {
		fmt.Println("other color")
		return true
	}

	already[dropPoint] = color
	dropPoint.X = dropPoint.X - 1
	fmt.Println("chek (x - 1)")
	if !recursion(chessboard, already, dropPoint, color) {
		return false
	}
	fmt.Println("chek (x + 1)")
	dropPoint.X = dropPoint.X + 2
	if !recursion(chessboard, already, dropPoint, color) {
		return false
	}
	dropPoint.X = dropPoint.X - 1
	dropPoint.Y = dropPoint.Y - 1
	fmt.Println("chek (y - 1)")
	if !recursion(chessboard, already, dropPoint, color) {
		return false
	}
	dropPoint.Y = dropPoint.Y + 2
	fmt.Println("chek (y + 1)")
	if !recursion(chessboard, already, dropPoint, color) {
		return false
	}
	return true
}
