package main

import "fmt"

func main() {
	var num1, num2 float32
	var opr string
	fmt.Print("Введите числа:")
	fmt.Scan(&num1, &num2)
	fmt.Print("Введите оператор:")
	fmt.Scan(&opr)
	switch opr {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(num1 / num2)
		}
	case "%":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(int(num1) % int(num2))
		}
	default:
		fmt.Println("Ошибка, такой операции не сушествует!")
	}
}
