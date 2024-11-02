package main //   	go run laba_1/main.go

import "fmt"

func main() {
	var Matr [3][3]float64 = [3][3]float64{
		{1.54, 1.70, 1.62},
		{3.69, 3.73, 3.59},
		{2.45, 2.43, 2.25},
	}
	var Stolb [3]float64 = [3]float64{-1.97, -3.74, -2.26}
	var znamenatel float64
	for g := 0; g < 3; g++ {
		i := g
		if i == 2 {
			i = 1
			znamenatel = Matr[i+1][1] / Matr[i][1]
		} else if Matr[i+1][0] >= Matr[i][0] {
			znamenatel = Matr[i+1][0] / Matr[i][0]
		} else {
			znamenatel = Matr[i][0] / Matr[i+1][0]
		}
		for j := 0; j < 3; j++ {
			Matr[i+1][j] -= Matr[i][j] * znamenatel
		}
		Stolb[i+1] -= Stolb[i] * znamenatel
	}
	fmt.Println(Matr)
	fmt.Println(Stolb)
}
