// Цель программы: вывести значения температуры в Цельсиях
// Входные данные представлены в Цельсиях и Фарингейтах
package main

import "fmt"

// Объявляем интерфейс для температуры в Цельсиях

type t interface {
	showCelsTemp()
}

// Объявляем структуру температуры в Цельсиях

type CelsiusTemp struct {
	tempC float32
}

// Объявляем структуру температуры в Фарингейтах
type PharyngeteTemp struct {
	tempPh float32
}

// // Объявляем структуру адаптера температуры

type TempAdapter struct {
	PhCelsTemp *PharyngeteTemp
}

// // Объявляем пустую структуру клинта, который будет делать запросы
type client struct {
}

// Функция, показывающая температуру в Цельсиях
// Для структуры CelsiusTemp

func (c *CelsiusTemp) showCelsTemp() {
	// Выводит в консоль значение температуры
	fmt.Printf("The temperature is %.1f degrees\n", c.tempC)
}

// Функция, показывающая температуру в Фарингейтах
// Для структуры PharyngeteTemp

func (ph *PharyngeteTemp) showPharTemp() {
	// Выводит в консоль значение температуры
	fmt.Printf("The temperature is %.1f degrees\n", ph.tempPh)
}

// Функция запроса клиентом температуры
// Соответствующей интерфейсу температуры в Цельсиях
func (cl *client) askCelsTemp(t t) {
	t.showCelsTemp()
}

// Функция адаптера, преобразующая запрос к температуре в Фарингейтах
func (ta *TempAdapter) showCelsTemp() {
	// Преобразуем значение
	ta.PhCelsTemp.tempPh = (ta.PhCelsTemp.tempPh - 32) / 1.8
	ta.PhCelsTemp.showPharTemp()
}

func main() {
	// Создаем клиента
	client := &client{}

	// Создаем экземпляр стуктуры температуры в Цельсиях
	tC := CelsiusTemp{tempC: 18.0}
	fmt.Println("tC:", tC)

	// Создаем экземпляр стуктуры температуры в Фарингейтах
	tF := PharyngeteTemp{tempPh: 78.8}
	fmt.Println("tF:", tF)

	// Создаем адаптера со ссылкой на температуру в Фарингейтах
	tempAdapter := &TempAdapter{PhCelsTemp: &tF}

	// Клиент запрашивает температуру от обоих экземпляров структур
	client.askCelsTemp(&tC)
	client.askCelsTemp(tempAdapter)
}
