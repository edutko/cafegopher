package java

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_castToBool(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected bool
	}{
		{true, isOk, true},
		{false, isOk, false},
		{javaBoolean(true), isOk, true},
		{javaBoolean(false), isOk, false},

		{nil, isNotOk, false},
		{int8(1), isNotOk, false},
		{int8(-1), isNotOk, false},
		{uint8(1), isNotOk, false},
		{int16(1), isNotOk, false},
		{int16(-1), isNotOk, false},
		{uint16(1), isNotOk, false},
		{1, isNotOk, false},
		{-1, isNotOk, false},
		{uint(1), isNotOk, false},
		{int32(1), isNotOk, false},
		{int32(-1), isNotOk, false},
		{uint32(1), isNotOk, false},
		{int64(1), isNotOk, false},
		{int64(-1), isNotOk, false},
		{uint64(1), isNotOk, false},
		{float32(math.Pi), isNotOk, false},
		{math.E, isNotOk, false},
		{javaByte(1), isNotOk, false},
		{javaCharacter('a'), isNotOk, false},
		{javaShort(1), isNotOk, false},
		{javaShort(-1), isNotOk, false},
		{javaInteger(1), isNotOk, false},
		{javaInteger(-1), isNotOk, false},
		{javaLong(1), isNotOk, false},
		{javaLong(-1), isNotOk, false},
		{javaFloat(math.Pi), isNotOk, false},
		{javaDouble(math.E), isNotOk, false},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToBool(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToFloat32(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected float32
	}{
		{float32(math.Pi), isOk, math.Pi},
		{javaFloat(math.Pi), isOk, math.Pi},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(1), isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{uint8(1), isNotOk, 0},
		{int16(1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{uint16(1), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaByte(1), isNotOk, 0},
		{javaShort(1), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToFloat32(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToFloat64(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected float64
	}{
		{float32(math.Pi), isOk, float64(float32(math.Pi))},
		{math.E, isOk, math.E},
		{javaFloat(math.Pi), isOk, float64(float32(math.Pi))},
		{javaDouble(math.E), isOk, math.E},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(1), isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{uint8(1), isNotOk, 0},
		{int16(1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{uint16(1), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaByte(1), isNotOk, 0},
		{javaShort(1), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToFloat64(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToInt(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected int
	}{
		{int8(127), isOk, 127},
		{int8(-128), isOk, -128},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{int16(-32768), isOk, -32768},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int(math.MinInt16) - 1, isOk, math.MinInt16 - 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int32(math.MinInt16 - 1), isOk, math.MinInt16 - 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaShort(-32768), isOk, -32768},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},
		{javaInteger(math.MinInt16 - 1), isOk, math.MinInt16 - 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToInt(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToInt8(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected int8
	}{
		{int8(127), isOk, 127},
		{int8(-128), isOk, -128},
		{javaByte(127), isOk, 127},
		{javaByte(-128), isOk, -128},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint8(1), isNotOk, 0},
		{int16(1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{uint16(1), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(1), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToInt8(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToInt16(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected int16
	}{
		{int8(127), isOk, 127},
		{int8(-128), isOk, -128},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{int16(-32768), isOk, -32768},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaShort(-32768), isOk, -32768},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint16(65535), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToInt16(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToInt32(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected int32
	}{
		{int8(127), isOk, 127},
		{int8(-128), isOk, -128},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{int16(-32768), isOk, -32768},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int(math.MinInt16) - 1, isOk, math.MinInt16 - 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int32(math.MinInt16 - 1), isOk, math.MinInt16 - 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaShort(-32768), isOk, -32768},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},
		{javaInteger(math.MinInt16 - 1), isOk, math.MinInt16 - 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToInt32(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToInt64(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected int64
	}{
		{int8(127), isOk, 127},
		{int8(-128), isOk, -128},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{int16(-32768), isOk, -32768},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int(math.MinInt16) - 1, isOk, math.MinInt16 - 1},
		{uint(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{int32(math.MinInt16 - 1), isOk, math.MinInt16 - 1},
		{uint32(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int64(math.MaxUint32) + 1, isOk, math.MaxUint32 + 1},
		{int64(math.MinInt32 - 1), isOk, math.MinInt32 - 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaShort(-32768), isOk, -32768},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},
		{javaInteger(math.MinInt16 - 1), isOk, math.MinInt16 - 1},
		{javaLong(math.MaxUint32 + 1), isOk, math.MaxUint32 + 1},
		{javaLong(math.MinInt32 - 1), isOk, math.MinInt32 - 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToInt64(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToUint(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected uint
	}{
		{int8(127), isOk, 127},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint32(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{-1, isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(math.MaxUint32 + 1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToUint(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToUint8(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected uint8
	}{
		{int8(127), isOk, 127},
		{int8(-1), isOk, 255}, // uint8 is special because consumers will expect Java bytes to map to Go bytes
		{uint8(255), isOk, 255},
		{javaByte(127), isOk, 127},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{uint16(1), isNotOk, 0},
		{int16(1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(1), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToUint8(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToUint16(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected uint16
	}{
		{int8(127), isOk, 127},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{uint16(65535), isOk, 65535},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{1, isNotOk, 0},
		{-1, isNotOk, 0},
		{uint(1), isNotOk, 0},
		{int32(1), isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{uint32(1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToUint16(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToUint32(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected uint32
	}{
		{int8(127), isOk, 127},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint32(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{-1, isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{int64(1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{uint64(1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToUint32(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func Test_castToUint64(t *testing.T) {
	testCases := []struct {
		value    any
		ok       okay
		expected uint64
	}{
		{int8(127), isOk, 127},
		{uint8(255), isOk, 255},
		{int16(32767), isOk, 32767},
		{uint16(65535), isOk, 65535},
		{int(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int32(math.MaxUint16) + 1, isOk, math.MaxUint16 + 1},
		{uint32(math.MaxInt32) + 1, isOk, math.MaxInt32 + 1},
		{int64(math.MaxUint32) + 1, isOk, math.MaxUint32 + 1},
		{uint64(math.MaxInt64) + 1, isOk, math.MaxInt64 + 1},
		{javaByte(127), isOk, 127},
		{javaShort(32767), isOk, 32767},
		{javaInteger(math.MaxUint16 + 1), isOk, math.MaxUint16 + 1},
		{javaLong(math.MaxUint32 + 1), isOk, math.MaxUint32 + 1},

		{nil, isNotOk, 0},
		{true, isNotOk, 0},
		{false, isNotOk, 0},
		{int8(-1), isNotOk, 0},
		{int16(-1), isNotOk, 0},
		{-1, isNotOk, 0},
		{int32(-1), isNotOk, 0},
		{int64(-1), isNotOk, 0},
		{float32(math.Pi), isNotOk, 0},
		{math.E, isNotOk, 0},
		{javaBoolean(true), isNotOk, 0},
		{javaBoolean(false), isNotOk, 0},
		{javaShort(-1), isNotOk, 0},
		{javaInteger(-1), isNotOk, 0},
		{javaLong(-1), isNotOk, 0},
		{javaFloat(math.Pi), isNotOk, 0},
		{javaDouble(math.E), isNotOk, 0},
	}

	for _, tc := range testCases {
		t.Run(stringify(tc.value), func(t *testing.T) {
			actual, ok := castToUint64(tc.value)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, bool(tc.ok), ok)
		})
	}
}

func stringify(val any) string {
	if v, ok := val.(Object); ok {
		return fmt.Sprintf("%s(%v)", v.ClassDesc.ClassName, v.ClassData[v.ClassDesc.ClassName]["value"])
	}
	return fmt.Sprintf("%T(%v)", val, val)
}

type okay bool

const isOk = true
const isNotOk = false
