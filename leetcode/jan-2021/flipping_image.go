package main

func main() {
	//image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	image := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	//fmt.Println("Image before: ", image)
	image = flipAndInvertImage(image)
	//fmt.Println("Image after flipping: ", image)
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		reverse(row)
		invert(row)
	}
	return A
}

func reverse(A []int) {
	start := 0
	end := len(A) - 1

	for start < end {
		k := A[start]
		A[start] = A[end]
		A[end] = k
		start++
		end--
	}
}

func invert(A []int) {
	for i, e := range A {
		if e == 0 {
			A[i] = 1
		} else {
			A[i] = 0
		}
	}
}
