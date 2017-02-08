package matrix

import (
	"testing";
	"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)


//************************
// Constructor Test Suite
//************************

type ConstructorsTestSuite struct {
	suite.Suite
	EmptyMatrix1 Matrix
	EmptyMatrix2 Matrix
	MatrixFromSlices1 Matrix
	MatrixFromSlices2 Matrix 
}

//************************
// END Constructor Test Suite
//************************

type GetEntryTestSuite struct {}

//********************************
// Addition of Matrices Test Suite
//********************************


// This data structure is a test suite 
// that tests different cases for adding Matrices. 
type AdditionTestSuite struct {
	suite.Suite
	DiffDimMatrix1 Matrix
	DiffDimMatrix2 Matrix
	AddToItselfMatrix Matrix
	ResultAddToItself Matrix
	SquareMatrix1 Matrix
	SquareMatrix2 Matrix

}

//********************************
// END Addition of Matrices Test Suite
//********************************

type EqualsTestSuite struct{}

// Initializes all matrices to be tested 
// in the addition test sweet. 
func (suite *AdditionTestSuite) SetupTest() {
	DiffDimMatrix1 := BlankMatrix{10, 2}
	DiffDimMatrix1 := BlankMatrix{1, 5}

	AddToItselfMatrix := Matrix{[]float64{1, 0, 0}, []float64{0, 1, 0}, []float64{0, 0, 1}}
	ResultAddToItself := Matrix{[]float64{2, 0, 0}, []float64{0, 2, 0}, []float64{0, 0, 2}}

	SquareMatrix1 := Matrix{[]float64{1, 6}, []float64{5, -7}}
	SquareMatrix1 := Matrix{[]float64{1.05, -10}, []float64{-103, 4}}
}


// Different Dimension Addition Test
// Adding two matrices of different dimensions should 
// raise and Error
func (suite *AdditionTestSuite) TestDifferentDimAddition() {
	sum, err := suite.DiffDimMatrix1.Add(suite.DiffDimMatrix2)
	assert.NotEqual(suite.T(), nil, err) 
}


// The addition method uses pointers as receivers 
// to avoid the inefficiency of copying the Matrix arguments
// This tests to make sure that we don't modify the original
// arguments when we add a matrix to itself
func (suite *AdditionTestSuite) TestAddToItself() {
	// Copy the matrix before we add it to itself
	// Test this copy against the one that added to itself
	// after adding
	copyMatrix := suite.AddToItselfMatrix

	sum        := suite.AddToItselfMatrix.Add(suite.AddToItselfMatrix)

	assert.Equal(suite.T(), copyMatrix, suite.AddToItselfMatrix)
	assert.Equal(suite.T(), suite.ResultAddToItself, sum)
}



//********************************
// Multiplication of Matrices Test Suite
//********************************

type MultiplicationTestSuite struct{

}



type GetRowTestSuite struct {}

type GetColTestSuite struct {}

type DeterminantTestSuite struct {}

type InverseTestSuite struct {}

type GetEigenValuesTestSuite struct{}










