package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/nddq/aoc2023/common"
)

func getDistanceTravel(timePressed, timeAllowed int) int {
	return 1 * timePressed * (timeAllowed - timePressed)
}

func solveQuadratic(a, b, c float64) (float64, float64, bool) {
	// Calculate the discriminant
	delta := b*b - 4*a*c

	// Check if the discriminant is negative (no real solutions)
	if delta < 0 {
		return 0, 0, false
	}

	// Calculate the solutions using the quadratic formula
	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	return x1, x2, true
}

func SolveOne() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	total := 1
	races := common.ParseBytesToStringSlices(input)
	for _, race := range races {
		waysToBeat := 0
		race := strings.Split(race, " ")
		timeAllowed, err := strconv.Atoi(race[0])
		if err != nil {
			panic(err)
		}
		distanceToBeat, err := strconv.Atoi(race[1])
		if err != nil {
			panic(err)
		}
		for timePressed := 0; timePressed <= timeAllowed; timePressed++ {
			distanceTravelled := getDistanceTravel(timePressed, timeAllowed)
			if distanceTravelled > distanceToBeat {
				waysToBeat++
			}
		}
		if waysToBeat > 0 {
			total *= waysToBeat
		}
	}
	fmt.Println(total)
}

// distance = speed * time
// speed = timePressed
// time = totalTime - timePressed
// distance = timePressed * (totalTime-timePress)
// distance = timePressed*totalTime - timePress^2
// xt - x^2 > 200
// We can model the problem as a quadratic function
// x^2−xt+C<0 where x is the time pressed, t is the total time, and C is the distance to beat
func SolveTwo() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	races := common.ParseBytesToStringSlices(input)
	for _, race := range races {
		race := strings.Split(race, " ")
		timeAllowed, err := strconv.ParseFloat(race[0], 64)
		if err != nil {
			panic(err)
		}
		distanceToBeat, err := strconv.ParseFloat(race[1], 64)
		if err != nil {
			panic(err)
		}
		a := float64(1)
		b := -timeAllowed
		c := distanceToBeat
		// Call the function to solve the quadratic equation
		x1, x2, hasRealSolutions := solveQuadratic(a, b, c)

		if hasRealSolutions {
			fmt.Printf("The solutions are x1 = %.2f and x2 = %.2f\n", x1, x2)
		} else {
			panic("No real solutions (discriminant is negative)")
		}
		// round x1 down
		// round x2 up
		x1 = math.Floor(x1)
		x2 = math.Ceil(x2)
		fmt.Println(int64((x1 - x2) + 1))
	}
}

func main() {
	SolveTwo()
}
