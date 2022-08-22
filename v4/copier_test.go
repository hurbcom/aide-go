package aidego

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	intValue1    = 299792458
	floatValue1  = math.Pi
	stringValue1 = "O ignorante afirma, o sábio duvida, o sensato reflete"
	boolValue1   = true
	intValue2    = 97838163
	floatValue2  = math.E
	stringValue2 = "A beleza das coisas existe no espírito de quem as contempla"
	boolValue2   = false
	intValue3    = 1602176565
	floatValue3  = math.Ln10
	stringValue3 = "O que não provoca minha morte faz com que eu fique mais forte"
	boolValue3   = true
)

type TestStruct1 struct {
	IntValue      int
	StringValue   string
	FloatValue    float64
	BoolValue     bool
	StructValue   TestStruct2
	IntPointer    *int
	StringPointer *string
	FloatPointer  *float64
	BoolPointer   *bool
	StructPointer *TestStruct3
}

type TestStruct2 struct {
	A bool     `copier:"BoolValue"`
	B *float64 `copier:"FloatPointer"`
	C string   `copier:"StringValue"`
	D *string  `copier:"StringPointer"`
	E *int     `copier:"IntPointer"`
	F float64  `copier:"FloatValue"`
	G int      `copier:"IntValue"`
	H *bool    `copier:"BoolPointer"`
}

type TestStruct3 struct {
	Z int      `copier:"IntValue"`
	Y float64  `copier:"FloatValue"`
	X string   `copier:"StringValue"`
	W bool     `copier:"BoolValue"`
	V *int     `copier:"IntPointer"`
	U *float64 `copier:"FloatPointer"`
	T *string  `copier:"-"`
	S *bool    `copier:"BoolPointer"`
}
type SliceTestStructSource struct {
	IntSlice    []int
	StringSlice []string
	FloatSlice  []float64
	BoolSlice   []bool
	StructSlice []TestStruct2

	IntSlicePointer    *[]int
	StringSlicePointer *[]string
	FloatSlicePointer  *[]float64
	BoolSlicePointer   *[]bool
	StructSlicePointer *[]TestStruct3

	IntPointerSlice    []*int
	StringPointerSlice []*string
	FloatPointerSlice  []*float64
	BoolPointerSlice   []*bool
	StructPointerSlice []*TestStruct2

	IntSlicePointerPointer    *[]*int
	StringPointerSlicePointer *[]*string
	FloatPointerSlicePointer  *[]*float64
	BoolPointerSlicePointer   *[]*bool
	StructPointerSlicePointer *[]*TestStruct3
}

type SliceTestStructDestination struct {
	A []int         `copier:"IntSlice"`
	B []string      `copier:"StringSlice"`
	C []float64     `copier:"FloatSlice"`
	D []bool        `copier:"BoolSlice"`
	E []TestStruct2 `copier:"StructSlice"`

	F *[]int         `copier:"IntSlicePointer"`
	G *[]string      `copier:"StringSlicePointer"`
	H *[]float64     `copier:"FloatSlicePointer"`
	I *[]bool        `copier:"BoolSlicePointer"`
	J *[]TestStruct3 `copier:"StructSlicePointer"`

	K []*int         `copier:"IntPointerSlice"`
	L []*string      `copier:"StringPointerSlice"`
	M []*float64     `copier:"FloatPointerSlice"`
	N []*bool        `copier:"BoolPointerSlice"`
	O []*TestStruct2 `copier:"StructPointerSlice"`

	P *[]*int         `copier:"IntSlicePointerPointer"`
	Q *[]*string      `copier:"StringPointerSlicePointer"`
	R *[]*float64     `copier:"FloatPointerSlicePointer"`
	S *[]*bool        `copier:"BoolPointerSlicePointer"`
	T *[]*TestStruct3 `copier:"StructPointerSlicePointer"`
}

