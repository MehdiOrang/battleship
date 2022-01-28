package tests

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
}

type ShootResult struct {
	Hit  bool
	Sunk bool
}

func NewGrid(ships []Position) *Grid {
	var realShips []*ship
	for _, p := range ships {
		ship := ship{
			start: p.Start,
			end:   p.End,
		}
		realShips = append(realShips, &ship)
	}

	grid := Grid{ships: realShips}
	return &grid
}
func (grid *Grid) Shoot(shotNum int, shotLetter string) (ShootResult, error) {
	//TODO: implement here
	return ShootResult{}, nil
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
		{shots: []shot{{1, "G", false}, {1, "H", true}, {10, "P", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
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

func getShips() []Position {
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
	var ships []Position
	ships = append(ships, Position{
		Start: Coordinate{2, 'A'},
		End:   Coordinate{2, 'A'},
	})
	ships = append(ships, Position{
		Start: Coordinate{3, 'E'},
		End:   Coordinate{3, 'E'},
	})
	ships = append(ships, Position{
		Start: Coordinate{1, 'H'},
		End:   Coordinate{4, 'H'},
	})
	ships = append(ships, Position{
		Start: Coordinate{5, 'B'},
		End:   Coordinate{5, 'C'},
	})
	ships = append(ships, Position{
		Start: Coordinate{7, 'F'},
		End:   Coordinate{8, 'F'},
	})
	ships = append(ships, Position{
		Start: Coordinate{7, 'I'},
		End:   Coordinate{9, 'I'},
	})
	ships = append(ships, Position{
		Start: Coordinate{10, 'D'},
		End:   Coordinate{10, 'H'},
	})

	return ships
}
