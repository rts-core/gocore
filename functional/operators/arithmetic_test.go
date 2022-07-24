package operators_test

import (
	"github.com/rts-core/gocore/functional/operators"
	"github.com/stretchr/testify/assert"
	"testing"
)

var addCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, 4},
	{-3, 1, -2},
	{-3, -1, -4},
	{3, -1, 2},
	{3, 0, 3},
}

func Test_Add_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range addCases {
		a.Equal(c.expected, operators.Add(c.x, c.y))
	}
}

var mulCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, 3},
	{-3, 1, -3},
	{-3, -1, 3},
	{3, -1, -3},
	{3, 0, 0},
}

func Test_Mul_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range mulCases {
		a.Equal(c.expected, operators.Mul(c.x, c.y))
	}
}

var subCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, -2},
	{-3, 1, -4},
	{-3, -1, -2},
	{3, -1, 4},
	{3, 0, 3},
}

func Test_Sub_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range subCases {
		a.Equal(c.expected, operators.Sub(c.x, c.y))
	}
}

var divCases = []struct {
	x, y, expected int
}{
	{1, 3, 0},
	{-3, 1, -3},
	{-3, -1, 3},
	{3, -1, -3},
}

func Test_Div_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range divCases {
		a.Equal(c.expected, operators.Div(c.x, c.y))
	}
}

var xorCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, 2},
	{-3, 1, -4},
	{-3, -1, 2},
	{3, -1, -4},
	{3, 0, 3},
}

func Test_Xor_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range xorCases {
		a.Equal(c.expected, operators.Xor(c.x, c.y))
	}
}

var maxCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, 3},
	{-3, 1, 1},
	{-3, -1, -1},
	{3, -1, 3},
	{3, 0, 3},
}

func Test_Max_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range maxCases {
		a.Equal(c.expected, operators.Max(c.x, c.y))
	}
}

var minCases = []struct {
	x, y, expected int
}{
	{0, 0, 0},
	{1, 3, 1},
	{-3, 1, -3},
	{-3, -1, -3},
	{3, -1, -1},
	{3, 0, 0},
}

func Test_Min_CalledWithGiven_ReturnsExpected(t *testing.T) {
	a := assert.New(t)
	for _, c := range minCases {
		a.Equal(c.expected, operators.Min(c.x, c.y))
	}
}
