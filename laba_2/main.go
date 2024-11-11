package main //			go run laba_2/main.go
import "fmt"

func main() {
	// 		Метод Якоби
	var Matr [4][4]float64 = [4][4]float64{
		{4.3000, 0.0100, 0.0037, -0.0027},
		{0.0217, 3.4000, 0.0090, 0.0027},
		{0.0270, 0.0207, 2.5000, 0.0080},
		{0.0324, 0.0260, 0.0197, 1.6000},
	}
	var Stolb [4]float64 = [4]float64{0.6632, 2.7779, 2.5330, 1.9285}
	// Начальные условия
	var x [4]float64 = [4]float64{0, 0, 0, 0} // Вектор переменных
	var x_new [4]float64                      // Вектор для хранения новых значений переменных
	iterations := 0                           // Счетчик итераций
	for j := 0; j < 3; j++ {
		for i := 0; i < 4; i++ {
			x_new[i] = (Stolb[i] + Matr[i][i]*x[i] /* Сложением нивелирую математику*/ - Matr[i][0]*x[0] - Matr[i][1]*x[1] - Matr[i][2]*x[2] - Matr[i][3]*x[3]) / Matr[i][i]
		}
		x = x_new
		iterations++
	}
	fmt.Println("Решение методом Якоби: ", x)
	fmt.Println("Кол-во операций: ", iterations)
	iterations = 0

	//			Метод зейделя

	x = [4]float64{0, 0, 0, 0}
	for j := 0; j < 3; j++ {
		for i := 0; i < 4; i++ {
			x_new[i] = (Stolb[i] + Matr[i][i]*x[i] /* Сложением нивелирую математику*/ - Matr[i][0]*x[0] - Matr[i][1]*x[1] - Matr[i][2]*x[2] - Matr[i][3]*x[3]) / Matr[i][i]
			x[i] = x_new[i]
		}
		iterations++
	}
	fmt.Println("Решение методом Зейделя: ", x)
	fmt.Println("Кол-во операций: ", iterations)

	//		Метод верхней релаксации (обобщенный метод Зейде-ля)

	fmt.Println("Решение методом Зейделя: ")
	x = [4]float64{0, 0, 0, 0}
	for w := 0.2; w <= 1.8; w += 0.2 {
		iterations = 0
		for j := 0; j < 30; j++ {
			for i := 0; i < 4; i++ {
				x_new[i] = w*((Stolb[i]+Matr[i][i]*x[i] /*!!!*/ -Matr[i][0]*x[0]-Matr[i][1]*x[1]-Matr[i][2]*x[2]-Matr[i][3]*x[3])/Matr[i][i]) + (1-w)*x[i]
				x[i] = x_new[i]
			}
			iterations++
		}
		for f := 0; f < 4; f++ {
			fmt.Print("|  X", f+1, ": ")
			fmt.Printf("%.3f", x[f])
		}
		x = [4]float64{0, 0, 0, 0}
		fmt.Print("|  W: ")
		fmt.Printf("%.1f", w)
		fmt.Println("|  Кол-во операций: ", iterations)
	}
}

//			go run laba_2/main.go
