package gochess

//NewDefaultChessManual is
func NewDefaultChessManual() *ChessManual {
	return &ChessManual{
		&Chessboard{
			Size:   DefaultChessboardSize,
			Points: [DefaultChessboardSize][DefaultChessboardSize]int{},
		},
		[2]*Player{nil, nil},
		0,
		6.5,
		[]DropPoint{},
		DropPoint{0, 0, 0, ""},
		false,
	}
}

//GetPointState is
func (c *Chessboard) GetPointState(x, y int) (bool, int) {
	if x < 0 || y < 0 || x >= c.Size || y >= c.Size {
		return false, -1
	}
	return true, c.Points[x][y]
}

//SetState is
func (c *Chessboard) SetState(dropPoint DropPoint) bool {
	if dropPoint.X < 0 || dropPoint.Y < 0 || dropPoint.X >= c.Size || dropPoint.Y >= c.Size {
		return false
	}
	c.Points[dropPoint.X][dropPoint.Y] = dropPoint.State
	return true
}

//CountQi is
func (c *Chessboard) CountQi(dropPoint DropPoint) int {
	qi := 0
	var b bool
	var s int
	b, s = c.GetPointState(dropPoint.X-1, dropPoint.Y)
	if b && (s == PointStateBlank || s == dropPoint.State) {
		qi++
	}
	b, s = c.GetPointState(dropPoint.X+1, dropPoint.Y)
	if b && (s == PointStateBlank || s == dropPoint.State) {
		qi++
	}
	b, s = c.GetPointState(dropPoint.X, dropPoint.Y-1)
	if b && (s == PointStateBlank || s == dropPoint.State) {
		qi++
	}
	b, s = c.GetPointState(dropPoint.X, dropPoint.Y+1)
	if b && (s == PointStateBlank || s == dropPoint.State) {
		qi++
	}
	return qi
}

//GetDiffPoint is
func (c *Chessboard) GetDiffPoint(dropPoint DropPoint) []DropPoint {
	var b bool
	var s int
	dropPoints := *new([]DropPoint)
	temp := DropPoint{}

	b, s = c.GetPointState(dropPoint.X-1, dropPoint.Y)
	if b && s != dropPoint.State && s != PointStateBlank {
		temp.X = dropPoint.X - 1
		temp.Y = dropPoint.Y
		temp.State = s

		dropPoints = append(dropPoints, temp)
	}

	b, s = c.GetPointState(dropPoint.X+1, dropPoint.Y)
	if b && s != dropPoint.State && s != PointStateBlank {
		temp.X = dropPoint.X + 1
		temp.Y = dropPoint.Y
		temp.State = s

		dropPoints = append(dropPoints, temp)
	}

	b, s = c.GetPointState(dropPoint.X, dropPoint.Y-1)
	if b && s != dropPoint.State && s != PointStateBlank {
		temp.X = dropPoint.X
		temp.Y = dropPoint.Y - 1
		temp.State = s

		dropPoints = append(dropPoints, temp)
	}

	b, s = c.GetPointState(dropPoint.X, dropPoint.Y+1)
	if b && s != dropPoint.State && s != PointStateBlank {
		temp.X = dropPoint.X
		temp.Y = dropPoint.Y + 1
		temp.State = s

		dropPoints = append(dropPoints, temp)
	}

	return dropPoints
}

//CheckFowardPoint is
func (c *ChessManual) CheckFowardPoint(dropPoint DropPoint) bool {
	if len(c.DropPoints) < 2 {
		return false
	}
	point := c.DropPoints[len(c.DropPoints)-2]
	if dropPoint.X == point.X && dropPoint.Y == point.Y && dropPoint.State == point.State {
		return true
	}
	return false
}

//CheckFowardPoint1 two continuation drop must different
func (c *ChessManual) CheckFowardPoint1(dropPoint DropPoint) bool {
	if len(c.DropPoints) < 1 {
		return true
	}
	point := c.DropPoints[len(c.DropPoints)-1]
	if dropPoint.X == point.X && dropPoint.Y == point.Y && dropPoint.State == point.State {
		return false
	}
	return true
}

//CheckKoPoint is
func (c *ChessManual) CheckKoPoint(dropPoint DropPoint) bool {
	if dropPoint.X == c.Ko.X && dropPoint.Y == c.Ko.Y && dropPoint.State == c.Ko.State {
		return true
	}
	return false
}
