package sum

import "errors"
import "fmt"

func SumInt8(a, b int8) (int8, error) {
	result := a + b
	if a > 0 && b > 0 { // Dersom a og b er positive
		if result < 0 { // kan ikke resultatet være negativt.
			errorstring := fmt.Sprintf("%d + %d results in an overflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	} else if a < 0 && b < 0 { // Dersom a og b er negative
		if result >= 0 { // kan ikke resultatet være positivt eller 0.
			errorstring := fmt.Sprintf("%d + %d results in an underflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	}
	return result, nil
}

func SumInt32(a, b int32) (int32, error) {
	result := a + b
	if a > 0 && b > 0 {
		if result < 0 {
			errorstring := fmt.Sprintf("%d + %d results in an overflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	} else if a < 0 && b < 0 {
		if result >= 0 {
			errorstring := fmt.Sprintf("%d + %d results in an underflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	}
	return result, nil
}

func SumInt64(a, b int64) (int64, error) {
	result := a + b
	if a > 0 && b > 0 {
		if result < 0 {
			errorstring := fmt.Sprintf("%d + %d results in an overflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	} else if a < 0 && b < 0 {
		if result >= 0 {
			errorstring := fmt.Sprintf("%d + %d results in an underflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	}
	return result, nil
}

func SumUint32(a, b uint32) (uint32, error) {
	result := a + b
	check := a * b // Enkel måte å sjekke om uint32 har overflowet
	if check/b != a {
			errorstring := fmt.Sprintf("%d + %d results in an overflow", a, b)
			err := errors.New(errorstring)
			return result, err
	}
	return result, nil

}

func SumFloat64(a, b float64) (float64, error) {
	result := a + b
	if a > 0 && b > 0 { // Dette fungerer ikke for floats.. Usikker på hvordan man kan sjekke
		if result < 0 { // presisjonsfeil og lignende.
			errorstring := fmt.Sprintf("%d + %d results in an overflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	} else if a < 0 && b < 0 {
		if result >= 0 {
			errorstring := fmt.Sprintf("%d + %d results in an underflow", a, b)
			err := errors.New(errorstring)
			return result, err
		}
	}
	return result, nil

}
