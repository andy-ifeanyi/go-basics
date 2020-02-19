package main

import (
	"fmt"
	"math"
)

func main() {
	// boolean operations
	fmt.Println("!false =", !false)
	fmt.Println("!true =", !true)
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && false =", false && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || true)
	fmt.Println("false || true =", true || true)
	fmt.Println("false || false =", true || true)
	fmt.Println("7 > 9", 7 > 9)
	fmt.Println(`"hi" == "hey"`, "hi" == "hey")

	// integer operations
	fmt.Println(7)
	fmt.Println("4 + 14 =", 4+14)
	fmt.Println("4 + 14 =", 4+14)
	fmt.Println("-17 =", -17)
	fmt.Println("-170 + 62 =", -170+62)
	fmt.Println("-10 - -2 =", -10 - -2)
	fmt.Println("-10 - (-2) =", -10-(-2))
	fmt.Println("-11 / 3 =", -11/3)

	// floating point operations
	fmt.Println("18.0 / 36 =", 18.0/36)
	fmt.Println("Pi =", math.Pi)

	// complex number operations
	fmt.Println("14 + 3i =", 14+3i)
	fmt.Println("3i =", 3i)
}
