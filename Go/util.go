package gochess

import "fmt"

//Lazi is
func Lazi(chessManual *ChessManual, dropPoint DropPoint) []DropPoint {
	fmt.Printf("\n\nLazi X:%d, Y:%d, State:%d\n", dropPoint.X, dropPoint.Y, dropPoint.State)
	f, dropPoints := isLaziInvalid(chessManual, dropPoint)
	if f {
		for _, d := range dropPoints {
			chessManual.Points[d.X][d.Y] = PointStateBlank
		}
		chessManual.DropPoints = append(chessManual.DropPoints, dropPoint)
		fmt.Println("eat up: ", dropPoints)
		return dropPoints
	}
	return nil
}

func isLaziInvalid(chessManual *ChessManual, dropPoint DropPoint) (bool, []DropPoint) {
	chessboard := chessManual.Chessboard
	f, s := chessboard.GetPointState(dropPoint.X, dropPoint.Y)
	if !f {
		fmt.Println("out of range")
		return false, nil
	}
	if s != PointStateBlank {
		fmt.Println("already has chess piece there")
		return false, nil
	}
	chessboard.SetState(dropPoint)

	diffPoints := chessboard.GetDiffPoint(dropPoint)
	fmt.Println("diff points is: ", diffPoints)

	kill := *new([]DropPoint)
	for _, d := range diffPoints {
		a, temp := IsDead(chessboard, d)
		if a {
			kill = append(kill, temp...)
		}
	}
	if len(kill) > 1 {
		fmt.Println("kill others killing > 1")
		return true, kill
	}
	// judge Ko (dajie)
	if len(kill) == 1 {
		fmt.Println("killing == 1")
		if chessboard.CountQi(dropPoint) == 0 && chessManual.CheckFowardPoint(dropPoint) {
			fmt.Println("Ko you must stop one time in here")
			chessManual.Points[dropPoint.X][dropPoint.Y] = PointStateBlank
			return false, nil
		}
		return true, kill
	}
	if b, _ := IsDead(chessboard, dropPoint); b && len(kill) == 0 {
		fmt.Println("invalid drop, not allowed sucide!")
		chessManual.Points[dropPoint.X][dropPoint.Y] = PointStateBlank
		return false, nil
	}

	return true, nil
}

//IsDead is
func IsDead(chessboard *Chessboard, dropPoint DropPoint) (bool, []DropPoint) {
	already := make(map[DropPoint]int)
	if recursion(chessboard, already, dropPoint, dropPoint.State) {
		dead := new([]DropPoint)
		for k := range already {
			*dead = append(*dead, k)
		}
		return true, *dead
	}
	return false, nil
}

// return false means not dead
func recursion(chessboard *Chessboard, already map[DropPoint]int, dropPoint DropPoint, color int) bool {
	if dropPoint.X < 0 || dropPoint.X >= chessboard.Size || dropPoint.Y < 0 || dropPoint.Y >= chessboard.Size {
		fmt.Println("out of range")
		return true
	}
	fmt.Printf("X is:%d, Y is:%d, Point:%d\n", dropPoint.X, dropPoint.Y, chessboard.Points[dropPoint.X][dropPoint.Y])
	if chessboard.Points[dropPoint.X][dropPoint.Y] == PointStateBlank {
		// Not Dead
		fmt.Println("is blank, alive")
		return false
	}
	if _, ok := already[dropPoint]; ok {
		fmt.Println("already in map")
		return true
	}
	if chessboard.Points[dropPoint.X][dropPoint.Y] != color {
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
