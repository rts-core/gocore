package functional_test

import (
	"github.com/rts-core/gocore/functional"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SliceMap_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0, 1, 2, 3, 4, 5}
	expected := []int{0, 2, 4, 6, 8, 10}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceMap_CalledWithSingleItem_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{5}
	expected := []int{10}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceMap_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []int
	expected := make([]int, 0)
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceMap_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]int, 0)
	expected := make([]int, 0)
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFlatMap_CalledWithFullEqualSizedArrays_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := [][]int{{0, 1, 2, 3}, {0, 1, 2, 3}}
	expected := []int{0, 2, 4, 6, 0, 2, 4, 6}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceFlatMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFlatMap_CalledWithFullNonEqualSizedArrays_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := [][]int{{0, 1, 2, 3}, {0, 1, 2}}
	expected := []int{0, 2, 4, 6, 0, 2, 4}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceFlatMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFlatMap_CalledWithSingleEqualSizedArrays_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := [][]int{{1}, {2}}
	expected := []int{2, 4}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceFlatMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFlatMap_CalledWithNonFullNonEqualSizedArrays_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := [][]int{{}, {2}}
	expected := []int{4}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceFlatMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFlatMap_CalledWithNilNonEqualSizedArrays_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := [][]int{nil, {2}}
	expected := []int{4}
	functor := func(x int) int { return x * 2 }

	actual := functional.SliceFlatMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFilter_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0, 1, 2, 3, 4, 5}
	expected := []int{0, 2, 4}
	functor := func(x int) bool { return x%2 == 0 }

	actual := functional.SliceFilter(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFilter_CalledWithSingleArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0}
	expected := []int{0}
	functor := func(x int) bool { return x%2 == 0 }

	actual := functional.SliceFilter(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFilter_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]int, 0)
	expected := make([]int, 0)
	functor := func(x int) bool { return x%2 == 0 }

	actual := functional.SliceFilter(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFilter_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []int
	expected := make([]int, 0)
	functor := func(x int) bool { return x%2 == 0 }

	actual := functional.SliceFilter(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFoldl_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []string{"1", "2", "3", "4", "5"}
	expected := "012345"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldl(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceFoldl_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]string, 0)
	expected := "0"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldl(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceFoldl_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []string
	expected := "0"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldl(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceFoldr_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []string{"1", "2", "3", "4", "5"}
	expected := "054321"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldr(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceFoldr_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]string, 0)
	expected := "0"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldr(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceFoldr_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []string
	expected := "0"
	functor := func(a, b string) string { return a + b }

	actual := functional.SliceFoldr(functor, initial, "0")

	a.Equal(expected, actual)
}

func Test_SliceReduce_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0, 1, 2, 3, 4, 5}
	expected := 15
	functor := func(x, y int) int { return x + y }

	actual := functional.SliceReduce(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceReduce_CalledWithLenTwoArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0, 1}
	expected := 1
	functor := func(x, y int) int { return x + y }

	actual := functional.SliceReduce(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceReduce_CalledWithLenOneArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{42}
	expected := 42
	functor := func(x, y int) int { return x + y }

	actual := functional.SliceReduce(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceReduce_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]int, 0)
	expected := 0 // default value
	functor := func(x, y int) int { return x + y }

	actual := functional.SliceReduce(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceReduce_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []int
	expected := 0 // default value
	functor := func(x, y int) int { return x + y }

	actual := functional.SliceReduce(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceToMap_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{0, 1, 2, 3, 4, 5}
	expected := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	functor := func(x int) int { return x }

	actual := functional.SliceToMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceToMap_CalledWithSingleItem_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{5}
	expected := map[int]int{5: 5}
	functor := func(x int) int { return x }

	actual := functional.SliceToMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceToMap_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []int
	expected := make(map[int]int)
	functor := func(x int) int { return x }

	actual := functional.SliceToMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceToMap_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]int, 0)
	expected := make(map[int]int)
	functor := func(x int) int { return x }

	actual := functional.SliceToMap(functor, initial)

	a.Equal(expected, actual)
}

func Test_SliceFirst_CalledWithFullArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{1, 2, 3, 4, 5}
	expected := 2
	functor := func(x int) bool { return x%2 == 0 }

	actual, found := functional.SliceFirst(functor, initial)

	a.Equal(expected, actual)
	a.Equal(true, found)
}

func Test_SliceFirst_CalledWithSingleArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := []int{1}
	functor := func(x int) bool { return x%2 == 0 }

	_, found := functional.SliceFirst(functor, initial)

	a.Equal(false, found)
}

func Test_SliceFirst_CalledWithEmptyArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	initial := make([]int, 0)
	functor := func(x int) bool { return x%2 == 0 }

	_, found := functional.SliceFirst(functor, initial)

	a.Equal(false, found)
}

func Test_SliceFirst_CalledWithNilArray_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	var initial []int
	functor := func(x int) bool { return x%2 == 0 }

	_, found := functional.SliceFirst(functor, initial)

	a.Equal(false, found)
}
