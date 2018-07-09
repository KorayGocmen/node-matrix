package main

import (
	"testing"
)

func TestSuccessfulCreateMatrix(t *testing.T) {
	rowCount1 := 2
	colCount1 := 2
	x1, _ := New(rowCount1, colCount1, []float64{
		4, -1,
		0, 5,
	})
	if x1.RowCount != rowCount1 || x1.ColCount != colCount1 {
		t.Errorf("row or col count was wrong")
	}

	rowCount2 := 2
	colCount2 := 1
	x2, _ := New(rowCount2, colCount2, []float64{
		1,
		2,
	})
	if x2.RowCount != rowCount2 || x2.ColCount != colCount2 {
		t.Errorf("row or col count was wrong")
	}

	rowCount3 := 1
	colCount3 := 2
	x3, _ := New(rowCount3, colCount3, []float64{
		1, 2,
	})
	if x3.RowCount != rowCount3 || x3.ColCount != colCount3 {
		t.Errorf("row or col count was wrong")
	}
}

func TestUnsuccessfulCreateMatrix(t *testing.T) {
	rowCount1 := 2
	colCount1 := 3
	_, err1 := New(rowCount1, colCount1, []float64{
		4, -1,
		0, 5,
	})
	if err1 == nil {
		t.Errorf("row or col count was wrong")
	}

	rowCount2 := 1
	colCount2 := 2
	_, err2 := New(rowCount2, colCount2, []float64{
		4, 2,
		0,
	})
	if err2 == nil {
		t.Errorf("row or col count was wrong")
	}
}

func TestSuccessfulRowAt(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x2, _ := New(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})

	x1RowAtReal, _ := x1.RowAt(0)
	x1RowAtExpected := []float64{4, -1}
	if !arrayEqual(x1RowAtReal, x1RowAtExpected) {
		t.Errorf("row at error, %v, %v", x1RowAtReal, x1RowAtExpected)
	}

	x2RowAtReal, _ := x2.RowAt(0)
	x2RowAtExpected := []float64{1, 8, 0}
	if !arrayEqual(x2RowAtReal, x2RowAtExpected) {
		t.Errorf("row at error, %v, %v", x2RowAtReal, x2RowAtExpected)
	}

	x3RowAtReal, _ := x2.RowAt(1)
	x3RowAtExpected := []float64{6, -2, 3}
	if !arrayEqual(x3RowAtReal, x3RowAtExpected) {
		t.Errorf("row at error, %v, %v", x3RowAtReal, x3RowAtExpected)
	}
}

func TestUnsuccessfulRowAt(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		4, -1,
		0, 5,
	})

	_, err1 := x1.RowAt(2)
	if err1 == nil {
		t.Errorf("row at error")
	}
}

func TestSuccessfulColAt(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x2, _ := New(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})

	x1ColAtReal, _ := x1.ColAt(0)
	x1ColAtExpected := []float64{4, 0}
	if !arrayEqual(x1ColAtReal, x1ColAtExpected) {
		t.Errorf("col at error, %v, %v", x1ColAtReal, x1ColAtExpected)
	}

	x2ColAtReal, _ := x2.ColAt(0)
	x2ColAtExpected := []float64{1, 6}
	if !arrayEqual(x2ColAtReal, x2ColAtExpected) {
		t.Errorf("col at error, %v, %v", x2ColAtReal, x2ColAtExpected)
	}

	x3ColAtReal, _ := x2.ColAt(1)
	x3ColAtExpected := []float64{8, -2}
	if !arrayEqual(x3ColAtReal, x3ColAtExpected) {
		t.Errorf("col at error, %v, %v", x3ColAtReal, x3ColAtExpected)
	}
}

func TestUnsuccessfulColAt(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		4, -1,
		0, 5,
	})

	_, err1 := x1.ColAt(2)
	if err1 == nil {
		t.Errorf("row at error")
	}
}

func TestSuccessfulDotProduct(t *testing.T) {
	x1, _ := New(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	x2, _ := New(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})
	expectedResult1, _ := New(2, 2, []float64{
		58, 64,
		139, 154,
	})

	result1, err1 := Dot(x1, x2)

	if err1 != nil || !deepEqual(expectedResult1, result1) {
		t.Errorf("dot product error")
	}

	x3, _ := New(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x4, _ := New(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})
	expectedResult2, _ := New(2, 3, []float64{
		-2, 34, -3,
		30, -10, 15,
	})

	result2, err2 := Dot(x3, x4)

	if err2 != nil || !deepEqual(expectedResult2, result2) {
		t.Errorf("dot product error")
	}
}

