package main //			go run laba_3/main.go

import (
	"fmt"
	"math"
)

func main() {
	var Matr [2][2]float64 = [2][2]float64{
		{1.06, 0.991},
		{0.994, 0.943},
	}
	var Stolb [2]float64 = [2]float64{2.54, 2.44}
	//		Метод регуляризации
	MatrT := [2][2]float64{
		{1.06, 0.994},
		{0.991, 0.943},
	}
	var AtA [2][2]float64 //транспонированная матрица умноженная на матрицу
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum := 0.0
			for k := 0; k < 2; k++ {
				sum += MatrT[i][k] * Matr[k][j]
			}
			AtA[i][j] = sum
		}
	}
	var AtB [2]float64 //транспонированная матрица умноженная на вектор b
	for i := 0; i < 2; i++ {
		sum := 0.0
		for j := 0; j < 2; j++ {
			sum += MatrT[i][j] * Stolb[j]
		}
		AtB[i] = sum
	}
	fmt.Println(AtA, AtB)
	// юзаю метод гауса чтобы выбить ответ
	for i := 0; i < 2; i++ {
		for k := i + 1; k < 2; k++ {
			factor := Matr[k][i] / Matr[i][i]
			for j := i; j < 2; j++ {
				Matr[k][j] -= factor * Matr[i][j]
			}
			AtB[k] -= factor * AtB[i]
		}
	}
	// Обратный ход
	var x [2]float64
	for i := 1; i >= 0; i-- {
		x[i] = AtB[i]
		for j := i + 1; j < 2; j++ {
			x[i] -= Matr[i][j] * x[j]
		}
		x[i] /= Matr[i][i]
	}
	fmt.Println("Решение методом Гауса после регуляризации: ")
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %.6f\n", i, x[i])
	}

	//		Метод вращения

	for i := 0; i < len(Matr)-1; i++ {
		for j := i + 1; j < len(Matr); j++ {
			c, s := givensRotation(Matr[i][i], Matr[j][i])
			applyGivensRotation(Matr, i, j, c, s)
		}
	}
	fmt.Println("Матрица после применения метода Гивенса:")
	for _, row := range Matr {
		fmt.Println(row)
	}
	// юзаю метод гауса чтобы выбить ответ
	for i := 0; i < 2; i++ {
		for k := i + 1; k < 2; k++ {
			factor := Matr[k][i] / Matr[i][i]
			for j := i; j < 2; j++ {
				Matr[k][j] -= factor * Matr[i][j]
			}
			AtB[k] -= factor * AtB[i]
		}
	}
	// Обратный ход
	for i := 1; i >= 0; i-- {
		x[i] = AtB[i]
		for j := i + 1; j < 2; j++ {
			x[i] -= Matr[i][j] * x[j]
		}
		x[i] /= Matr[i][i]
	}
	fmt.Println("Решение методом Гауса после вращения: ")
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %.6f\n", i, x[i])
	}
}

// Функция для выполнения вращения Гивенса
func givensRotation(a, b float64) (float64, float64) {
	if b == 0 {
		return 1, 0
	}
	r := math.Hypot(a, b) // Вычисляем гипотенузу
	c := a / r            // Косинус угла
	s := -b / r           // Синус угла
	return c, s
}

// Функция для применения вращения Гивенса к матрице
func applyGivensRotation(A [2][2]float64, i, j int, c, s float64) {
	n := len(A)
	for k := 0; k < n; k++ {
		temp := c*A[i][k] + s*A[j][k]
		A[j][k] = -s*A[i][k] + c*A[j][k]
		A[i][k] = temp
	}
}

//			go run laba_3/main.go
