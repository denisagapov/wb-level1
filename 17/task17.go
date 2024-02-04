package main

import "fmt"

// функция выполняет бинарный поиск элемента в отсортированном слайсе
func binSearch(arr []int, number int) int {
	left := 0             // левая граница слайса
	right := len(arr) - 1 // правая граница

	for left <= right { //цикл выполняетсяпока левая граница не будет <= правой
		mid := left + (right-left)/2 // вычисляем середину
		// возвращаем индекс середины, если искомы элемент равен среднему элементу
		if arr[mid] == number {
			return mid
		}
		// если средний элемент меньше искомого, сдвигаем левую границу за середину
		if arr[mid] < number {
			left = mid + 1
			// иначе сдвигаем правую границу перед серединой
		} else {
			right = mid - 1
		}
	}
	// если элемент не найде, возвращаем -1
	return -1
}

func main() {
	// отсортированный слайс
	arr := []int{1, 3, 4, 5, 13, 20, 25, 40, 42, 44, 53}
	// искомое число
	number := 42
	index := binSearch(arr, number)
	// если функция не возвращает -1 (элемент не найдет)
	if index != -1 {
		fmt.Printf("Найден элемент %d с индексом %d\n", number, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", number)
	}
}