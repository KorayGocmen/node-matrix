// Package matrix implements a simple library for matrix operations.
// A matrix struct is used to create new matrixes and operations are
// performed using this structure.
package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Matrix is the main matrix data structure.
type Matrix struct {
	RowCount int         `json:"rowCount"`
	ColCount int         `json:"colCount"`
	Matrix   [][]float64 `json:"matrix"`
}

// New creates a new matrix struct with the
// given row, col count and initial values.
func New(rowCount, colCount int, values []float64) (*Matrix, error) {
	if values != nil && len(values) != rowCount*colCount {
		return nil, errors.New("initial values and dimensions of matrix do not match")
	}

	matrix := make([][]float64, rowCount)
	for i := range matrix {
		matrix[i] = make([]float64, colCount)
	}

	index := 0
	for i := range matrix {
		for j := range matrix[i] {
			ijValue := 0.0
			if values != nil {
				ijValue = values[index]
			}
			matrix[i][j] = ijValue
			index++
		}
	}

	m := Matrix{RowCount: rowCount, ColCount: colCount, Matrix: matrix}
	return &m, nil
}

// RowAt returns the row with the given row index.
func (m Matrix) RowAt(rowIndex int) ([]float64, error) {
	if rowIndex >= m.RowCount {
		return nil, errors.New("given row index is out of range in given matrix")
	}
	return m.Matrix[rowIndex], nil
}

// ColAt returns the column with the given column index.
func (m Matrix) ColAt(colIndex int) ([]float64, error) {
	var col []float64
	for i := 0; i < m.RowCount; i++ {
		if colIndex >= m.ColCount {
			return nil, errors.New("given col index is out of range in given matrix")
		}
		col = append(col, m.Matrix[i][colIndex])
	}
	return col, nil
}

// Dot does matrix multiplication and returns a new
// matrix for the result and returns it.
func Dot(x, y *Matrix) (*Matrix, error) {

	if x.ColCount != y.RowCount {
		return nil, errors.New("unable to perform dot multiplication due to dimensions of the matrices")
	}

	var wg sync.WaitGroup
	out, _ := New(x.RowCount, y.ColCount, nil)

	for rowID := 0; rowID < x.RowCount; rowID++ {
		row, _ := x.RowAt(rowID)
		wg.Add(1)
		go (func(rowID int, row []float64, out *Matrix) {
			for colID := 0; colID < y.ColCount; colID++ {
				col, _ := y.ColAt(colID)
				for i := 0; i < len(row); i++ {
					out.Matrix[rowID][colID] += row[i] * col[i]
				}
			}
			wg.Done()
		})(rowID, row, out)
	}

	wg.Wait()
	return out, nil
}

// Scale performs scalar multiplication with the given factor.
// Creates a new scaled Matrix struct and returns it.
func Scale(A float64, x *Matrix) (*Matrix, error) {
	out, _ := New(x.RowCount, x.ColCount, nil)

	for rowID := 0; rowID < x.RowCount; rowID++ {
		for colID := 0; colID < x.ColCount; colID++ {
			out.Matrix[rowID][colID] = A * x.Matrix[rowID][colID]
		}
	}

	return out, nil
}

// Transpose creates a new transpose matrix for the given matrix.
// Newly created matrix is returned.
func Transpose(x *Matrix) (*Matrix, error) {
	var outValues []float64
	for colID := 0; colID < x.ColCount; colID++ {
		col, _ := x.ColAt(colID)
		for _, val := range col {
			outValues = append(outValues, val)
		}
	}

	out, _ := New(x.ColCount, x.RowCount, outValues)
	return out, nil
}

// Add adds two matrixes.
// Returns the result in a new matrix struct.
func Add(x, y *Matrix) (*Matrix, error) {
	if x.RowCount != y.RowCount || x.ColCount != y.ColCount {
		return nil, errors.New("unable to perform matrix sum due to matrix dimensions")
	}

	out, _ := New(x.RowCount, x.ColCount, nil)
	for rowID := 0; rowID < x.RowCount; rowID++ {
		for colID := 0; colID < x.ColCount; colID++ {
			out.Matrix[rowID][colID] = x.Matrix[rowID][colID] + y.Matrix[rowID][colID]
		}
	}

	return out, nil
}

