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
			println(matrix[i][j], "*", vector[j])
			result[i] += matrix[i][j] * vector[j]
		}
	}

	return result, nil
}

func calcDeltaOfArrays(array1 []float64, array2 []float64) float64 {
	if array2 == nil {
		return 1000
	}
	var sumOfArray1 float64
	var sumOfArray2 float64
	for i := 0; i < len(array1); i++ {
		sumOfArray1 += array1[i]
		sumOfArray2 += array2[i]
	}
	return math.Abs(sumOfArray1 - sumOfArray2)
}

func main() {

	// input data
	/*var epsilon = 0.2
	var adjacencyMatrix = [][]int{
		{0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	}*/

	var epsilon = 0.2

	var adjacencyMatrix = [][]int{
		{0, 1, 0, 1},
		{0, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 1, 0},
	}

	// power matrix - contains probabilities of transition from one node to another
	var powerMatrix = make([][]float64, len(adjacencyMatrix))

	// number of nodes
	var countOfNodes = len(adjacencyMatrix)

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

	fmt.Println("Power matrix: ", powerMatrix)

	// pi 0
	piprev := []float64{0.25, 0.25, 0.25, 0.25}
	//pi 1

	// pi n+1
	var pinext, err = doMatrixMultiplication(powerMatrix, piprev)

	fmt.Println("pi n+1: ", pinext)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	piprev = pinext
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
