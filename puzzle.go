package main

import (
	"fmt"
	"time"
)

func findZero(puzzle [][]int) (x int, y int) {

	for y, i := range(puzzle) {
		for x, j := range(i) {
			if j == 0 {
				return x, y
				break;
			}
		}
	}

	return -1, -1

}

func checkIfDirectionAvailable(puzzle [][]int, direction int) bool {

	x, y := findZero(puzzle)

	switch direction {
		case 1:
			if y > 0 {
				return true
			}
			break
		case 2:
			if x < len(puzzle[0]) - 1 {
				return true
			}
			break
		case 3:
			if y < len(puzzle) - 1 {
				return true
			}
			break
		case 4:
			if x > 0 {
				return true
			}
			break
		default:
			return false
			break
	}

	return false

}

func move(puzzle [][]int, direction int) [][]int {

	//ifDirectionAvailable := checkIfDirectionAvailable(puzzle, direction)
	x, y := findZero(puzzle)

	var tempPuzzle = make([][]int, 0)

	for _, sv := range(puzzle) {
		var temp []int
		for _, v := range(sv) {
			temp = append(temp, v)
		}
		tempPuzzle = append(tempPuzzle, temp)
	}
	
	//if ifDirectionAvailable {
	
		switch direction {
			case 1:
				tempPuzzle[y][x] = tempPuzzle[y-1][x]
				tempPuzzle[y-1][x] = 0
				break
			case 2:
				tempPuzzle[y][x] = tempPuzzle[y][x+1]
				tempPuzzle[y][x+1] = 0
				break
			case 3:
				tempPuzzle[y][x] = tempPuzzle[y+1][x]
				tempPuzzle[y+1][x] = 0
				break
			case 4:
				tempPuzzle[y][x] = tempPuzzle[y][x-1]
				tempPuzzle[y][x-1] = 0
			default:
				break
		}
	
	//}

	return tempPuzzle

}

func getAvailableDirections(puzzle [][]int) []int {

	var directions []int = make([]int, 0)

	for i:=1; i<=4; i++ {
		if checkIfDirectionAvailable(puzzle, i) {
			directions = append(directions, i)
		}
	}

	return directions

}

func getAvailableDirectionCoordinates(puzzle [][]int) [][2]int {

	var availableDirectionCoordinate [][2]int = make([][2]int, 0)

	x, y := findZero(puzzle)

	if checkIfDirectionAvailable(puzzle, 1) {
		availableDirectionCoordinate = append(availableDirectionCoordinate, [2]int{x,y-1})
	}
	if checkIfDirectionAvailable(puzzle, 2) {
		availableDirectionCoordinate = append(availableDirectionCoordinate, [2]int{x+1,y})
	}
	if checkIfDirectionAvailable(puzzle, 3) {
		availableDirectionCoordinate = append(availableDirectionCoordinate, [2]int{x,y+1})
	}
	if checkIfDirectionAvailable(puzzle, 4) {
		availableDirectionCoordinate = append(availableDirectionCoordinate, [2]int{x-1,y})
	}

	return availableDirectionCoordinate

}

func checkIfSolve(puzzle [][]int) bool {

	var currentNumber int = 1

	for y, sy := range(puzzle) {
		for x, number := range(sy) {
			if number != currentNumber && (y != len(puzzle) - 1 || x != len(puzzle[0]) - 1) {
				return false
				break
			}
			currentNumber++
		}
	}

	return true

}

func makeStep(stepId int, lastStepId int, direction int) [3]int {

	return [3]int{stepId, lastStepId, direction}

}

func traceSteps(allSteps [][3]int) []int {

	var solvingDirections []int
	var currentStep [3]int

	solvingDirections = append(solvingDirections, allSteps[len(allSteps)-1][2])
	currentStep = allSteps[len(allSteps)-1]

	if len(allSteps) > 1 {
		for i:=len(allSteps)-2; i>=0; i-- {
			if (allSteps[i][0] == currentStep[1]) {
				solvingDirections = append(solvingDirections, allSteps[i][2])
				currentStep = allSteps[i]
			}
		}
	}

    for i, j:=0, len(solvingDirections)-1; i<j; i, j=i+1, j-1 {
        solvingDirections[i], solvingDirections[j] = solvingDirections[j], solvingDirections[i]
    }

	return solvingDirections

}

func haveEverMoved(allPuzzles [][][]int, currentPuzzle [][]int) bool {

	for _, puzzle := range(allPuzzles) {
		if puzzleEqual(puzzle, currentPuzzle) {
			return true
			break
		}
	}

	return false

}

func puzzleEqual(puzzleA [][]int, puzzleB[][]int) bool {

	for sk, sv := range(puzzleA) {
		for k, v := range(sv) {
			if (v != puzzleB[sk][k]) {
				return false
			}
		}
	}

	return true

}

func solve(puzzle [][]int) []int {

	var allSteps [][3]int
	var stepsQueue [][3]int
	var allPuzzles [][][]int
	var puzzleQueue [][][]int

	var stepId int = 1
	var stepCounting int = 0

	var directions []int

	allPuzzles = append(allPuzzles, puzzle)

	directions = getAvailableDirections(puzzle)

	for _, direction := range(directions) {
		forwardPuzzle := move(puzzle, direction)
		step := makeStep(stepId, 0, direction)
		stepsQueue = append(stepsQueue, step)
		allSteps = append(allSteps, step)
		stepCounting++
		if checkIfSolve(forwardPuzzle) {
			return traceSteps(allSteps)
			break
		}
		puzzleQueue = append(puzzleQueue, forwardPuzzle)
		allPuzzles = append(allPuzzles, forwardPuzzle)
		stepId++
	}

	var currentPuzzle [][]int
	var currentStep [3]int

	for len(puzzleQueue) > 0 {

		currentPuzzle, puzzleQueue = puzzleQueue[0], puzzleQueue[1:]
		currentStep, stepsQueue = stepsQueue[0], stepsQueue[1:]

		directions = getAvailableDirections(currentPuzzle)
		for _, direction := range(directions) {
			forwardPuzzle := move(currentPuzzle, direction)
			if haveEverMoved(allPuzzles, forwardPuzzle) {
				continue
			}
			step := makeStep(stepId, currentStep[0], direction)
			stepsQueue = append(stepsQueue, step)
			allSteps = append(allSteps, step)
			stepCounting++
			if checkIfSolve(forwardPuzzle) {
				fmt.Printf("共计计算了%d步\n", stepCounting)
				return traceSteps(allSteps)
				break
			}
			puzzleQueue = append(puzzleQueue, forwardPuzzle)
			allPuzzles = append(allPuzzles, forwardPuzzle)
			stepId++
		}

	}

	return make([]int, 0)

}

func printPuzzleQueue(puzzleQueue [][][]int) {
	for _, puzzle := range(puzzleQueue) {
		printPuzzle(puzzle)
	}
}

func printPath(path []int) {
	for _, direction := range(path) {
		switch direction {
			case 1:
				fmt.Print("上")
				break
			case 2:
				fmt.Print("右")
				break
			case 3:
				fmt.Print("下")
				break
			case 4:
				fmt.Print("左")
				break
			default:
				break
		}
	}
}

func printPuzzle(puzzle [][]int) {
	for _, sy := range(puzzle) {
		for _, number := range(sy) {
			fmt.Print(number)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main () {

    start := time.Now()

	var puzzle = [][]int{
		{5,0,3},
		{2,1,6},
		{8,4,7}}

	path := solve(puzzle)
	printPath(path)

	elapsed := time.Since(start)
	fmt.Printf("\n耗时：%s", elapsed)

}