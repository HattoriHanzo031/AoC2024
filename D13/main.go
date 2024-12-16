package main

import (
	"common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	scanner, close := common.FileScaner("input.txt")
	defer close()

	totalCoinsP1 := 0
	totalCoinsP2 := 0
	for scanner.Scan() {
		buttonA := strings.Fields(scanner.Text())
		Ax := common.Must(strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(buttonA[2], "X+"), ",")))
		Ay := common.Must(strconv.Atoi(strings.TrimPrefix(buttonA[3], "Y+")))
		scanner.Scan()
		buttonB := strings.Fields(scanner.Text())
		Bx := common.Must(strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(buttonB[2], "X+"), ",")))
		By := common.Must(strconv.Atoi(strings.TrimPrefix(buttonB[3], "Y+")))
		scanner.Scan()
		prize := strings.Fields(scanner.Text())
		Px := common.Must(strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(prize[1], "X="), ",")))
		Py := common.Must(strconv.Atoi(strings.TrimPrefix(prize[2], "Y=")))

		scanner.Scan()

		totalCoinsP1 += calculateCoins(Ax, Ay, Bx, By, Px, Py)
		totalCoinsP2 += calculateCoins(Ax, Ay, Bx, By, Px+10000000000000, Py+10000000000000)
	}
	fmt.Println("total coins:", totalCoinsP1)
	fmt.Println("total coins:", totalCoinsP2)
}

func calculateCoins(Ax, Ay, Bx, By, Px, Py int) int {
	bPresses := int(math.Round(float64(Ax) / float64((By*Ax)-(Ay*Bx)) * (float64(Py) - (float64(Px*Ay) / float64(Ax)))))
	aPresses := int(math.Round(float64(Px-(Bx*bPresses)) / float64(Ax)))
	if aPresses*Ax+bPresses*Bx != Px || aPresses*Ay+bPresses*By != Py {
		return 0
	}
	return 3*aPresses + bPresses
}
