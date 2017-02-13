package main

import "./algorithms"
import "fmt"
import "math/rand"
import "time"

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	var list []int = []int{}
	for _, v := range rand.Perm(100) {
		list = append(list, v)
	}
	fmt.Println(list)
	algorithms.Bubble_sort_modified(list)
	fmt.Println(list)
}
