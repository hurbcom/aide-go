package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCheckStringJSONData(t *testing.T) {
	var s string
	result1 := CheckStringJSONData(s)
	assert.Nil(t, result1)

	result2 := CheckStringJSONData("")
	assert.Nil(t, result2)

	result3 := CheckStringJSONData("test")
	assert.NotNil(t, result3)
	assert.Equal(t, "test", *result3)
}

func TestShouldCheckInt64JSONData(t *testing.T) {
	var i1 int64
	result1 := CheckInt64JSONData(i1)
	assert.Nil(t, result1)

	result2 := CheckInt64JSONData(0)
	assert.Nil(t, result2)

	result3 := CheckInt64JSONData(987654)
	assert.NotNil(t, result3)
	assert.Equal(t, int64(987654), *result3)
}

func TestShouldCheckFloat64JSONData(t *testing.T) {
	var f1 float64
	result1 := CheckFloat64JSONData(f1)
	assert.Nil(t, result1)

	result2 := CheckFloat64JSONData(0)
	assert.Nil(t, result2)

	result3 := CheckFloat64JSONData(0.00)
	assert.Nil(t, result3)

	result4 := CheckFloat64JSONData(9876.54)
	assert.NotNil(t, result4)
	assert.Equal(t, float64(9876.54), *result4)
}

func TestIsArray(t *testing.T) {
	actual := IsArray([]string{"foo", "bar"})
	assert.Equal(t, true, actual)

	actual = IsArray([]int{65485, 19734})
	assert.Equal(t, true, actual)

	actual = IsArray([]int64{65485, 19734})
	assert.Equal(t, true, actual)

	actual = IsArray(nil)
	assert.Equal(t, false, actual)

	actual = IsArray(65485)
	assert.Equal(t, false, actual)

	actual = IsArray("foo")
	assert.Equal(t, false, actual)

	actual = IsArray(false)
	assert.Equal(t, false, actual)
}

func TestIsString(t *testing.T) {
	actual := IsString("foo")
	assert.Equal(t, true, actual)

	actual = IsString("")
	assert.Equal(t, true, actual)

	actual = IsString("123")
	assert.Equal(t, true, actual)

	actual = IsString("123.456")
	assert.Equal(t, true, actual)

	actual = IsString("true")
	assert.Equal(t, true, actual)

	actual = IsString([]int{1, 2})
	assert.Equal(t, false, actual)

	actual = IsString([]string{"a", "b"})
	assert.Equal(t, false, actual)

	actual = IsString(nil)
	assert.Equal(t, false, actual)

	actual = IsString(123)
	assert.Equal(t, false, actual)

	actual = IsString(123.456)
	assert.Equal(t, false, actual)

	actual = IsString(false)
	assert.Equal(t, false, actual)
}

func TestIsPointer(t *testing.T) {
	var pStr *string
	actual := IsPointer(pStr)
	assert.Equal(t, true, actual)

	var pInt *int
	actual = IsPointer(pInt)
	assert.Equal(t, true, actual)

	var pInt64 *int64
	actual = IsPointer(pInt64)
	assert.Equal(t, true, actual)

	var pFloat64 *float64
	actual = IsPointer(pFloat64)
	assert.Equal(t, true, actual)

	var pInter *interface{}
	actual = IsPointer(pInter)
	assert.Equal(t, true, actual)

	var pSlice *[]string
	actual = IsPointer(pSlice)
	assert.Equal(t, true, actual)

	var strVar string
	actual = IsPointer(strVar)
	assert.Equal(t, false, actual)

	var intVar int
	actual = IsPointer(intVar)
	assert.Equal(t, false, actual)

	var int64Var int64
	actual = IsPointer(int64Var)
	assert.Equal(t, false, actual)

	var float64Var float64
	actual = IsPointer(float64Var)
	assert.Equal(t, false, actual)

	var interVar interface{}
	actual = IsPointer(interVar)
	assert.Equal(t, false, actual)

	var sliceValue []string
	actual = IsPointer(sliceValue)
	assert.Equal(t, false, actual)
}
