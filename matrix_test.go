package golinal

import (
    "github.com/stretchr/testify/suite";
    "testing"
)

//****************
// Global Matrices
//****************

var NonsquareMatrix = NewMatrix([]float64{1}, []float64{-7})
var NonsquareMatrix2 = NewMatrix([]float64{3, 4})

var ThreeIdentity = NewMatrix([]float64{1, 0, 0}, []float64{0, 1, 0}, []float64{0, 0, 1})

var RandFourMatrix = NewMatrix (
        []float64{0.223548, 7.51484, 7.94393, 7.95676}, 
        []float64{9.44692, -2.05097, -3.59421, -7.9301}, 
        []float64{-5.90911, -9.56427, -6.67171, -8.09466}, 
        []float64{-9.43214, 6.42982, 7.37722, 3.8219})



var RandMatrix = NewMatrix(
        []float64{2.99875, -0.722266, -0.237451, -1.11405, -2.127, 8.88714, -1.65288, -5.27189, -5.92509, -6.02403}, 
        []float64{3.26164, 5.86218, 2.81815, -2.06958, 0.366388, -0.271817, -3.51731, 3.22294, -4.71693, -8.95407}, 
        []float64{6.53936, 0.653704, 5.51595, 8.75519, 4.50956, -2.18589, 1.44052, -7.2319, -6.35739, 9.8645}, 
        []float64{-7.58145, 8.1194, -7.58264, 9.88342, 1.48929, -3.66263, 1.87859, 8.37529, 0.772604, -6.30053}, 
        []float64{-9.39711, -6.49522, 1.94943, 9.03285, 3.40668, -7.61823, -7.2272, 4.2087, -8.91554, 4.15006}, 
        []float64{-2.80846, 5.17557, 9.52006, -9.47033, 5.67815, 6.8402, 0.818774, -6.92541, 8.22727, 5.83063}, 
        []float64{0.627448, -3.28157, 5.50732, 9.96776, -7.11846, -7.3921, -6.67718, -0.621603, -7.81631, -2.35664}, 
        []float64{-3.93259, -2.36996, 4.96011, 7.4244, -6.90362, -1.08256, 5.21275, -8.66966, -8.88118, -3.51107}, 
        []float64{8.40158, -6.03886,  7.62759, -2.43493, 0.36834, 9.46535, -4.76645, 7.49544, -7.4696, -7.31169}, 
        []float64{-5.5276, -5.28959, -2.90794, -0.984676, 7.66101, -4.13466, -6.00181, 1.37331, -9.62865, 1.92439})


// Nil pointer for Matrix
var nilMatrixP *Matrix


//************************
// Constructor Test Suite
//************************

type ConstructorsTestSuite struct {
	suite.Suite

	EmptyMatrix, 
	MatrixFromSlices *Matrix 
}


func (suite *ConstructorsTestSuite) SetupTest() {
    suite.EmptyMatrix = BlankMatrix(2, 3)

    suite.MatrixFromSlices = NewMatrix([]float64{1, 0, 0}, []float64{0, 1, 0}, []float64{0, 0, 1})
}

func (suite *ConstructorsTestSuite) TestConstructors() {
    suite.Equal(suite.EmptyMatrix.NumRows(), 2, "They should be equal")
    suite.Equal(suite.EmptyMatrix.NumCols(), 3, "They Should be equal")

    suite.Equal(suite.MatrixFromSlices.NumRows(), 3, "They should be equal")
    suite.Equal(suite.MatrixFromSlices.NumCols(), 3, "They Should be equal")
}


//********************************
// Addition of Matrices Test Suite
//********************************


// This data structure is a test suite 
// that tests different cases for adding Matrices. 
type AdditionTestSuite struct {
	suite.Suite

	DiffDimMatrix1,
	DiffDimMatrix2 *Matrix

	AddToItselfMatrix,
    CopyOfItselfMatrix,
	ResultAddToItself *Matrix

	SquareMatrix1,
	SquareMatrix2,
    ResultMatrix *Matrix

}

