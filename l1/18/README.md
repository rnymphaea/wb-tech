# L1.18 Конкурентный счетчик
## Задание
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин). По завершению программы структура должна выводить итоговое значение счётчика.
## Пример работы программы
```
├── atomic.go // с использованием atomic
└── mutex.go  // с использованием мьютекса
```
Вывод программ не отличается. Поэтому для демонстрации работы будет использован файл mutex.go
#### Справка
```bash
$ go run mutex.go --help
```
```
Usage of /path/to/exe/mutex:
  -debug
    	show logs to check the correctness of execution
  -goroutines int
    	amount of goroutines (default 10)
```

#### Стандартный вывод 
```bash
$ go run mutex.go
```
```
Amount of incrementing operations for each goroutine: 20
Result:  200
```

#### Режим отображения логов
```bash 
$ go run --race mutex.go --debug --goroutines=5
```
```
Debug mode started
Amount of goroutines: 5
Amount of incrementing operations for each goroutine: 5
[DEBUG] expected value: 25
[DEBUG] goroutine #3: incrementing counter (current = 0)
[DEBUG] goroutine #0: incrementing counter (current = 0)
[DEBUG] goroutine #3: incrementing counter (current = 1)
[DEBUG] goroutine #0: incrementing counter (current = 2)
[DEBUG] goroutine #3: incrementing counter (current = 3)
[DEBUG] goroutine #0: incrementing counter (current = 4)
[DEBUG] goroutine #3: incrementing counter (current = 5)
[DEBUG] goroutine #0: incrementing counter (current = 6)
[DEBUG] goroutine #3: incrementing counter (current = 7)
[DEBUG] goroutine #0: incrementing counter (current = 8)
[DEBUG] goroutine #1: incrementing counter (current = 0)
[DEBUG] goroutine #4: incrementing counter (current = 1)
[DEBUG] goroutine #1: incrementing counter (current = 11)
[DEBUG] goroutine #1: incrementing counter (current = 13)
[DEBUG] goroutine #1: incrementing counter (current = 14)
[DEBUG] goroutine #1: incrementing counter (current = 15)
[DEBUG] goroutine #4: incrementing counter (current = 12)
[DEBUG] goroutine #2: incrementing counter (current = 0)
[DEBUG] goroutine #4: incrementing counter (current = 17)
[DEBUG] goroutine #4: incrementing counter (current = 19)
[DEBUG] goroutine #2: incrementing counter (current = 18)
[DEBUG] goroutine #4: incrementing counter (current = 20)
[DEBUG] goroutine #2: incrementing counter (current = 21)
[DEBUG] goroutine #2: incrementing counter (current = 23)
[DEBUG] goroutine #2: incrementing counter (current = 24)
Result:  25
```
