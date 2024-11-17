package main //			go run laba_4/main.go
import (
	"fmt"
	"math"
)

func main() {
	//		Метод простых итераций
	var x [10]float64
	var y [10]float64
	for i := 1; i < 10; i++ {
		x[i] = 1 - (math.Sin(y[i-1]) / 2)
		y[i] = 0.7 - math.Cos(x[i-1]-1)
	}
	fmt.Println("Решение методом простых итераций: ")
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %.6f;	y[%d] = %.6f\n", i, x[i], i, y[i])
	}

	//			Метод ньютона

	x0 := 1.0
	y0 := -0.5
	// Начальные значения
	// x0 := 1.0
	// y0 := -0.5
	tol := 1e-6
	maxIter := 10

	// Запускаем метод Ньютона
	X, Y := tonMethod(x0, y0, tol, maxIter)

	// Выводим результат
	fmt.Printf("Решение: x = %.6f, y = %.6f\n", X, Y)
	// for i := 0; i < 10; i++ {
	// 	var J [2][2]float64 = [2][2]float64{
	// 		{2, math.Cos(y[i])},
	// 		{-math.Sin(x[i] - 1), 1},
	// 	}
	// 	var F [2]float64 = [2]float64{math.Sin(y[i]) + 2*x[i] - 2, math.Cos(x[i]-1) + y[i] - 0.7}
	// }
}

// Определяем функции f1 и f2
func f1(x, y float64) float64 {
	return math.Sin(y) + 2*x - 2
}

func f2(x, y float64) float64 {
	return math.Cos(x-1) + y - 0.7
}

// Определяем частные производные для якобиана
func df1dx() float64 {
	return 2
}

func df1dy(y float64) float64 {
	return math.Cos(y)
}

func df2dx(x float64) float64 {
	return -math.Sin(x - 1)
}

func df2dy() float64 {
	return 1
}

// Метод Ньютона для решения системы уравнений
func tonMethod(x0, y0 float64, tol float64, maxIter int) (float64, float64) {
	x, y := x0, y0

	for i := 0; i < maxIter; i++ {
		// Вычисляем значения функций
		f1Val := f1(x, y)
		f2Val := f2(x, y)

		// Вычисляем якобиан
		j11 := df1dx()
		j12 := df1dy(y)
		j21 := df2dx(x)
		j22 := df2dy()

		// Вычисляем определитель якобиана
		detJ := j11*j22 - j12*j21
		if math.Abs(detJ) < 1e-10 {
			fmt.Println("Якобиан близок к нулю, метод не сходится.")
			return x, y
		}

		// Вычисляем обратный якобиан
		invJ11 := j22 / detJ
		invJ12 := -j12 / detJ
		invJ21 := -j21 / detJ
		invJ22 := j11 / detJ

		// Вычисляем новые значения x и y
		xNew := x - (invJ11*f1Val + invJ12*f2Val)
		yNew := y - (invJ21*f1Val + invJ22*f2Val)

		// Проверяем на сходимость
		if math.Abs(xNew-x) < tol && math.Abs(yNew-y) < tol {
			return xNew, yNew
		}

		x, y = xNew, yNew
	}

	fmt.Println("Достигнуто максимальное количество итераций.")
	return x, y
}

//			go run laba_4/main.go
