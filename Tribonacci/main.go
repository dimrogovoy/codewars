package main

import "fmt"

func Xbonacci(signature []float64, n int) []float64 {
	array := signature[:]
	l := len(array)
	for j := l; j < n; j++ {
		var total float64
		for i := j - 1; i >= j-l; i-- {
			total += array[i]
		}
		array = append(array, total)
	}
	return array[:n]
}

func Tribonacci(signature [3]float64, n int) []float64 {
	array := signature[:]
	//switch {
	//case n == 0:
	//	return array
	//case n == 1:
	//	return append(array, signature[0])
	//case n == 2:
	//	return append(array, signature[:2]...)
	//}

	//array = append(array, signature[:n]...)
	for j := 3; j < n; j++ {
		array = append(array, array[j-1]+array[j-2]+array[j-3])
	}
	return array[:n]

}

func main() {

	fmt.Println(Xbonacci([]float64{1, 1, 1, 1}, 10))
	fmt.Println(Xbonacci([]float64{0, 0, 0, 0, 1}, 10))
	fmt.Println(Xbonacci([]float64{1, 0, 0, 0, 0, 0, 1}, 10))
	//fmt.Println(Tribonacci([3]float64{20, 2, 18}, 3))
	//fmt.Println(Tribonacci([3]float64{17, 9, 15}, 2))
}