// Subtract subtracts two matrixes.
// Returns the result in a new matrix struct.
func Subtract(x, y *Matrix) (*Matrix, error) {
	minusY, _ := Scale(-1, y)
	result, err := Add(x, minusY)
	return result, err
}

// AddScalar adds a scalar value to each element in
// the matrix. Returns a new result matrix.
func AddScalar(A float64, x *Matrix) (*Matrix, error) {
	out, _ := New(x.RowCount, x.ColCount, nil)

	for rowID := 0; rowID < x.RowCount; rowID++ {
		for colID := 0; colID < x.ColCount; colID++ {
			out.Matrix[rowID][colID] = A + x.Matrix[rowID][colID]
		}
	}

	return out, nil
}

// SubtractScalar subtracts a scalar value to each element
// in the matrix. Returns a new result matrix.
func SubtractScalar(A float64, x *Matrix) (*Matrix, error) {
	out, _ := AddScalar(-1*A, x)
	return out, nil
}

func popMatrix(s []string) (string, []string) {
	return s[0], s[1:]
}

func (m *Matrix) printableMatrix() string {
	matrix := "["
	for i, row := range m.Matrix {
		matrix += "["
		for j, val := range row {
			matrix += strconv.FormatFloat(val, 'f', 5, 64)
			if j != m.ColCount-1 {
				matrix += ","
			}
		}
		matrix += "]"
		if i != m.RowCount-1 {
			matrix += ","
		}
	}
	matrix += "]"
	return matrix
}

func main() {
	command := flag.String("command", "_UNKNOWN_", "command name")
	input1Raw := flag.String("input1", "_UNKNOWN_", "command inputs")
	input2Raw := flag.String("input2", "_UNKNOWN_", "command inputs")
	flag.Parse()

	switch *command {
	case "dot":
		{
			input1 := strings.Split(*input1Raw, ",")
			input2 := strings.Split(*input2Raw, ",")

			var matrix1RowRaw, matrix1ColRaw string
			matrix1RowRaw, input1 = popMatrix(input1)
			matrix1Row, _ := strconv.Atoi(matrix1RowRaw)
			matrix1ColRaw, input1 = popMatrix(input1)
			matrix1Col, _ := strconv.Atoi(matrix1ColRaw)

			var matrix1Values []float64
			for i := 0; i < (matrix1Row * matrix1Col); i++ {
				var valRaw string
				valRaw, input1 = popMatrix(input1)
				val, _ := strconv.ParseFloat(valRaw, 64)
				matrix1Values = append(matrix1Values, val)
			}

			matrix1, _ := New(matrix1Row, matrix1Col, matrix1Values)

			var matrix2RowRaw, matrix2ColRaw string
			matrix2RowRaw, input2 = popMatrix(input2)
			matrix2Row, _ := strconv.Atoi(matrix2RowRaw)
			matrix2ColRaw, input2 = popMatrix(input2)
			matrix2Col, _ := strconv.Atoi(matrix2ColRaw)

			var matrix2Values []float64
			for i := 0; i < (matrix2Row * matrix2Col); i++ {
				var valRaw string
				valRaw, input2 = popMatrix(input2)
				val, _ := strconv.ParseFloat(valRaw, 64)
				matrix2Values = append(matrix2Values, val)
			}

			matrix2, _ := New(matrix2Row, matrix2Col, matrix2Values)

			dotProduct, err := Dot(matrix1, matrix2)
			if err != nil {
				fmt.Println(false, err)
				fmt.Printf("{\"success\":false,\"err\":%s}\n", err.Error())
			} else {
				resultMatrix := dotProduct.printableMatrix()
				fmt.Printf("{\"success\":true,\"out\":{\"rowCount\":%d,\"colCount\":%d,\"matrix\":%s}}\n", dotProduct.RowCount, dotProduct.ColCount, resultMatrix)
			}
		}
	}
}
