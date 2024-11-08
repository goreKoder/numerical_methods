package main //			go run laba_2/main.go

import (
	"fmt"
	"math"
)

func main() {
	// 		Метод Якоби
	var Matr [4][4]float64 = [4][4]float64{
		{4.3000, 0.0100, 0.0037, -0.0027},
		{0.0217, 3.4000, 0.0090, 0.0027},
		{0.0270, 0.0207, 2.5000, 0.0080},
		{0.0324, 0.0260, 0.0197, 1.6000},
	}
	var Stolb [4]float64 = [4]float64{0.6632, 2.7779, 2.5330, 1.9285}
	// var MatrD [4][4]float64
	// var MatrL [4][4]float64
	// var MatrU [4][4]float64
	// Начальные условия
	var x [4]float64       // Вектор переменных
	var x_new [4]float64   // Вектор для хранения новых значений переменных
	const epsilon = 0.0001 // Порог сходимости
	iterations := 1        // Счетчик итераций

	// Итерационный процесс
	for {
		// Обновляем значения переменных
		for i := 0; i < 4; i++ {
			sum := 0.0
			// Суммируем все элементы, кроме диагонального
			for j := 0; j < 4; j++ {
				if i != j {
					sum += Matr[i][j] * x[j]
				}
			}
			// Вычисляем новое значение переменной
			x_new[i] = (Stolb[i] - sum) / Matr[i][i]
		}

		// Проверка на сходимость
		converged := true
		for i := 0; i < 4; i++ {
			if math.Abs(x_new[i]-x[i]) > epsilon {
				converged = false
				break
			}
		}

		// Если все значения сошлись, выходим из цикла
		if converged {
			break
		}

		// Обновляем вектор переменных для следующей итерации
		copy(x[:], x_new[:])
		iterations++
	}

	// Выводим результат
	fmt.Printf("Решение системы уравнений после %d итераций:\n", iterations)
	for i := 0; i < 4; i++ {
		fmt.Printf("x[%d] = %.6f\n", i, x[i])
	}
	fmt.Println("Колличкство итераций: ", iterations)

	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		if i == j {
	// 			MatrD[i][j] = Matr[i][j]
	// 			MatrU[i][j] = 0
	// 			MatrL[i][j] = 0
	// 		} else if i > j {
	// 			MatrD[i][j] = 0
	// 			MatrU[i][j] = 0
	// 			MatrL[i][j] = Matr[i][j]
	// 		} else {
	// 			MatrD[i][j] = 0
	// 			MatrU[i][j] = Matr[i][j]
	// 			MatrL[i][j] = 0
	// 		}
	// 	}
	// }
	// fmt.Println("Диагональ: ", MatrD)
	// fmt.Println("Нижний треугольник: ", MatrL)
	// fmt.Println("Верхний треугольник: ", MatrU)
	//x(n+1) = Bx(x) + d, где B — матрица, а d — вектор
}

//			go run laba_2/main.go
