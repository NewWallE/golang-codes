package main

import "sort"

func main() {

}

func numMovesStones(a int, b int, c int) []int {
	stones := []int{a, b, c}
	sort.Ints(stones)

	max := stones[2] - stones[0] - 2

	if max == 0 {
		return []int{0, 0}
	}

	min := 2
	if stones[0]+1 == stones[1] ||
		stones[1]+1 == stones[2] ||
		stones[0]+2 == stones[1] ||
		stones[1]+2 == stones[2] {
		min = 1
	}

	return []int{min, max}

}
