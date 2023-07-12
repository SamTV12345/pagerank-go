package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetOut(t *testing.T) {

	got := getOut([]int{0, 0, 1, 1, 0, 0, 1})
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetOutWithLargerArray(t *testing.T) {

	got := getOut([]int{0, 1, 0, 1, 0, 0, 1, 0, 0, 1})
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDoMatrixMultiplicationSmallNumbers(t *testing.T) {
	matrix := [][]float64{
		{0.05, 0.45, 0.05, 0.45},
		{0.05, 0.05, 0.05, 0.85},
		{0.45, 0.05, 0.05, 0.45},
		{0.05, 0.05, 0.85, 0.05},
	}

	vector := []float64{0.25, 0.25, 0.25, 0.25}

	got, _ := doMatrixMultiplication(matrix, vector)
	fmt.Println(got)
	want := []float64{0.15, 0.15, 0.25, 0.45}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("Not equal")
	}
}

func TestCalcDeltaOfSameArray(t *testing.T) {
	var array1 = []float64{1, 2, 3}

	var res = calcDeltaOfArrays(array1, array1)

	if res != 0 {
		t.Errorf("An identical array should have delta of 0")
	}
}
