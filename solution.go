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
	const SidesCircle myBeautifulType = 0
	const SidesTriangle myBeautifulType = 3
	const SidesSquare myBeautifulType = 4

	if sidesNum == SidesCircle {
		return 2 * math.Pi * sideLen
	} else {
		if sidesNum == SidesTriangle {
			return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
		} else {
			if sidesNum == SidesSquare {
				return math.Pow(sideLen, 2)
			} else {
				return 0
			}
		}
	}

	//switch sidesNum {
	//case SidesCircle:
	//	return 2 * math.Pi * sideLen
	//case SidesTriangle:
	//	return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	//case SidesSquare:
	//	return math.Pow(sideLen, 2)
	//default:
	//	return 0
	//}
}
