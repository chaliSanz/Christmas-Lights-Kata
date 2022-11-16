package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Intruncion struct {
	steps  []string
	matrix [][]int
}

type Position struct {
	start, end int
}

type Step struct {
	pos []Position

	option string
}

func main() {
	//fmt.Print("test")
	fmt.Println(createMatrix(2, 2))
}

func createMatrix(row, column int) [][]int {

	m := make([][]int, row)
	for i := range m {
		m[i] = make([]int, column)
	}
	return m

}

func checkIfAllLighTurnOff(matrix [][]int) bool {

	return countLight(matrix) > 0

}

func countLight(matrix [][]int) int {

	total := 0

	for i, row := range matrix {

		for j := range row {

			total = total + matrix[i][j]

		}
	}

	return total
}

func turnOnAllLight(matrix [][]int) [][]int {

	for i, row := range matrix {
		for j := range row {
			matrix[i][j] = 1
		}
	}

	return matrix

}

func checkIfAllLightTurnOn(matrix [][]int) bool {

	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 0 {
				return false
			}

		}
	}

	return true
}

func executeInstructions(intruncions []string, matrix [][]int) int {

	steps := make([]Step, len(intruncions))

	for i, current := range intruncions {

		err, step := readStep(current)
		if err != nil {
			//todo error
		}
		steps[i] = step

	}

	for _, step := range steps {
		err, mtrix := executeStep(step, matrix)
		if err != nil {
			//todo error
		}

		matrix = mtrix
	}

	return countLight(matrix)

}

func executeStep(step Step, matrix [][]int) (error, [][]int) {

	if step.option == "on" {
		return nil, turnOnLight(step.pos[0], step.pos[1], matrix)
	} else if step.option == "off" {
		return nil, turnOffLight(step.pos[0], step.pos[1], matrix)
	} else if step.option == "toggle" {
		return nil, toggleLight(step.pos[0], step.pos[1], matrix)
	} else {
		return errors.New("Options no found."), nil
	}

}

func readStep(intruncions string) (error, Step) {

	steps := Step{}

	option := getOptionFromStep(intruncions)

	err, positions := getPositionsFromStep(intruncions)

	if option == "" || err != nil {
		return errors.New("a problem as ocurred."), steps
	}

	steps.option = option
	steps.pos = positions

	return nil, steps

}

func getPositionsFromStep(step string) (error, []Position) {

	re := regexp.MustCompile("[\\d]*,[\\d]*")
	findPositions := re.FindAllString(step, -1)

	pos := make([]Position, 2)

	for i := 0; i < len(findPositions); i++ {

		positions := strings.Split(findPositions[i], ",")

		if len(positions) > 2 || len(findPositions) > 2 {
			return errors.New("cant process more that 2 pair of positions"), pos
		}
		if i == 0 {
			pos[i].start = int(parseInt(positions[0]))

			pos[i+1].start = int(parseInt(positions[1]))
		} else {
			pos[i-1].end = int(parseInt(positions[0]))

			pos[i].end = int(parseInt(positions[1]))
		}

	}
	return nil, pos

}

func getOptionFromStep(step string) string {

	finStr := strings.Contains(step, "turn")

	if finStr {
		str := strings.Split(step, " ")

		if strings.EqualFold(str[1], "on") {
			return "on"
		} else if strings.EqualFold(str[1], "off") {
			return "off"

		} else {
			return ""
		}
	} else {
		str := strings.Split(step, " ")
		if strings.EqualFold(str[0], "toggle") {
			return "toggle"
		} else {
			return ""
		}
	}

}

func turnOnLight(postionX, positionY Position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {

			if matrix[i][j] == 0 {
				matrix[i][j]++
				continue
			}

		}
	}

	return matrix

}

func turnOffLight(postionX, positionY Position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {

			if matrix[i][j] != 0 {
				matrix[i][j]--
			}

		}
	}

	return matrix

}

func toggleLight(postionX, positionY Position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {

			if matrix[i][j] == 0 || matrix[i][j] == 1 {
				matrix[i][j] = matrix[i][j] + 2
				continue
			}

		}
	}

	return matrix

}

func parseInt(text string) int64 {

	value, err := strconv.ParseInt(text, 0, 64)

	if err != nil {
		panic(err)
	}

	return value

}
