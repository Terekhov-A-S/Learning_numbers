package main

import (
	"fmt"
	"math"
)

func main() {

	//var num1 int64 = -4
	//var num2 int64 = 0
	//var num3 int64 = 5
	//var num4 int64 = 16
	//var num5 int64 = 8
	//var num6 int64 = -3

	//printNumberInfo(num1)
	//printNumberInfo(num2)
	//printNumberInfo(num3)
	//printNumberInfo(num4)
	//printNumberInfo(num5)
	//printNumberInfo(num6)
	printNumberInfo(getInput())
}

func getInput() int64 {
	var getNum int64
	fmt.Print("Введите интересующее число: ")
	fmt.Scanln(&getNum)
	return getNum
}

func printNumberInfo(num int64) {
	sqrtNum := math.Sqrt(float64(num))
	truncatedValue := math.Trunc(sqrtNum)

	switch {
	case num > 0 && num%2 == 0 && sqrtNum == truncatedValue:
		fmt.Printf("Число %d положительное.\n", num)
		fmt.Printf("Число %d четное.\n", num)
		fmt.Printf("Квадратный корень числа %d является целым числом и равен %.0f.\n", num, truncatedValue)
	case num > 0 && num%2 != 0 && sqrtNum == truncatedValue:
		fmt.Printf("Число %d положительное.\n", num)
		fmt.Printf("Число %d нечетное.\n", num)
		fmt.Printf("Квадратный корень числа %d является целым числом и равен %.0f.\n", num, truncatedValue)
	case num > 0 && num%2 == 0 && sqrtNum != truncatedValue:
		fmt.Printf("Число %d положительное.\n", num)
		fmt.Printf("Число %d четное.\n", num)
		fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrtNum)
	case num > 0 && num%2 != 0 && sqrtNum != truncatedValue:
		fmt.Printf("Число %d положительное.\n", num)
		fmt.Printf("Число %d нечетное.\n", num)
		fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrtNum)
	case num < 0 && num%2 == 0:
		fmt.Printf("Число %d отрицательное.\n", num)
		fmt.Printf("Число %d четное.\n", num)
	case num < 0 && num%2 != 0:
		fmt.Printf("Число %d отрицательное.\n", num)
		fmt.Printf("Число %d нечетное.\n", num)
	case num == 0:
		fmt.Println("Число равно 0.")
		fmt.Println("Число 0 четное.")
	default:
		fmt.Println("Вы ввели недопустимое число")
	}
}
