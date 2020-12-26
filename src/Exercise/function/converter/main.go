package main

import (
	"fmt"
)

func main() {
	radians := Pi / 4
	degrees := rad2deg(radians)
	fmt.Printf("radian: %g, degree: %g\n", degrees, radians)
	radians = deg2rad(degrees)
	fmt.Printf("radian: %g, degree: %g\n", degrees, radians)
}

func rad2deg(radian float64) float64 {
	return (radian * (180 / pi))
}

//Named return variables
func deg2rad(degree float64) (radian float64) {
	radian = degree * pi / 180.0
	return
}

/*exercise:
1. import value of Pi from cognologix.com/mathconst package
*/