var (
	intPointer1    = intValue1
	floatPointer1  = floatValue1
	stringPointer1 = stringValue1
	boolPointer1   = boolValue1

	intPointer2    = intValue2
	floatPointer2  = floatValue2
	stringPointer2 = stringValue2
	boolPointer2   = boolValue2

	testStruct1 = TestStruct1{
		IntValue:      intValue2,
		StringValue:   stringValue2,
		FloatValue:    floatValue2,
		BoolValue:     boolValue2,
		StructValue:   testStruct2,
		IntPointer:    &intPointer1,
		StringPointer: &stringPointer1,
		FloatPointer:  &floatPointer1,
		BoolPointer:   &boolPointer1,
		StructPointer: &testStruct3,
	}

	testStruct2 = TestStruct2{
		A: boolValue3,
		B: &floatPointer2,
		C: stringValue3,
		D: &stringPointer2,
		E: &intPointer2,
		F: floatValue3,
		G: intValue3,
		H: &boolPointer2,
	}

	testStruct3 = TestStruct3{
		Z: intValue2,
		Y: floatValue2,
		X: stringValue2,
		W: boolValue2,
		V: &intPointer1,
		U: &floatPointer1,
		T: &stringPointer1,
		S: &boolPointer1,
	}

	testSliceSource = SliceTestStructSource{
		IntSlice:    []int{intValue1, intValue2, intValue3},
		StringSlice: []string{stringValue1, stringValue2, stringValue3},
		FloatSlice:  []float64{floatValue1, floatValue2, floatValue3},
		BoolSlice:   []bool{boolValue1, boolValue2},
		StructSlice: []TestStruct2{testStruct2},

		IntSlicePointer:    &[]int{intValue1, intValue2, intValue3},
		StringSlicePointer: &[]string{stringValue1, stringValue2, stringValue3},
		FloatSlicePointer:  &[]float64{floatValue1, floatValue2, floatValue3},
		BoolSlicePointer:   &[]bool{boolValue1, boolValue2},
		StructSlicePointer: &[]TestStruct3{testStruct3},

		IntPointerSlice:    []*int{&intPointer1, &intPointer2},
		StringPointerSlice: []*string{&stringPointer1, &stringPointer2},
		FloatPointerSlice:  []*float64{&floatPointer1, &floatPointer2},
		BoolPointerSlice:   []*bool{&boolPointer1, &boolPointer2},
		StructPointerSlice: []*TestStruct2{&testStruct2},

		IntSlicePointerPointer:    &[]*int{&intPointer1, &intPointer2},
		StringPointerSlicePointer: &[]*string{&stringPointer1, &stringPointer2},
		FloatPointerSlicePointer:  &[]*float64{&floatPointer1, &floatPointer2},
		BoolPointerSlicePointer:   &[]*bool{&boolPointer1, &boolPointer2},
		StructPointerSlicePointer: &[]*TestStruct3{&testStruct3},
	}
)

