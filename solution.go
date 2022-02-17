package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myBeautifulType int

func CalcSquare(sideLen float64, sidesNum myBeautifulType) float64 {
	switch sidesNum {
	case 0:
		SidesCircle := (2 * math.Pi * sideLen)
		return SidesCircle
	case 3:
		SidesTriangle := (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
		return SidesTriangle
	case 4:
		SidesSquare := math.Pow(sideLen, 2)
		return SidesSquare
	default:
		return 0
	}
}
