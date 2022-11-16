package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMatrix(t *testing.T) {

	expected := [][]int{{0, 0}, {0, 0}}
	got := createMatrix(2, 2)

	assert.Equal(t, expected, got)

}

func TestCheckIfAllLightTurnOff(t *testing.T) {

	expected := false

	matrix := createMatrix(10, 10)
	got := checkIfAllLighTurnOff(matrix)

	assert.Equal(t, expected, got, "all light are turn off")

}

func TestCheckifAllLightTurnOffFailed(t *testing.T) {

	expected := true
	matrix := createMatrix(10, 10)
	got := checkIfAllLighTurnOff(matrix)

	assert.NotEqual(t, expected, got, "all light are not turn off")
}

func TestCheckAllLightTurnOn(t *testing.T) {
	expected := true

	matrix := createMatrix(10, 10)
	matrix = turnOnAllLight(matrix)
	got := checkIfAllLightTurnOn(matrix)

	assert.Equal(t, expected, got)
}

func TestTurnOnLigh(t *testing.T) {

	expected := 1000000

	matrix := createMatrix(1000, 1000)

	var positionX, positionY Position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 999

	matrix = turnOnLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)
}

func TestTurnOffLight(t *testing.T) {
	expected := 999996

	matrix := createMatrix(1000, 1000)
	matrix = turnOnAllLight(matrix)

	var positionX, positionY Position

	positionX.start = 499
	positionX.end = 500
	positionY.start = 499
	positionY.end = 500

	matrix = turnOffLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)
}

func TestToggleLightOffToOn(t *testing.T) {

	expected := 1000
	matrix := createMatrix(1000, 1000)

	var positionX, positionY Position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 0

	matrix = toggleLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)

}

func TestToggleLightOnToOff(t *testing.T) {

	expected := 999000
	matrix := createMatrix(1000, 1000)
	matrix = turnOnAllLight(matrix)
	var positionX, positionY Position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 0

	matrix = toggleLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)

}

func TestParseStringToInteger(t *testing.T) {

	var expected int64 = 282

	got := parseInt("282")

	assert.Equal(t, expected, got)

}

func TestGetPositionsFromStep(t *testing.T) {

	positions := make([]Position, 2)

	positions[0].start = 887
	positions[0].end = 959
	positions[1].start = 9
	positions[1].end = 629

	err, got := getPositionsFromStep("turn on 887,9 through 959,629")

	assert.Equal(t, nil, err)
	assert.Equal(t, positions, got)

}

func TestGetPositionsFromStepMoreThan2Positions(t *testing.T) {

	positions := make([]Position, 2)
	expectedErrorMsg := "cant process more that 2 pair of positions"

	err, got := getPositionsFromStep("turn on 887,9 through 959,629 jfj 938,394")

	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)

	assert.Equal(t, positions, got)

}

func TestGetOptionFromStep(t *testing.T) {

	expecte := "on"

	got := getOptionFromStep("turn oN 887,9 through 959,629 ")

	assert.Equal(t, expecte, got)

}

func TestReadStep(t *testing.T) {

	positions := make([]Position, 2)

	positions[0].start = 887
	positions[0].end = 959
	positions[1].start = 9
	positions[1].end = 629

	var steps Step
	steps.option = "on"
	steps.pos = positions

	err, got := readStep("turturn on 887,9 through 959,629")

	assert.Equal(t, nil, err)
	assert.Equal(t, steps, got)

}

func TestexecuteStep(t *testing.T) {
	expected := 4

	positions := make([]Position, 2)

	positions[0].start = 0
	positions[0].end = 1
	positions[1].start = 0
	positions[1].end = 1

	var steps Step
	steps.option = "on"
	steps.pos = positions
	matrix := createMatrix(2, 2)
	err, mtrix := executeStep(steps, matrix)

	got := countLight(mtrix)

	assert.Equal(t, nil, err)

	assert.Equal(t, expected, got)

}

func TestexecuteStepFailed(t *testing.T) {
	expectedErrorMsg := "Options no found"

	positions := make([]Position, 2)

	positions[0].start = 0
	positions[0].end = 1
	positions[1].start = 0
	positions[1].end = 1

	var steps Step
	steps.option = "ons"
	steps.pos = positions
	matrix := createMatrix(2, 2)
	err, mtrix := executeStep(steps, matrix)

	got := countLight(mtrix)

	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	assert.Equal(t, nil, got)

}

func TestExecuteIntrustions(t *testing.T) {
	expected := 230022
	intruncions := []string{
		"turn on 887,9 through 959,629",
		"turn on 454,398 through 844,448",
		"turn off 539,243 through 559,965",
		"turn off 370,819 through 676,868",
		"turn off 145,40 through 370,997",
		"turn off 301,3 through 808,453",
		"turn on 351,678 through 951,908",
		"toggle 720,196 through 897,994",
		"toggle 831,394 through 904,860"}
	matrix := createMatrix(1000, 1000)

	got := executeInstructions(intruncions, matrix)

	assert.Equal(t, expected, got)

}