func TestCopier_Copy(t *testing.T) {

	t.Run("copier.Copy: success cases",
		func(t *testing.T) {

			var destination1 TestStruct1
			testCopier(t, false, &testStruct1, &destination1)

			assert.Equal(t, testStruct1.IntValue, destination1.IntValue)
			assert.Equal(t, testStruct1.StringValue, destination1.StringValue)
			assert.Equal(t, testStruct1.FloatValue, destination1.FloatValue)
			assert.Equal(t, testStruct1.BoolValue, destination1.BoolValue)
			assert.Equal(t, testStruct1.StructValue.A, destination1.StructValue.A)
			assert.Equal(t, testStruct1.StructValue.B, destination1.StructValue.B)
			assert.Equal(t, testStruct1.StructValue.C, destination1.StructValue.C)
			assert.Equal(t, testStruct1.StructValue.D, destination1.StructValue.D)
			assert.Equal(t, testStruct1.StructValue.E, destination1.StructValue.E)
			assert.Equal(t, testStruct1.StructValue.F, destination1.StructValue.F)
			assert.Equal(t, testStruct1.StructValue.G, destination1.StructValue.G)
			assert.Equal(t, testStruct1.StructValue.H, destination1.StructValue.H)
			assert.Equal(t, testStruct1.IntPointer, destination1.IntPointer)
			assert.Equal(t, testStruct1.StringPointer, destination1.StringPointer)
			assert.Equal(t, testStruct1.FloatPointer, destination1.FloatPointer)
			assert.Equal(t, testStruct1.BoolPointer, destination1.BoolPointer)
			assert.Equal(t, testStruct1.StructPointer, destination1.StructPointer)

			var destination2 TestStruct2
			testCopier(t, false, &testStruct1, &destination2)

			assert.Equal(t, testStruct1.IntValue, destination2.G)
			assert.Equal(t, testStruct1.StringValue, destination2.C)
			assert.Equal(t, testStruct1.FloatValue, destination2.F)
			assert.Equal(t, testStruct1.BoolValue, destination2.A)
			assert.Equal(t, testStruct1.IntPointer, destination2.E)
			assert.Equal(t, testStruct1.StringPointer, destination2.D)
			assert.Equal(t, testStruct1.FloatPointer, destination2.B)
			assert.Equal(t, testStruct1.BoolPointer, destination2.H)

			var destination3 TestStruct3
			testCopier(t, false, &testStruct1, &destination3)
			assert.Equal(t, testStruct1.IntValue, destination3.Z)
			assert.Equal(t, testStruct1.FloatValue, destination3.Y)
			assert.Equal(t, testStruct1.StringValue, destination3.X)
			assert.Equal(t, testStruct1.BoolValue, destination3.W)
			assert.Equal(t, testStruct1.IntPointer, destination3.V)
			assert.Equal(t, testStruct1.FloatPointer, destination3.U)
			assert.Empty(t, destination3.T)
			assert.Equal(t, testStruct1.BoolPointer, destination3.S)
		},
	)

	t.Run("copier.Copy: arrays and slices success cases",
		func(t *testing.T) {

			var destination1 SliceTestStructDestination
			testCopier(t, false, &testSliceSource, &destination1)

			assert.Equal(t, testSliceSource.IntSlice, destination1.A)
			assert.Equal(t, testSliceSource.StringSlice, destination1.B)
			assert.Equal(t, testSliceSource.FloatSlice, destination1.C)
			assert.Equal(t, testSliceSource.BoolSlice, destination1.D)
			assert.Equal(t, testSliceSource.StructSlice, destination1.E)
			assert.Equal(t, testSliceSource.IntSlicePointer, destination1.F)
			assert.Equal(t, testSliceSource.StringSlicePointer, destination1.G)
			assert.Equal(t, testSliceSource.FloatSlicePointer, destination1.H)
			assert.Equal(t, testSliceSource.BoolSlicePointer, destination1.I)
			assert.Equal(t, testSliceSource.StructSlicePointer, destination1.J)
			assert.Equal(t, testSliceSource.IntPointerSlice, destination1.K)
			assert.Equal(t, testSliceSource.StringPointerSlice, destination1.L)
			assert.Equal(t, testSliceSource.FloatPointerSlice, destination1.M)
			assert.Equal(t, testSliceSource.BoolPointerSlice, destination1.N)
			assert.Equal(t, testSliceSource.StructPointerSlice, destination1.O)
			assert.Equal(t, testSliceSource.IntSlicePointerPointer, destination1.P)
			assert.Equal(t, testSliceSource.StringPointerSlicePointer, destination1.Q)
			assert.Equal(t, testSliceSource.FloatPointerSlicePointer, destination1.R)
			assert.Equal(t, testSliceSource.BoolPointerSlicePointer, destination1.S)
			assert.Equal(t, testSliceSource.StructPointerSlicePointer, destination1.T)
		},
	)

	t.Run("copier.Copy: it should fail when there are no match with source and destination fields",
		func(t *testing.T) {

			var destination struct {
				A int
				B string
				C float64
				D bool
			}

			testCopier(t, true, &testStruct1, &destination)
		},
	)

	t.Run("copier.Copy: it should fail when no struct is passed as source or destination",
		func(t *testing.T) {

			intSource := intValue1
			var intDestination int
			testCopierParameterPassCombination(t, true, intSource, intDestination)

			floatSource := floatValue1
			var floatDestination float64
			testCopierParameterPassCombination(t, true, floatSource, floatDestination)

			stringSource := stringValue1
			var stringDestination string
			testCopierParameterPassCombination(t, true, stringSource, stringDestination)

			var boolDestination bool
			testCopierParameterPassCombination(t, true, boolValue1, boolDestination)
		},
	)

}

func testCopier(t *testing.T, fail bool, source, destination interface{}) {
	t.Helper()

	copier, err := NewCopier()
	assert.NoError(t, err)

	err = copier.Copy(source, destination)
	if fail {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}

	return
}

func testCopierParameterPassCombination(t *testing.T, fail bool, source, destination interface{}) {
	t.Helper()

	testCopier(t, fail, &source, &destination)
	testCopier(t, fail, source, destination)
	testCopier(t, fail, source, &destination)
	testCopier(t, fail, &source, destination)
}
