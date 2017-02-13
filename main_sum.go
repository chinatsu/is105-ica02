package main

import "./sum"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
	args := os.Args
	if len(args) == 3 {
		format := "int"
		var a_int64, b_int64 int64
		var a, b float64

		a_int64, a_err := strconv.ParseInt(args[1], 10, 64)
		b_int64, b_err := strconv.ParseInt(args[2], 10, 64)
		if a_err != nil || b_err != nil {
			a, a_err = strconv.ParseFloat(args[1], 64)
			b, b_err = strconv.ParseFloat(args[2], 64)
			format = "float"
			if a_err != nil || b_err != nil {
				fmt.Println("That's numberwang!")
				os.Exit(0)
			}
		}
		if format == "float" {
			fmt.Println(sum.SumFloat64(a, b))
		} else if format == "int" {
			fmt.Println(sum.SumInt64(a_int64, b_int64))
		}
	} else if len(args) == 4 {
		format := strings.ToLower(args[3])
		switch f := format; f {
			case "int8":
				a_8, a_err := strconv.ParseInt(args[1], 10, 8)
				b_8, b_err := strconv.ParseInt(args[2], 10, 8)
				TestError(a_err, b_err)
				a := int8(a_8)
				b := int8(b_8)
				fmt.Println(sum.SumInt8(a, b))
			case "uint32":
				a_u32, a_err := strconv.ParseUint(args[1], 10, 32)
				b_u32, b_err := strconv.ParseUint(args[2], 10, 32)
				TestError(a_err, b_err)
				a := uint32(a_u32)
				b := uint32(b_u32)
				fmt.Println(sum.SumUint32(a, b))
			case "int32":
				a_32, a_err := strconv.ParseInt(args[1], 10, 32)
				b_32, b_err := strconv.ParseInt(args[2], 10, 32)
				TestError(a_err, b_err)
				a := int32(a_32)
				b := int32(b_32)
				fmt.Println(sum.SumInt32(a, b))
			case "int64":
				a_64, a_err := strconv.ParseInt(args[1], 10, 64)
				b_64, b_err := strconv.ParseInt(args[2], 10, 64)
				TestError(a_err, b_err)
				a := int64(a_64)
				b := int64(b_64)
				fmt.Println(sum.SumInt64(a, b))
			case "float64":
				a_f64, a_err := strconv.ParseFloat(args[1], 64)
				b_f64, b_err := strconv.ParseFloat(args[2], 64)
				TestError(a_err, b_err)
				fmt.Println(sum.SumFloat64(a_f64, b_f64))
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
	if a_err != nil {
		fmt.Println(a_err)
		os.Exit(0)
	}
	if b_err != nil {
		fmt.Println(b_err)
		os.Exit(0)
	}
}
