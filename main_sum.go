package main

import "./sum"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
	args := os.Args
	if len(args) == 3 { // Om argumentlengden er 3, vil programmet bruke int64 og float64.
		format := "int"
		var a_int64, b_int64 int64
		var a, b float64

		a_int64, a_err := strconv.ParseInt(args[1], 10, 64) // Sjekk først om argumentene er hele tall
		b_int64, b_err := strconv.ParseInt(args[2], 10, 64) // vi bruker int64 til dette
		if a_err != nil || b_err != nil {
			a, a_err = strconv.ParseFloat(args[1], 64) // Dersom en av tallene ikke er hele tall
			b, b_err = strconv.ParseFloat(args[2], 64) // prøver vi å sjekke om argumentene er floats
			format = "float"
			if a_err != nil || b_err != nil {
				// Dersom det ikke er float eller int, kan vi ikke bruke det. Vi vet at en av
				// errorene ikke er nil, og bruker TestError for å printe den.
				TestError(a_err, b_err)
			}
		}
		if format == "float" {
			result, err := sum.SumFloat64(a, b)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		} else if format == "int" {
			result, err := sum.SumInt64(a_int64, b_int64)
			// Alle SumInt-funksjonene i pakken sum returnerer resultat, og en eventuell error
			// dersom resultatet her under- eller overflowet.
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}
	} else if len(args) == 4 { // Dersom argumentlengden er 4, vil det siste argumentet være type
		format := strings.ToLower(args[3]) // La type-argumentet være case-insensitive
		switch f := format; f {
			case "int8":
				a_8, a_err := strconv.ParseInt(args[1], 10, 8)
				b_8, b_err := strconv.ParseInt(args[2], 10, 8)
				TestError(a_err, b_err)
				a := int8(a_8)
				b := int8(b_8)
				result, err := sum.SumInt8(a, b)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			case "uint32":
				a_u32, a_err := strconv.ParseUint(args[1], 10, 32)
				b_u32, b_err := strconv.ParseUint(args[2], 10, 32)
				TestError(a_err, b_err)
				a := uint32(a_u32)
				b := uint32(b_u32)
				result, err := sum.SumUint32(a, b)
				if err != nil { // Uint32 vil returnere error dersom c = a*b
						// mens c/b != a
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			case "int32":
				a_32, a_err := strconv.ParseInt(args[1], 10, 32)
				b_32, b_err := strconv.ParseInt(args[2], 10, 32)
				TestError(a_err, b_err)
				a := int32(a_32)
				b := int32(b_32)
				result, err := sum.SumInt32(a, b)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			case "int64":
				a_64, a_err := strconv.ParseInt(args[1], 10, 64)
				b_64, b_err := strconv.ParseInt(args[2], 10, 64)
				TestError(a_err, b_err)
				a := int64(a_64)
				b := int64(b_64)
				result, err := sum.SumInt64(a, b)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			case "float64":
				a_f64, a_err := strconv.ParseFloat(args[1], 64)
				b_f64, b_err := strconv.ParseFloat(args[2], 64)
				TestError(a_err, b_err)
				result, err := sum.SumFloat64(a_f64, b_f64)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			default:
				fmt.Printf("%s is not a supported type\n", f)
		}
	} else {
		fmt.Println("Usages:\n")
		fmt.Println("main_sum 2 2")
		fmt.Println("\tParses as int64 and returns 4")
		fmt.Println("main_sum 4.5 2.1")
		fmt.Println("\tParses as float64 and returns 6.6")
		fmt.Println("main_sum 2 2 int8")
		fmt.Println("\tParses as int8 and returns 4")
		fmt.Println("\nAny of these types can be used as third parameter: int8, uint32, int32, int64, float64")
	}
}


func TestError(a_err, b_err error) {
	// Generisk funksjon for å sjekke errors, og lukke programmet dersom
	// en fins.
	if a_err != nil {
		fmt.Println(a_err)
		os.Exit(0)
	}
	if b_err != nil {
		fmt.Println(b_err)
		os.Exit(0)
	}
}