func TestUnsuccessfulDotProduct(t *testing.T) {
	x1, _ := New(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	x2, _ := New(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	_, err1 := Dot(x1, x2)

	if err1 == nil {
		t.Errorf("dot product error")
	}
}

func TestSuccessfulScale(t *testing.T) {
	x1, _ := New(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	expectedResult1, _ := New(2, 3, []float64{
		3, 6, 9,
		15, 18, 21,
	})

	result1, err1 := Scale(3, x1)

	if err1 != nil || !deepEqual(expectedResult1, result1) {
		t.Errorf("scale error")
	}

	x2, _ := New(2, 2, []float64{
		1, 2,
		5, 6,
	})

	expectedResult2, _ := New(2, 2, []float64{
		-2, -4,
		-10, -12,
	})

	result2, err2 := Scale(-2, x2)

	if err2 != nil || !deepEqual(expectedResult2, result2) {
		t.Errorf("scale error")
	}
}
func TestSuccessfulDeepEqual(t *testing.T) {
	x1, _ := New(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	x2, _ := New(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	if !deepEqual(x1, x2) {
		t.Errorf("deep equal error")
	}

	x3, _ := New(1, 1, []float64{
		1,
	})

	x4, _ := New(1, 1, []float64{
		1,
	})

	if !deepEqual(x3, x4) {
		t.Errorf("deep equal error")
	}
}

func TestUnsuccessfulDeepEqual(t *testing.T) {
	x1, _ := New(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})

	x2, _ := New(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	if deepEqual(x1, x2) {
		t.Errorf("deep equal error")
	}

	x3, _ := New(2, 2, []float64{
		1, 2,
		3, 4,
	})

	x4, _ := New(2, 2, []float64{
		1, 2,
		3, 5,
	})

	if deepEqual(x3, x4) {
		t.Errorf("deep equal error")
	}
}

func TestSuccessfulTranspose(t *testing.T) {
	x1, _ := New(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	expectedResult1, _ := New(3, 2, []float64{
		1, 5,
		2, 6,
		3, 7,
	})

	result1, err1 := Transpose(x1)

	if err1 != nil || !deepEqual(expectedResult1, result1) {
		t.Errorf("transpose error")
	}

	x2, _ := New(2, 2, []float64{
		1, 2,
		5, 6,
	})

	expectedResult2, _ := New(2, 2, []float64{
		1, 5,
		2, 6,
	})

	result2, err2 := Transpose(x2)

	if err2 != nil || !deepEqual(expectedResult2, result2) {
		t.Errorf("transpose error")
	}
}

func TestSuccessfulAdd(t *testing.T) {
	x1, _ := New(3, 2, []float64{
		1, 3,
		1, 0,
		1, 2,
	})
	x2, _ := New(3, 2, []float64{
		0, 0,
		7, 5,
		2, 1,
	})
	expectedResult1, _ := New(3, 2, []float64{
		1, 3,
		8, 5,
		3, 3,
	})

	result1, err1 := Add(x1, x2)
	if err1 != nil || !deepEqual(expectedResult1, result1) {
		t.Errorf("add error")
	}

	x3, _ := New(3, 2, []float64{
		1, 3,
		1, 0,
		1, 2,
	})
	x4, _ := New(3, 2, []float64{
		0, 0,
		-7, -5,
		-2, -1,
	})
	expectedResult2, _ := New(3, 2, []float64{
		1, 3,
		-6, -5,
		-1, 1,
	})

	result2, err2 := Add(x3, x4)
	if err2 != nil || !deepEqual(expectedResult2, result2) {
		t.Errorf("add error")
	}
}

func TestUnsuccessfulAdd(t *testing.T) {
	x1, _ := New(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	x2, _ := New(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	_, err1 := Add(x1, x2)

	if err1 == nil {
		t.Errorf("add error")
	}
}

func TestSuccessfulSubtract(t *testing.T) {
	x1, _ := New(2, 3, []float64{
		8, 4, 2,
		6, 1, 5,
	})
	x2, _ := New(2, 3, []float64{
		3, 10, 4,
		5, 6, 1,
	})
	expectedResult1, _ := New(2, 3, []float64{
		5, -6, -2,
		1, -5, 4,
	})

	result1, err1 := Subtract(x1, x2)
	if err1 != nil || !deepEqual(expectedResult1, result1) {
		t.Errorf("subtract error")
	}

	x3, _ := New(3, 2, []float64{
		-1, 2, 0,
		0, 3, 6,
	})
	x4, _ := New(3, 2, []float64{
		0, -4, 3,
		9, -4, -3,
	})
	expectedResult2, _ := New(3, 2, []float64{
		-1, 6, -3,
		-9, 7, 9,
	})

	result2, err2 := Subtract(x3, x4)
	if err2 != nil || !deepEqual(expectedResult2, result2) {
		t.Errorf("subtract error")
	}
}

func TestUnsuccessfulSubtract(t *testing.T) {
	x1, _ := New(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	x2, _ := New(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	_, err1 := Subtract(x1, x2)

	if err1 == nil {
		t.Errorf("subtract error")
	}
}

func TestSuccessfulAddScalar(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		1, 2,
		5, 6,
	})
	expectedResult, _ := New(2, 2, []float64{
		4, 5,
		8, 9,
	})

	result, _ := AddScalar(3, x1)
	if !deepEqual(expectedResult, result) {
		t.Errorf("add scalar error")
	}

	x2, _ := New(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	expectedResult2, _ := New(2, 3, []float64{
		1.5, 2.5, 3.5,
		4.5, 5.5, 6.5,
	})

	result2, _ := AddScalar(0.5, x2)
	if !deepEqual(expectedResult2, result2) {
		t.Errorf("add scalar error")
	}
}

func TestSuccessfulSubtractScalar(t *testing.T) {
	x1, _ := New(2, 2, []float64{
		1, 2,
		5, 6,
	})
	expectedResult, _ := New(2, 2, []float64{
		0, 1,
		4, 5,
	})

	result, _ := SubtractScalar(1, x1)
	if !deepEqual(expectedResult, result) {
		t.Errorf("subtract scalar error")
	}

	x2, _ := New(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	expectedResult2, _ := New(2, 3, []float64{
		-1, 0, 1,
		2, 3, 4,
	})

	result2, _ := SubtractScalar(2, x2)
	if !deepEqual(expectedResult2, result2) {
		t.Errorf("subtract scalar error")
	}
}

func arrayEqual(arr1, arr2 []float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func deepEqual(x, y *Matrix) bool {
	if x.RowCount != y.RowCount || x.ColCount != y.ColCount {
		return false
	}

	for rowID := 0; rowID < x.RowCount; rowID++ {
		for colID := 0; colID < x.ColCount; colID++ {
			if x.Matrix[rowID][colID] != y.Matrix[rowID][colID] {
				return false
			}
		}
	}

	return true
}