// Initializes all matrices to be tested 
// in the addition test sweet. 
func (suite *AdditionTestSuite) SetupTest() {
    suite.DiffDimMatrix1 = BlankMatrix(10, 2)
    suite.DiffDimMatrix2 = BlankMatrix(1, 5)

    
    suite.AddToItselfMatrix = NewMatrix([]float64{1, 0, 0}, []float64{0, 1, 0}, []float64{0, 0, 1})
    suite.CopyOfItselfMatrix = ThreeIdentity
    suite.ResultAddToItself = NewMatrix([]float64{2, 0, 0}, []float64{0, 2, 0}, []float64{0, 0, 2})

    
    suite.SquareMatrix1 = NewMatrix([]float64{1, 6}, []float64{5, -7})
    suite.SquareMatrix2 = NewMatrix([]float64{1.05, -10}, []float64{-103, 4})
    suite.ResultMatrix  = NewMatrix([]float64{2.05, -4}, []float64{-98, -3})
}


// Different Dimension Addition Test
// Adding two matrices of different dimensions should 
// raise and Error
func (suite *AdditionTestSuite) TestDifferentDimAddition() {
    err := suite.DiffDimMatrix1.Add(suite.DiffDimMatrix2)
    suite.NotEqual(nil, err)
}


// The addition method uses pointers as receivers 
// to avoid the inefficiency of copying the Matrix arguments
// This tests to make sure that we don't modify the original
// arguments when we add a Matrix to itself
func (suite *AdditionTestSuite) TestAddToItself() {
    // Copy the Matrix before we add it to itself
    // Test this copy against the one that added to itself
    // after adding
    

    err := suite.AddToItselfMatrix.Add(suite.AddToItselfMatrix)

    suite.Equal(err, nil, "They should be equal")
    suite.NotEqual(suite.CopyOfItselfMatrix, suite.AddToItselfMatrix, "They should not be equal")
    suite.Equal(suite.ResultAddToItself, suite.AddToItselfMatrix, "They should be equal")
}

func (suite *AdditionTestSuite) TestSimpleAdd() {
    err := suite.SquareMatrix1.Add(suite.SquareMatrix2)

    suite.Equal(err, nil, "They should be equal")
    suite.Equal(suite.ResultMatrix, suite.SquareMatrix1, "They should be equal")

}


//**************************************
// Multiplication of Matrices Test Suite
//**************************************

type MultiplicationTestSuite struct {
    suite.Suite
    MismatchRowCol1,
    MismatchRowCol2,
    SquaredRandMatrix *Matrix
    
}


func (suite *MultiplicationTestSuite) SetupTest() {
    suite.MismatchRowCol1 = NewMatrix([]float64{1, 6, 3}, []float64{5, -7, 3})
    suite.MismatchRowCol2 = NonsquareMatrix

    suite.SquaredRandMatrix = NewMatrix(
        []float64{11.7719, 129.774, 21.9211, -153.57, 33.6721, 95.1895, 65.75, -102.2, 240.369, 90.2049}, 
        []float64{55.3243, 93.0136, 31.7845, 3.5077, -63.0053, 39.0486, 87.8007, -105.999, 48.3232, -17.1332}, 
        []float64{-123.397, 31.8222, -152.806, 154.054, 140.097, -139.273, -99.3154, 62.8844, -122.848, 47.4363}, 
        []float64{-114.917, 102.382, -47.9481, 156.305, -135.274, -115.198, 43.1011, 158.743, -15.1189, -230.882}, 
        []float64{-234.646, 27.4006, -238.697, 210.004, 59.1005, -249.836, 115.081, 64.1364, 33.6, 122.233}, 
        []float64{134.585, -106.579, 231.721, -100.351, 184.561, 46.2524, -168.852, -10.9457, -132.449, 161.053}, 
        []float64{-15.112, 155.872, -231.474, 108.712, 1.80904, -52.6207, 175.303, -1.27922, 125.886, 15.0193}, 
        []float64{6.69661, 167.354, -130.272, 86.6587, 3.83158, -125.846, 68.3412, 18.3789, 197.724, 88.4347}, 
        []float64{-11.0144, 72.0043, 107.567, -6.93323, -10.8936, 117.14, 169.24, -331.361, 99.8063, 176.223}, 
        []float64{-206.458, -43.6637, -151.877, 62.1927, 42.2995, -180.585, 44.5284, 8.04939, 61.2176, 149.294})
}

