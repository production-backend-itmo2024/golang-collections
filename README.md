# golang-collections

В данной домашке мы попрактикуемся с основными коллекциями языка Golang, а так же попробуем написать свою коллекцию, цель данной ЛР познакомиться с синтаксисом языка

## Небольшая методичка по синтаксису

Структура программы
```go
package main // функция main должна находиться в package main

const first = "first" // const - не имеет типа вместо этого она inlineится на стадии компиляции

const ( // можно обьявлять сразу несколько констант
	second = "second"
	third  = "third"
)

const ( // если мы не перезаписываем величину, то она копируется
	one   = 1
	two   // 1
	three // 1
)

const ( // iota специальная константа, которая автоматически увеличивается с каждой строкой
	Sunday    = iota     // 0
	_                    // 1
	Monday               // 2
	Wednesday = iota - 1 // 2
	Thursday             // 3
	Friday               // 4
	Saturday             // 5
)

var variable int64      // аналогично константам можно обьявлять переменные вне функции, но у них должне быть тип
var otherVariable = 123 // обьявление со значнием, by default ставится zero value для int это 0

func functionName(inputValue int) (firstReturnValue int, secondReturnValue error) { // сигнатура функции выглядит следующим образом
	return 0, nil // golang поддерживает возврат нескольких значений из функции
}

type myStruct struct { // обьявление структуры
	privateField int64  // обозначение поля стуктуры это поле не экспортируемое(не видно вне пакета)
	PublicField  string // это поле экспортируемое(видно вне пакета)
} // ТО myStruct - не экспортируемая стуктура с экспортируемым полем :)

func (ms *myStruct) method() { // добавляем к стуртуре метод, теперь у myStruct есть метод method()
	// ms - это ресивер, анало self или this
}

type myInterface interface { // обьявление интерфейса
	myMethod(inputValue int) (returnValue1 string, err error) // похоже на функцию, только без func
} // в golang утиная типизация, те любой тип, который имеет методы интерфейса, имплементит интерфейс

func main() { // Точка входа в любую программу написанную на go
	var stringVar string = "123" // полная форма
	intVar := int64(123)         // обьявление переменной, тип выводится автоматически
	var otherStringVar = "123"   // обьявление переменной, тип выводится автоматически

	_ = intVar
	_ = stringVar
	_ = otherStringVar

	for i := 1; i < 10; i++ { // синтаксис for в си style

	}

	var slice []int64            // обьявление пустного слайса
	slice = make([]int64, 0, 10) // можно так же сконструировать слайс через make

	for index, value := range slice { // можно использовать range для коллекций, пример для slice
		_, _ = index, value
	}

	var mp map[string]int64 // обьявление map
	// после обьявления map нужно обязательно сконструировать его
	mp = make(map[string]int64) // работа с некоструированным mapом приводит к панике

	key := "123"
	value := int64(123)
	nonExistentKey := "321"

	mp[key] = value                     // кладем в map
	otherValue := mp[key]               // достаем из map
	zeroValue := mp[nonExistentKey]     // если данного ключа нет в map, то вернется zero value
	zeroValue, ok := mp[nonExistentKey] // можно использовать с 2мя аргуметнами в ok положится true, если значние есть, false, если нет

	_, _, _ = otherValue, zeroValue, ok
}

```