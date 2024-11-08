package main //   	go run laba_1/gaus.go

import (
	"fmt"
	"math"
)

func main() {

	var Matr [3][3]float64 = [3][3]float64{
		{1.54, 1.70, 1.62},
		{3.69, 3.73, 3.59},
		{2.45, 2.43, 2.25},
	}
	var Stolb [3]float64 = [3]float64{-1.97, -3.74, -2.26}
	var Matr2 [3][3]float64 = Matr
	var Stolb2 [3]float64 = Stolb

	// Прямой ход метода Гаусса
	for i := 0; i < 3; i++ {
		// Поиск максимального элемента в столбце
		maxRow := i
		for k := i + 1; k < 3; k++ {
			if math.Abs(Matr[k][i]) > math.Abs(Matr[maxRow][i]) {
				maxRow = k
			}
		}

		// Меняем местами текущую строку и строку с максимальным элементом
		if maxRow != i {
			for k := i; k < 3; k++ {
				Matr[i][k], Matr[maxRow][k] = Matr[maxRow][k], Matr[i][k]
			}
			Stolb[i], Stolb[maxRow] = Stolb[maxRow], Stolb[i]
		}

		// Обнуляем элементы ниже главной диагонали
		for k := i + 1; k < 3; k++ {
			factor := Matr[k][i] / Matr[i][i]
			for j := i; j < 3; j++ {
				Matr[k][j] -= factor * Matr[i][j]
			}
			Stolb[k] -= factor * Stolb[i]
		}
	}

	// Обратный ход
	var x [3]float64
	for i := 3 - 1; i >= 0; i-- {
		x[i] = Stolb[i]
		for j := i + 1; j < 3; j++ {
			x[i] -= Matr[i][j] * x[j]
		}
		x[i] /= Matr[i][i]
	}

	fmt.Println("Решение методом Гауса: ")
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %.6f\n", i, x[i])
	}

	// var Matr [3][3]float64 = [3][3]float64{
	// 	{1.54, 1.70, 1.62},
	// 	{3.69, 3.73, 3.59},
	// 	{2.45, 2.43, 2.25},
	// }
	// var Stolb [3]float64 = [3]float64{-1.97, -3.74, -2.26}
	// var Matr2 [3][3]float64 = Matr

	// var Stolb2 [3]float64 = Stolb
	// // метод Гауса
	// var znamenatel float64
	// fmt.Println(Matr[0][1])
	// for g := 0; g < 3; g++ {
	// 	i := g
	// 	if i == 2 {
	// 		i = 1
	// 		znamenatel = Matr[i+1][1] / Matr[1][1]
	// 	} else {
	// 		znamenatel = Matr[i+1][0] / Matr[0][0]
	// 	}
	// 	for j := 0; j < 3; j++ {
	// 		if g == 2 {
	// 			Matr[i+1][j] -= Matr[1][j] * znamenatel
	// 		} else {
	// 			Matr[i+1][j] -= Matr[0][j] * znamenatel
	// 		}
	// 	}
	// 	Stolb[i+1] -= Stolb[0] * znamenatel
	// }
	// fmt.Println("Преобразованная матрица: ", Matr)
	// // fmt.Println("Столбец свободных членов: ", Stolb)
	// x3 := Stolb[2] / Matr[2][2]
	// x2 := Stolb[1] / (Matr[1][1] * Matr[1][2] * x3)
	// x1 := Stolb[0] / (Matr[0][0] * Matr[0][1] * x2 * Matr[0][2] * x3)
	// fmt.Println("Решение методом Гауса: ")
	// fmt.Println("X1 = ", x1, "X2 = ", x2, "X3 = ", x3)

	//метод Холецкого
	var znamenatel float64
	var Matr2tr [3][3]float64 = Matr2
	for i := 0; i < len(Matr2); i++ {
		for j := 0; j < len(Matr2); j++ {
			Matr2tr[i][j] = Matr2[j][i]
		}
	}
	var Matr3 [3][3]float64 = Matr2
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				Matr3[i][j] += Matr2[i][k] * Matr2tr[k][j]
			}
		}
	}
	// fmt.Println("Преобразованная матрица2: ", Matr3)
	Matr3[0][0] = math.Sqrt(Matr3[0][0])
	znamenatel = Matr3[1][0]  //!!!
	Matr3[1][0] = Matr3[0][1] //!!!
	Matr3[1][0] /= Matr3[0][0]
	Matr3[1][1] = math.Sqrt(Matr3[1][1] - (Matr3[1][0] * Matr3[1][0]))
	Matr3[1][0] = znamenatel / Matr3[0][0] //!!!
	Matr3[2][0] /= Matr3[0][0]
	Matr3[2][1] = (Matr3[2][1] - Matr3[1][0]*Matr2[2][0]) / Matr3[2][0]
	Matr3[2][2] = math.Sqrt(Matr3[2][2] - Matr2[2][0]*Matr3[2][0] - Matr3[2][1]*Matr3[2][1])
	Matr3[0][1] = 0
	Matr3[0][2] = 0
	Matr3[1][2] = 0
	// fmt.Println("Преобразованная матрица2: ", Matr3)

	var Matr3tr [3][3]float64 = Matr3
	for i := 0; i < len(Matr3); i++ {
		for j := 0; j < len(Matr3); j++ {
			Matr3tr[i][j] = Matr3[j][i]
		}
	}
	y3 := Stolb2[0] / Matr3[0][0]
	y2 := Stolb2[1] / (Matr3[1][1] * Matr3[1][0] * y3)
	y1 := Stolb2[0] / (Matr3[2][0] * Matr3[2][1] * y2 * Matr[2][2] * y3)
	// fmt.Println("y1 = ", y1, "y2 = ", y2, "y3 = ", y3)

	fmt.Println("Решение методом Холецкого: ")
	X3 := y3 / Matr3tr[2][2]
	X2 := y2 / (Matr3tr[1][1] * Matr3tr[1][2] * X3)
	X1 := y1 / (Matr3tr[0][0] * Matr3tr[0][1] * X2 * Matr3tr[0][2] * X3)
	Matr[1][1] = X1
	fmt.Println("X1 = ", x[0], "X2 = ", x[1], "X3 = ", x[2])
}

//				go run laba_1/gaus.go
