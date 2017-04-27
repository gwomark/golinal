package matrix

import (
	"errors";
	"math";
	"fmt";
	"github.com/gonum/matrix/mat64";
)

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
func Matrix(slices... []float64) (*matrix) {
	m := new(matrix)
	m.numRows = int(len(slices))
	m.numCols = int(len(slices[0]))
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
func BlankMatrix(rows, cols int) (*matrix) {

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


func Identity(n int) (*matrix) {

	m := BlankMatrix(n, n)
	for i := 0; i < n; i++ {
		m.elems[i][i] = 1
	}

	return m
}


func (m matrix) Dims() (int, int) {
	return m.numRows, m.numCols
}


// brief: Get the row,col'th entry of a matrix
//
// note: it is undefined behavior to use invalid indices with Get()
//
// returns: number of cols
func (m matrix) At(row, col int) (float64) {

	return m.elems[row][col]
}


// brief: Gets number of rows in a matrix
//
// returns: number of rows
func (m matrix) NumRows() int {
	return m.numRows
}


// brief: Gets number of cols in a matrix
//
// returns: number of cols
func (m matrix) NumCols() int {
	return m.numCols
}


// brief: Adds two matrices together
// 
// inputs: a matrix pointer
//
// returns: an error if dimensions of the 
//          matrices to be summed aren't equal
func (m *matrix) Add(q *matrix) error {
	if (m.numRows != q.numRows) || (m.numCols != q.numRows) {
		return errors.New("Dimensions aren't equal")
	} else {
		// Loop through each entry, store sum of entries in m
		for i := 0; i < m.numRows; i++ {
			for j := 0; j < q.numRows; j++ {
				m.elems[i][j] += q.elems[i][j]
			}
		}
		return nil
	}
}


// brief: Scales a matrix by a real number
// 
// inputs: A float 
func (m *matrix) Scale(x float64) {
	for i := 0; i < m.numRows; i++ {
			for j := 0; j < m.numRows; j++ {
				m.elems[i][j] *= x
			}
		}	
}


// brief: Multiplys the matrix m by the matrix q
//
// details: O(n^3)
// 
// returns: product of m and q
func (m matrix) Multiply(q *matrix) (*matrix, error) {

	if (m.numCols != q.numRows) {
		return nil,errors.New("Dimensions can't be multiplied")
	} else {
		result := BlankMatrix(m.numRows, q.numCols)
		for i := 0; i < m.numRows; i++ {
			for j := 0; j < q.numCols; j++ {
				for k := 0; k < q.numRows; k++ {
					result.elems[i][j] += m.elems[i][k] * q.elems[k][j]
				}
			}
		}

		return result, nil	
	}
}

// brief: Calculates transpose of matrix, wrapper for T()
//
// details: Implemented for *matrix
// 
// returns: a transposed version of m as a pointer to a matrix

func(m *matrix) Transpose() *matrix {
	transpose := BlankMatrix(m.numCols, m.numRows)

	for i := 0; i < m.numCols; i++ {
		for j := 0; j < m.numRows; j++ {
			fmt.Printf("i= %d\n", i)
			fmt.Printf("j= %d\n", j)
			transpose.elems[i][j] = m.elems[j][i]
		}
	}

	return transpose
}



// brief: Calculates transpose of matrix
//
// details: Implemented for mat64.Matrix interface
// 
// returns: a transposed version of m
func(m *matrix) Inverse() (*matrix, error) {
	if !m.IsSqaure(){
		return nil, errors.New("Matrix is not square")
	}

	inverseT := BlankMatrix(m.numRows,m.numRows)
	identity := Identity(m.numRows)

	for i, _ := range inverseT.elems {
		r, err := m.Gauss(identity.elems[i])
		if err != nil {
			return nil, err
		}

		fmt.Println(r)
		inverseT.elems[i] = r
		
	}

	fmt.Println(inverseT.elems)

	return inverseT, nil
}


// brief: Calculates determinant of a matrix
//
// details: Uses LU decomposition, O(n^3)
// 
// returns: a determinant of m
func (m *matrix) Determinant() (det float64, err error) {
	L, U, P, err := m.LUP()
	if err != nil {
		return 0.0, err
	}

	det, err = determinant(*L, *U, *P)
	return
}



// brief: Calculates the LUP decomposition of a matrix
//
// details: Decomposes m in to the product of three matrices:
//              P: A row permutation matrix
//              U: An upper triangular matrix
//              L: A lower triangular matrix
//          Hence Pm = LU,
//          O(n^3)
//
// returns: a determinant of m
func (m *matrix) LUP() (*matrix, *matrix, *matrix, error) {
	
	// No LUP if matrix isn't square
	if !m.IsSqaure() {
		return nil, nil, nil, errors.New("LUP requires square matrix")
	}

	n := m.numRows
	L := BlankMatrix(n,n)
	U := BlankMatrix(n,n)
	P := m.pivotMatrix()
	Pm, _ := P.Multiply(m)

	

	for j := 0; j < n; j++ {
		
		L.elems[j][j] = 1

		// Populate U using the following formula (compile it in LaTex)
		// u_{ij} = a_{ij} - \sum_{k=1}^{i-1} u_{kj} l_{ik} 
		for i := 0; i < j+1; i++ {
			sum := 0.0
			for k := 0; k < i; k++ {
				sum += (U.elems[k][j] * L.elems[i][k])
			}
			
			U.elems[i][j] = Pm.elems[i][j] - sum	

		}
		
		
		// Populate L using the following formula (compile it in LaTex)
		// l_{ij} = \frac{1}{u_{jj}} (a_{ij} - \sum_{k=1}^{j-1} u_{kj} l_{ik}
		for i := j; i < n; i++ {
			sum := 0.0

			for k := 0; k < j; k++ {
				sum += (U.elems[k][j] * L.elems[i][k])
			}
			L.elems[i][j] = (Pm.elems[i][j] - sum) / U.elems[j][j]
		}
	}

	return L, U, P, nil

}


func (m *matrix) RREF() *matrix {
	
	// loop through each row
	for i := 0; i < m.numRows; i++ {
		
		for j := 0; j < m.numCols; j++ {

		}

	}
	return BlankMatrix(4, 4)

}



// brief: Solves equations of the form
//       Ax = b
// where A is a matrix, and x,b are vectors
//
// inputs: b a slice of floats to solve for
//
// details: If A is square, LU decomposition is used
//
func (A *matrix) Gauss(b []float64) ([]float64, error) {
	
	
	L, U, P, err := A.LUP()
	if err != nil {
		return nil, err
	}

	return gauss(b, L, U, P)

}


// brief: Finds the eigenvalues of a square matrix m
//      
//
// returns: a slice of complex numbers
func (m *matrix) Eigenvalues() ([]complex128, error) {
	if m.IsSqaure() == false {
		err := errors.New("Matrix should be square")
		return nil, err
	}

	// Create an Eigen type 
	var eigen mat64.Eigen

	// Perform eigenvalue decomposition
	eigen.Factorize(*m, true)
	
	return eigen.Values(nil), nil	
}

///////////////////////////////
//         HELPER            //
//         FUNCTIONS         //  
///////////////////////////////


func (m *matrix) IsSqaure() bool {
	return (m.numRows == m.numCols)
}



// brief: Calculates transpose of matrix
//
// details: Implemented for mat64.Matrix interface
// 
// returns: a transposed version of m
func (m matrix) pivotMatrix() *matrix { 

	dim := m.numRows

	// Permutation matrix to return
	P := Identity(dim)
	
	// For each column, find the row below the diagonal
	// with the highest entry in that column
	for j, row := range m.elems {
		currentMax := row[j]
		maxRow := j

		
		for i := j; i < dim; i++ {
			if m.elems[i][j] > currentMax {
				currentMax = m.elems[i][j]
				maxRow = i
			}
		}

		// If the j'th entry in maxRow is 
		// not on the diagonal, swap the
		// corresponding rows in P
		if maxRow != j {
			P.elems[j], P.elems[maxRow] = P.elems[maxRow], P.elems[j]
		}
	}

	return P
}


func permDeterminant(m *matrix) float64 {
	numberOfSwitches := 0.0
	
	for i := 0; i < m.numRows; i++ {
		if m.elems[i][i] != 1 {
			numberOfSwitches += 1
		}
	}

	return math.Pow(-1.0, (numberOfSwitches / 2.0))
}


func determinant(L, U, P matrix) (float64, error) {
	n := L.numRows
	
	// det(m) = det(L)det(P^-1)det(U)
	detL := 1.0
	detU := 1.0
	detP := permDeterminant(&P)

	// Multiply diagonals to calculate det(L) and det(U)
	for i := 0; i < n; i++ {
		detL = L.elems[i][i] * detL
	}

	for i := 0; i < n; i++ {
		detU = U.elems[i][i] * detU
	}

	return (detL * detU * detP), nil

}


func gauss(b []float64, L *matrix, U *matrix, P *matrix) ([]float64, error) {
	
	// Multiply P 
	Pb, err := P.Multiply(Matrix(b).Transpose())

	
	if err != nil {
		return nil, err
	}

	// Solve Ly = Pb
	y := make([]float64, len(b))
	for i, row := range L.elems {
		sum := 0.0
		for j := 0; j < i; j++ {
			sum += row[j]*y[j]
		}

		y[i] = (Pb.elems[i][0] - sum) / row[i]
	 
	}

	//fmt.Println(b)

	// Solve Ux = y 
	// Perform backwards substitution
	x := make([]float64, len(b))
	for i := len(b)- 1; i >= 0; i++ {
		sum := 0.0
		for j := len(b)-1; j > i; j++ {
			sum += U.elems[i][j]*x[j]
		}

		x[i] = (y[i] - sum) / U.elems[i][0]
	}

	fmt.Println(x)
	return x, nil
}

// brief: Finds the max entry in a 
//        slice of floats and it index
//
// returns: the max value in the slice and its index
func Max(s []float64) (float64, int) {
	n:= len(s)
	currentMax := s[0]
	index := 0

	for i := 1; i < n; i++ {
		if s[i] > currentMax {
			currentMax = s[i]
			index = i
		}
	} 

	return currentMax, index
}

// brief: Calculates transpose of matrix
//
// details: Implemented for mat64.Matrix interface
// 
// returns: a transposed version of m
func (m matrix) T() mat64.Matrix {
	transpose := m.Transpose()


	return *transpose

}



