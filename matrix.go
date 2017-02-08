// Matrix Struct Definition

type matrix struct {
	numRows, numCols int
	elems [][]float64
}


// brief: Parameterized constructor that takes slice 
// of slices of floats
//
// details: Allows us to create a matrix with only 
// providing slices and not specifying columns or rows
//
// note: It is on the user to make sure that all columns have 
// the same length, otherwise the matrix will break
//
// returns: a pointer to a matrix
func NewMatrix(slices... []float64) *matrix {
	m := new(matrix)
	m.numRows = len(elems)
	m.numCols = len(elems[0])
	m.elems = slices
	return m
}



// brief: Parameterized constructor that takes slice 
// of slices of floats
//
// details: Allows us to create a matrix with only 
// providing slices and not specifying columns or rows
//
// note: It is on the user to make sure that all columns have 
// the same length, otherwise the matrix will break
//
// returns: a pointer to a matrix
func BlankMatrix(rows, cols int) *matrix {
	m := new(matrix)
	m.numRows = rows
	m.numCols = cols

	// make empty slice of slices
	slices := make([][]float64, rows)
	for i:= range slices {
		slices[i] = make([]float64, cols)
	}
	m.elems = slices

	return m
}