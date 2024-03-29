package main

import (
	"fmt"
	"time"
)

// завершение горутины с помощью закрытия канала
func main() {
	c := make(chan int) // создание канала для отправки чисел

	go func() { // запуск горутины
		for {
			val, ok := <-c // получения данных из канала
			if !ok {       // если канал отправляет ok == false
				return // работа горутины завершается
			}
			fmt.Println("Получено значение: ", val) // вывод данных из канала
		}
	}()

	// цикл отправки чисел от 0 до 9 в канал
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Second) // задержка для имитации работы
	}
	close(c) // закрытие канала после выполнения цикла
	fmt.Println("Канал закрыт")

}
