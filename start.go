package main

import (
	"fmt"
	"math"
	"os"
)

func doMatrixMultiplication(matrix [][]float64, vector []float64) ([]float64, error) {
	matrixCols := len(matrix[0])
	matrixRows := len(matrix)

	result := make([]float64, matrixRows)

	if matrixCols != len(vector) {
		return result, fmt.Errorf("matrix and vector have different sizes")
	}

	for i := 0; i < matrixRows; i++ {
		for j := 0; j < matrixCols; j++ {
			var innerMult = matrix[j][i] * vector[j]
			result[i] += innerMult
		}
	}

	return result, nil
}

func calcDeltaOfArrays(array1 []float64, array2 []float64) float64 {
	if array2 == nil || array1 == nil {
		return 1000
	}
	delta := 0.0

	for i := 0; i < len(array1); i++ {
		diff := math.Pow(array1[i]-array2[i], 2)
		delta += diff
	}
	return delta
}

func main() {

	// input data
	var epsilon = 0.2
	var adjacencyMatrix = [][]int{
		{0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	}

	/*
		// Lecture data
		var epsilon = 0.2

		var adjacencyMatrix = [][]int{
			{0, 1, 0, 1},
			{0, 0, 0, 1},
			{1, 0, 0, 1},
			{0, 0, 1, 0},
		}
	*/
	calculation, err := runCalculation(adjacencyMatrix, epsilon)
	if err != nil {
		return
	}
	fmt.Println("Calculation:", calculation)
}

func runCalculation(adjacencyMatrix [][]int, epsilon float64) (counterToReturn int, errorReturned error) {

	// power matrix - contains probabilities of transition from one node to another
	var powerMatrix = make([][]float64, len(adjacencyMatrix))

	// number of nodes
	var countOfNodes = len(adjacencyMatrix)
	// pi 0
	piprev := make([]float64, countOfNodes)

	for i := 0; i < countOfNodes; i++ {
		piprev[i] = 0.2
	}

	// probability of selecting a node without an edge (jump to random node)
	var noEdge = epsilon / float64(countOfNodes)

	for i := 0; i < countOfNodes; i++ {
		powerMatrix[i] = make([]float64, countOfNodes)
		for j := 0; j < countOfNodes; j++ {
			if adjacencyMatrix[i][j] == 1 {
				powerMatrix[i][j] = noEdge + (1-epsilon)/float64(getOut(adjacencyMatrix[i]))
			} else {
				powerMatrix[i][j] = noEdge
			}
		}
	}

	//pi 1
	var pinext []float64
	var err error
	var counter = 0

	for calcDeltaOfArrays(piprev, pinext) > math.Pow(10, -5) {
		if pinext != nil {
			piprev = pinext
		}
		pinext, err = doMatrixMultiplication(powerMatrix, piprev)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printVecWithTwoDecimals(pinext, counter)

		counter++
	}
	return counter, err
}

func printVecWithTwoDecimals(array []float64, counter int) {
	print(counter, ": ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%.2f, ", math.Round(array[i]*100)/100)
	}
	fmt.Println()
}

func getOut(vec []int) int {
	var sumOfVertices int
	for _, element := range vec {
		if element == 1 {
			sumOfVertices += 1
		}
	}
	return sumOfVertices
}
