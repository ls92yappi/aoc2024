package main

import "fmt"
import . "github.com/ls92yappi/aoc"

///////////////////////////////////////////////////////////

func main() {
	list1 := []int{10, 15}
	list3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list4 := []int{9, 8, 7, 5, 11}
	fmt.Println(LCM(list1))
	fmt.Println(LCM(list3))
	// Below shows how to call a variadic using a slice
	fmt.Println(LCMv(list3[0], list3[1], list3[2:]...))
	fmt.Println(LCM(list4))
	fmt.Printf("%v\n", list4)
	fmt.Printf("%v\n", FloorPow2(27720))
	fmt.Printf("%v\n", CeilPow2(27720))
	fmt.Println(Sum(list3))
	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(Prod(list4))
	fmt.Printf("%d %d %d %d\n", Abs(-7), Abs(3), Abs(2), Abs(0))
	fmt.Printf("%d %d %d %d %d %d %d\n", Min2(-4,-3), Min2(-3,5), Min2(-1,-2), Min2(-1,0), Min2(1,7), Min2(7,2), Min2(3,-7))
	fmt.Printf("%d %d %d %d %d %d %d\n", Max2(-4,-3), Max2(-3,5), Max2(-1,-2), Max2(-1,0), Max2(1,7), Max2(7,2), Max2(3,-7))
	//fmt.Printf("%d %d %d %d %d %d %d\n", WhichMin2(-4,-3), WhichMin2(-3,5), WhichMin2(-1,-2), WhichMin2(-1,0), WhichMin2(1,7), WhichMin2(7,2), WhichMin2(3,-7))
	fmt.Println(Min(list4))
	fl := Factor(27720*54321)
	fmt.Println("27720*54321",fl,Sum(fl))

	s := "3 1 4 1 5 9"
	il,ni,_ := IntSlice(s, " ")
	fmt.Printf("%d items, namely %v\n", ni, il)
}
