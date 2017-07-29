package chess

import "strconv"
import "fmt"
import s "strings"
import "math"
import "errors"

/*
 a b c d e
| | | | | |1
|x| | | | |2
| | |x| | |3
| | | | | |4

one axis must be 2,
one axis must be 1

*/


// CanKnightAttack checks if two knights on specified positions can attack each other
func CanKnightAttack(white, black string) (bool, error) {
	// validate inputs

	if len(white) != 2 || len(black) != 2 {
		return false, fmt.Errorf("Invalid coordinates: %s %s ", white, black)
	}

	if s.ToLower(white) == s.ToLower(black) {
		return false, fmt.Errorf("Invalid coordinates: cannot be equal %s %s ", white, black)
	}

	wcol, wrow := parseCoordinates(white)	
	bcol, brow := parseCoordinates(black)

	if wcol > 8 || bcol > 8 || wrow > 8 || brow > 8 || wcol < 1 || bcol < 1 || wrow < 1 || brow < 1 {
		return false, errors.New("Invalid coordinates: out of range")
	}

	colDif := math.Abs(float64(wcol - bcol))
	rowDif := math.Abs(float64(wrow - brow))

	// test if can attack

	return (colDif == 2 && rowDif == 1) || (colDif == 1 && rowDif == 2), nil
}

func parseCoordinates(input string) (int, int) {
	coords := s.Split(input, "")

	col, _ := strconv.ParseInt(coords[0], 25, 0)
	row, _ := strconv.Atoi(coords[1])
	
	return int(col - 9), row
}