func (suite *MultiplicationTestSuite) TestMultiplication() {
    mult1, err1 := suite.MismatchRowCol1.Multiply(suite.MismatchRowCol2)
    mult2, err2 := suite.MismatchRowCol2.Multiply(suite.MismatchRowCol1)
    mult3, err3 := ThreeIdentity.Multiply(ThreeIdentity)
    mult4, err4 := RandMatrix.Multiply(Identity(10))
    mult5, err5 := RandMatrix.Multiply(RandMatrix)

    suite.Equal(mult1, nilMatrixP, "They should be equal")
    suite.NotEqual(err1, nil, "They should not be equal")

    suite.Equal(mult2, nilMatrixP, "They should be equal")
    suite.NotEqual(err2, nil, "They should be equal")

    suite.Equal(mult3, ThreeIdentity, "They should be equal")
    suite.Equal(err3, nil, "They should be equal")

    suite.Equal(RandMatrix, mult4, "They should be equal")
    suite.Equal(err4, nil, "They should be equal")

    suite.Equal(suite.SquaredRandMatrix, mult5, "They should be equal")
    suite.Equal(err5, nil, "They should be equal")
}


//*****************************
// LUP Decomposition Test Suite
//*****************************

type LUPDecompTestSuite struct {
    suite.Suite

    ThreeMatrix,
    FourMatrix *Matrix

    LThree,
    UThree,
    PThree *Matrix

    LFour,
    UFour,
    PFour *Matrix
}

func (suite *LUPDecompTestSuite) SetupTest() {
    suite.ThreeMatrix = NewMatrix(
        []float64{1,  3,  5},
        []float64{2,  4,  7},
        []float64{1,  1,  0})

    suite.LThree = NewMatrix(
        []float64{1.00000,  0.00000,  0.00000},
        []float64{0.50000,  1.00000,  0.00000},
        []float64{0.50000, -1.00000,  1.00000})

    suite.UThree = NewMatrix(
        []float64{2.00000,  4.00000,  7.00000},
        []float64{0.00000,  1.00000,  1.50000},
        []float64{0.00000,  0.00000, -2.00000})

    suite.PThree = NewMatrix(
        []float64{0,  1,  0},
        []float64{1,  0,  0},
        []float64{0,  0,  1})

    suite.FourMatrix = NewMatrix(
        []float64{11,  9, 24,  2},
        []float64{ 1,  5,  2,  6},
        []float64{ 3, 17, 18,  1},
        []float64{ 2,  5,  7,  1})

    suite.LFour = NewMatrix(
        []float64{1.00000,  0.00000, 0.00000,  0.00000},
        []float64{0.27273,  1.00000, 0.00000,  0.00000},
        []float64{0.09091,  0.28750, 1.00000,  0.00000},
        []float64{0.18182,  0.23125, 0.00360,  1.00000})

    suite.UFour = NewMatrix(
        []float64{1.00000,  0.00000, 0.00000,  0.00000},
        []float64{0.27273,  1.00000, 0.00000,  0.00000},
        []float64{0.09091,  0.28750, 1.00000,  0.00000},
        []float64{0.18182,  0.23125, 0.00360,  1.00000})
    
    suite.PFour = NewMatrix(
        []float64{1, 0, 0, 0},
        []float64{0, 0, 1, 0},
        []float64{0, 1, 0, 0},
        []float64{0, 0, 0, 1})
}

func (suite *LUPDecompTestSuite) TestLUP() {
    L1, U1, P1, err1 := suite.ThreeMatrix.LUP()
    L2, U2, P2, err2 := suite.FourMatrix.LUP()

    suite.Equal(suite.LThree, L1, "They should be equal")
    suite.Equal(suite.UThree, U1, "They should be equal")
    suite.Equal(suite.PThree, P1, "They should be equal")
    suite.Equal(nil, err1, "There should be no error")

    suite.Equal(suite.LFour, L2, "They should be equal")
    suite.Equal(suite.UFour, U2, "They should be equal")
    suite.Equal(suite.PFour, P2, "They should be equal")
    suite.Equal(nil, err2, "There should be no error")
   
}


