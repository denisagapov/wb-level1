package main

import (
	"fmt"
)

// функция удваивает каждое полученное число и отправляет результат в канал out
func double(in <-chan int, out chan<- int) {
	for x := range in { // читаем каждое значение из канала in
		out <- x * 2 // удваиваем значение и отправляем в канал out
	}
	close(out) // закрываем канал out после обработки всех чисел
}

func main() {
	in := make(chan int)             // создаем канал для входных значений
	out := make(chan int)            // создаем канал для выходных значений
	numbers := [5]int{1, 2, 3, 4, 5} // массив чисел для обработки

	// анонимная функция для отправки чисел в канал
	go func() {
		for _, x := range numbers { // перебираем все числа массива
			in <- x // отправляем каждое число в канал in
		}
		close(in) // закрываем канал in после отправки всех чисел
	}()

	go double(in, out) // запускаем функцию double в горутине

	for result := range out { // читаем каждый результат из канала out
		fmt.Println(result) // выводим результат в консоль
	}
}