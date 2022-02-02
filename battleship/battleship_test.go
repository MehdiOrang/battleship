package battleship

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Coordinate struct {
	Num    int
	Letter rune
}

type Position struct {
	Start Coordinate
	End   Coordinate
}

var (
	ErrIncorrectLetter     = errors.New("incorrect input in string")
	ErrOutOfGridBoundaries = errors.New("shot out of grid boundaries")
)

var GridSize = Position{Start: Coordinate{Num: 1, Letter: 'A'}, End: Coordinate{Num: 10, Letter: 'J'}}

type Grid struct {
	ships []*ship
	Shots int
}

type ship struct {
	start Coordinate
	end   Coordinate
	cap   int
}

type ShootResult struct {
	Hit  bool
	Sunk bool
}

func NewGrid(ships []ship) *Grid {
	var realShips []*ship
	for _, ship := range ships {
		newship := ship
		realShips = append(realShips, &newship)
	}

	grid := Grid{ships: realShips}
	return &grid
}

func shootInRange(shoot Coordinate, ship *ship) bool {
	startPointNum := ship.start.Num
	endPointNum := ship.end.Num
	startPointLetter := ship.start.Letter
	endPointLetter := ship.end.Letter
	for i := startPointNum; i <= endPointNum; i++ {
		for j := startPointLetter; j <= endPointLetter; j++ {
			if shoot.Num == i && shoot.Letter == j {
				ship.cap -= 1
				return true
			}
		}
	}
	return false
}
func (grid *Grid) Shoot(shotNum int, shotLetter string) (ShootResult, error) {
	//TODO: implement here
	shoot := Coordinate{Num: shotNum, Letter: rune(shotLetter[0])}
	if shotNum > 10 || shotNum < 1 {
		return ShootResult{}, ErrOutOfGridBoundaries
	}
	if shotLetter[0] < 65 || shotLetter[0] > 74 {
		return ShootResult{}, ErrIncorrectLetter
	}
	shootResult := ShootResult{}

	for _, v := range grid.ships {

		if shootInRange(shoot, v) {
			shootResult.Hit = true
		}
		if v.cap == 0 {
			shootResult.Sunk = true
		}
	}
	return shootResult, nil
}

func (grid *Grid) ResetShips() {
	grid.Shots = 0
}

type shot struct {
	num         int
	letter      string
	expectedHit bool
}

type testCase struct {
	shots        []shot
	expectedSunk bool
	err          error
}

func TestShoot(t *testing.T) {
	ships := getShips()

	testCases := []testCase{
		{shots: []shot{{1, "G", false}, {1, "H", true}, {1, "I", false}}, expectedSunk: false},
		{shots: []shot{{1, "H", true}, {2, "H", true}, {3, "H", true}, {4, "H", true}}, expectedSunk: true},
		{shots: []shot{{1, "D", false}, {7, "F", true}, {8, "F", true}}, expectedSunk: true},
		{shots: []shot{{10, "D", true}, {9, "D", false}, {10, "C", false},
			{10, "E", true}, {10, "F", true}, {10, "G", true}, {10, "H", true}},
			expectedSunk: true},
		{shots: []shot{{7, "J", true}, {8, "I", false}, {9, "H", false}, {10, "G", true}}, expectedSunk: false},
		{shots: []shot{{1, "H", true}, {8, "I", false}, {10, "CC", false}, {10, "G", true}}, expectedSunk: false, err: ErrIncorrectLetter},
		{shots: []shot{{1, "G", false}, {1, "H", true}, {11, "G", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
		{shots: []shot{{1, "G", false}, {1, "H", true}, {10, "P", false}}, expectedSunk: false, err: ErrIncorrectLetter},
		{shots: []shot{{1, "G", false}, {12, "K", false}, {11, "G", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			grid := NewGrid(ships)
			var latestShootResult ShootResult
			for _, shot := range test.shots {
				var err error
				latestShootResult, err = grid.Shoot(shot.num, shot.letter)

				if !cmp.Equal(shot.expectedHit, latestShootResult.Hit) {
					t.Log(cmp.Diff(shot.expectedHit, latestShootResult.Hit))
					t.Fail()
				}
				if err != nil && test.err != nil {
					if !errors.Is(err, test.err) {
						t.Log("err is incorrect")
						t.Fail()
					}
				}
			}

			if !cmp.Equal(test.expectedSunk, latestShootResult.Sunk) {
				t.Log(cmp.Diff(test.expectedSunk, latestShootResult.Sunk))
				t.Fail()
			}

			grid.ResetShips()
		})
	}
}

func getShips() []ship {
	// count  name              size
	//   1    Aircraft Carrier   5
	//   1    Battleship         4
	//   1    Cruiser            3
	//   2    Destroyer          2
	//   2    Submarine          1
	//
	// 		  A B C D E F G H I J
	//		1               @
	//		2 @             @
	//		3         @     @
	//		4               @
	//		5   @ @
	//		6
	//		7           @       @
	//		8           @       @
	//		9                   @
	//	   10       @ @ @ @ @
	//
	var ships []ship
	newship := ship{
		start: Coordinate{2, 'A'},
		end:   Coordinate{2, 'A'},
		cap:   1,
	}
	ships = append(ships, newship)
	ships = append(ships, ship{
		start: Coordinate{3, 'E'},
		end:   Coordinate{3, 'E'},
		cap:   1,
	})
	ships = append(ships, ship{
		start: Coordinate{1, 'H'},
		end:   Coordinate{4, 'H'},
		cap:   4,
	})
	ships = append(ships, ship{
		start: Coordinate{5, 'B'},
		end:   Coordinate{5, 'C'},
		cap:   2,
	})
	ships = append(ships, ship{
		start: Coordinate{7, 'F'},
		end:   Coordinate{8, 'F'},
		cap:   2,
	})
	ships = append(ships, ship{
		start: Coordinate{7, 'J'},
		end:   Coordinate{9, 'J'},
		cap:   3,
	})
	ships = append(ships, ship{
		start: Coordinate{10, 'D'},
		end:   Coordinate{10, 'H'},
		cap:   5,
	})

	return ships
}
