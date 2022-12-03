package sum

// func sum(a int, b int) int {
// 	return a + b
// }

func sum(xs ...int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
