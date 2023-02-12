package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your list of numbers in order separated by comma (e.g. 1,2,3,4,5): ")
	var rawList string
	fmt.Scan(&rawList)
	fmt.Println("Given numbers ", rawList)

	numListStr := strings.Split(rawList, ",")
	numList := make([]int, len(numListStr))
	for i, numStr := range numListStr {
		numList[i], _ = strconv.Atoi(numStr)
	}

	middleIndex := (len(numList) + 1) / 2
	leftList := numList[:middleIndex]
	rightList := numList[middleIndex:]

	fmt.Printf("Split into %v and %v \n", leftList, rightList)

	Ra, sortedLleftList := sortCount(leftList)
	fmt.Println("Ra = ", Ra, " ", sortedLleftList)
	Rb, sortedRightList := sortCount(rightList)
	fmt.Println("Rb = ", Rb, " ", sortedRightList)
	Rx, sortedList := mergeCount(leftList, rightList)
	fmt.Println("Rx = ", Rx)
	fmt.Println("L = ", sortedList)
	fmt.Println("Total inversions = ", Rx+Ra+Rb)

}

func sortCount(list []int) (int, []int) {
	count := 0
	n := len(list)

	if n == 1 {
		return 0, list
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i] // swaps position
				count++
			}

		}
	}

	return count, list

}

func mergeCount(listL []int, listR []int) (int, []int) {

	count := 0
	LPointer := 0
	RPointer := 0
	var outputList []int

	for LPointer < len(listL) && RPointer < len(listR) { //both list not reach end (pointer not exceed len())

		if listL[LPointer] <= listR[RPointer] { // ai < bj
			outputList = append(outputList, listL[LPointer]) //append smallest
			LPointer++
		} else { // ai > bj
			outputList = append(outputList, listR[RPointer]) //append smallest
			RPointer++
			count += len(listL) - LPointer //inc count by elem. remaning in A(listL)
		}

	}

	//appends the remaning elements (1 would have already been completely empty due to for loop)
	outputList = append(outputList, listL[LPointer:]...)
	outputList = append(outputList, listR[RPointer:]...)

	return count, outputList

}
