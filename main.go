package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	var a = [][]int{
		{5, 4, 3, 7},
		{12, 3, 4, 6},
		{9, 2, 8, 5},
		{7, 6, 7, 2},
		{6, 6, 7, 2},
		{1, 6, 7, 2},
		{4, 6, 7, 2},
		{2, 6, 7, 2},
		{10, 6, 7, 2},
	}
	fmt.Print(a[5][0])
	/* output each array element's value */
	/*var i, j int

	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}*/

	var dam = [] int{7, 7, 4, 7}

	mins := min(dam, 8)
	fmt.Println("Minimele sunt:")
	for i := 0; i < len(mins); i++ {
		fmt.Printf("%d ", mins[i])
	}

	flood(a, 5, 0)

	fmt.Println("\npuncte vizitate")
	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d]=%d\n", i, m[i])
	}
}

func min(dam []int, level int) []int {
	var mins []int
	for i := 0; i < len(dam); i++ {
		if level >= dam[i] {
			mins = append(mins, dam[i])
		}
	}
	return mins
}

var m [9]int

func flood(a [][]int, row int, col int) {
	h_flood := 7
	if row < 0 || row > 8 {
		return
	}
	if col < 0 || col > 8 {
		return
	}
	if m[row] == -1 {
		return
	}
	if a[row][col] <= h_flood {
		m[row] = -1
	} else {
		return
	}
	fmt.Printf("\nGo North---> row=%d", row-1)
	flood(a, row-1, col)
	fmt.Printf("\nGo South---> row=%d", row+1)
	flood(a, row+1, col)
}