//*******************************
// Inverse of Matrices Test Suite
//*******************************
type InverseTestSuite struct {
    suite.Suite

    ZeroDeterminantMatrix,
    TenIdentity *Matrix

    RandMatrixInverse *Matrix
}

func (suite *InverseTestSuite) SetupTest() {


    suite.ZeroDeterminantMatrix = NewMatrix([]float64{2, -2}, []float64{-2, 2})

    suite.TenIdentity = Identity(10)

    suite.RandMatrixInverse = NewMatrix(
        []float64{-0.0460854, -0.0705092, 0.0326931, 0.045097, -0.164044, 0.0123061, 0.104152, -0.0198876, 0.0493236, 0.102875}, 
        []float64{0.141227, 0.325278, 0.0262483, -0.179546, 0.419574, -0.121419, -0.306189, -0.0357916, -0.162755, -0.362409}, 
        []float64{-0.0715951, -0.0268494, 0.0000621608, 0.029943, -0.0600711, 0.0617097, 0.0618292, 0.0345928, 0.0558978, 0.0424593}, 
        []float64{-0.0430982, -0.202563, 0.0235125, 0.164038, -0.235156, 0.0869331, 0.207031, -0.00181176, 0.0909108, 0.178495}, 
        []float64{-0.170888, -0.341571, 0.0211468, 0.281072, -0.546889, 0.193541, 0.36861, 0.0233994, 0.187385, 0.486657}, 
        []float64{0.133272, 0.0910483, 0.00621968, -0.0741133, 0.234174, -0.0504633, -0.16888, -0.0265523, -0.0423938, -0.202145}, 
        []float64{-0.0699641, 0.0156878, 0.0117219, -0.00723369, -0.00487834, -0.029769, -0.0838612, 0.0756502, 0.0214575, -0.0122182}, 
        []float64{0.0507206, 0.222235, -0.00168086, -0.150038, 0.356268, -0.126108, -0.272279, -0.0189743, -0.0668899, -0.298225}, 
        []float64{-0.133813, -0.329467, -0.0227785, 0.228099, -0.457624, 0.170425, 0.362524,-0.00890932, 0.141323, 0.346878}, 
        []float64{0.168294, 0.34461, 0.031145, -0.287184, 0.592122, -0.189706, -0.426741, -0.0494601, -0.190007, -0.486913})
}

func (suite *InverseTestSuite) TestInverse() {
    inverse1, err1 := NonsquareMatrix.Inverse()
    inverse2, err2 := suite.ZeroDeterminantMatrix.Inverse()
    inverse3, err3 := suite.TenIdentity.Inverse()
    inverse4, err4 := RandMatrix.Inverse()


    suite.Equal(nilMatrixP, inverse1, "Non sqaure Matrix should be nil")
    suite.NotEqual(nil, err1, "There should be an error")

    suite.Equal(nilMatrixP, inverse2, "Zero determinant should be nil")
    suite.NotEqual(nil, err2, "There should be an error")

    suite.Equal(suite.TenIdentity, inverse3, "Identity inverse is itself")
    suite.Equal(nil, err3, "There should be no error")

    suite.Equal(suite.RandMatrixInverse, inverse4, "Identity inverse is itself")
    suite.Equal(nil, err4, "There should be no error")
}




//****************************************************
// Eigenvalues and Determinants of Matrices Test Suite
//****************************************************
type EigValDeterminantTestSuite struct{
    suite.Suite

    Uppertriangular1,
    Uppertriangular2, 
    RandMatrix *Matrix

    IdentityEigenVals,
    Upper1EigenVals,
    Upper2EigenVals,
    RandEigenVals []complex128

}


