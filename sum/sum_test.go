package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{126, 1, 127},
	{-127, -1, -128},
}

var sum_tests_int32 = []struct {
	n1	 int32
	n2	 int32
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{32767, 1, 32768},
	{2147483647, 0, 2147483647},
}

var sum_tests_uint32 = []struct {
	n1	 uint32
	n2	 uint32
	expected uint32
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{32767, 1, 32768},
	{4294967293, 1, 4294967291},
}

var sum_tests_float64 = []struct {
	n1	 float64
	n2	 float64
	expected float64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{32767, 1, 32768},
	{1.0000000000000000000532, 1, 2}, // uh oh
}

var sum_tests_int64 = []struct {
	n1	 int64
	n2	 int64
	expected int64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{32767, 1, 32768},
	{9223372036854775807, 1, 9223372036854775807},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		val, err := SumInt8(v.n1, v.n2)
		if err != nil {
			t.Errorf("Sum(%d, %d) returned an error: %s", v.n1, v.n2, err)
		} else if val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		val, err := SumInt32(v.n1, v.n2)
		if err != nil {
			t.Errorf("Sum(%d, %d) returned an error: %s", v.n1, v.n2, err)
		} else if val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}

	}
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		val, err := SumUint32(v.n1, v.n2)
		if err != nil {
			t.Errorf("Sum(%d, %d) returned an error: %s", v.n1, v.n2, err)
		} else if val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}

	}
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		val, err := SumFloat64(v.n1, v.n2)
		if err != nil {
			t.Errorf("Sum(%f, %f) returned an error: %s", v.n1, v.n2, err)
		} else if val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		val, err := SumInt64(v.n1, v.n2)
		if err != nil {
			t.Errorf("Sum(%d, %d) returned an error: %s", v.n1, v.n2, err)
		} else if val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