func (suite *EigValDeterminantTestSuite) SetupTest() {

    suite.IdentityEigenVals = []complex128{1+0i, 1+0i, 1+0i}

    suite.Uppertriangular1 = NewMatrix(
        []float64{5, 10, 9, 3, 4}, 
        []float64{0, 4, -6, 7.234, -3}, 
        []float64{0, 0, 3, 13098.38, 239}, [
        ]float64{0, 0, 0, 2, -70}, 
        []float64{0, 0, 0, 0, 1})

    suite.Upper1EigenVals  = []complex128{5+0i, 4+0i, 3+0i, 2+0i, 1+0i}

    suite.Uppertriangular2 = NewMatrix(
        []float64{1, 10, 9, 3, 4}, 
        []float64{0, 3, -6, 7.234, -3}, 
        []float64{0, 0, 4, 13098.38, 239}, [
        ]float64{0, 0, 0, 5, -70}, 
        []float64{0, 0, 0, 0, 2})

    suite.Upper2EigenVals = []complex128{1+0i, 3+0i, 4+0i, 5+0i, 2+0i}

    suite.RandEigenVals = []complex128{18.5377+11.1238i, 18.5377-11.1238i, -13.4118+9.58048i, -13.4118-9.58048i, -13.5186+0i, 0.557657+12.1091i, 0.557657-12.1091i, 9.66163+0i, 7.42975+0i, -1.32488+0i}
}


func (suite *EigValDeterminantTestSuite) TestDeterminant() {
    det1, err1 := NonsquareMatrix.Determinant()

    det2, err2 := ThreeIdentity.Determinant()
    det3, err3 := suite.Uppertriangular1.Determinant()
    det4, err4 := suite.Uppertriangular2.Determinant()
    det5, err5 := RandMatrix.Determinant()
    det6, err6 := RandFourMatrix.Determinant()

    suite.Equal(0.0, det1, "There should be no determinant")
    suite.NotEqual(err1, nil, "There should be an error")

    suite.Equal(det2, 1.0, "They should be equal")
    suite.Equal(err2, nil, "There should be no error")

    suite.Equal(det3, 120.0, "They should be equal")
    suite.Equal(err3, nil, "There should be no error")
    
    suite.Equal(det4, 120.0,"They should be equal")
    suite.Equal(err4, nil, "There should be no error")

    suite.Equal(det5, 2.39872e+10, "They should be equal")
    suite.Equal(err5, nil, "There should be no error")

    suite.Equal(0.0, det6, "They should be equal")
    suite.Equal(nil, err6, "There should be no error")

}

func (suite *EigValDeterminantTestSuite) TestEigenvalues() {
    eval1, err1 := NonsquareMatrix.Eigenvalues()
    eval2, err2 := ThreeIdentity.Eigenvalues()
    eval3, err3 := suite.Uppertriangular1.Eigenvalues()
    eval4, err4 := suite.Uppertriangular2.Eigenvalues()
    eval5, err5 := RandMatrix.Eigenvalues()


    suite.Equal(eval1, nil, "There should be no eigenvalues")
    suite.NotEqual(err1, nil, "There should be an error")

    suite.Equal(eval2, suite.IdentityEigenVals, "They should be equal")
    suite.Equal(err2, nil, "There should be no error")

    suite.Equal(eval3, suite.Upper1EigenVals, "They should be equal")
    suite.Equal(err3, nil, "There should be no error")
    
    suite.Equal(eval4, suite.Upper2EigenVals,"They should be equal")
    suite.Equal(err4, nil, "There should be no error")

    suite.Equal(eval5, suite.RandEigenVals, "They should be equal")
    suite.Equal(err5, nil, "There should be no error")

}


// brief: Runs all test suites when "go test" is run
//
// 
//
func TestAll(t *testing.T) {
	suite.Run(t, new(ConstructorsTestSuite))
    suite.Run(t, new(AdditionTestSuite))
    suite.Run(t, new(MultiplicationTestSuite))
    suite.Run(t, new(LUPDecompTestSuite))
    //suite.Run(t, new(InverseTestSuite))
    suite.Run(t, new(EigValDeterminantTestSuite))
    
}







