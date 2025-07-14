# L1.7 Конкурентная запись в map
## Задание
Реализовать безопасную для конкуренции запись данных в структуру map.
## Пример работы программы
```
├── custom_map.go // с использованием мьютекса
└── sync_map.go   // с использованием sync.Map
```
Вывод программ не отличается. Поэтому для демонстрации работы будет использован файл custom_map.go
#### Справка
```bash
$ go run custom_map.go --help
```
```
Usage of /path/to/exe/custom_map:
  -debug
    	show logs to check the correctness of execution
  -goroutines int
    	amount of goroutines to write in map (default 10)
```

#### Стандартный вывод 
```bash
$ go run custom_map.go
```
```
map[0:0 1:8 2:16 3:24 4:32]
```

#### Режим отображения логов
```bash 
$ go run --race custom_map.go --debug --goroutines=3
```
```
Debug mode started
Amount of goroutines: 3
[DEBUG] goroutine 0: set: key - 0, value - 0
[DEBUG] goroutine 1: set: key - 0, value - 0
[DEBUG] goroutine 0: get: key - 0, value - 0
[DEBUG] goroutine 1: get: key - 0, value - 0
[DEBUG] goroutine 0: set: key - 1, value - 0
[DEBUG] goroutine 1: set: key - 1, value - 1
[DEBUG] goroutine 0: get: key - 1, value - 0
[DEBUG] goroutine 2: set: key - 0, value - 0
[DEBUG] goroutine 0: set: key - 2, value - 0
[DEBUG] goroutine 1: get: key - 1, value - 1
[DEBUG] goroutine 0: get: key - 2, value - 0
[DEBUG] goroutine 1: set: key - 2, value - 2
[DEBUG] goroutine 0: set: key - 3, value - 0
[DEBUG] goroutine 2: get: key - 0, value - 0
[DEBUG] goroutine 1: get: key - 2, value - 2
[DEBUG] goroutine 2: set: key - 1, value - 2
[DEBUG] goroutine 0: get: key - 3, value - 0
[DEBUG] goroutine 1: set: key - 3, value - 3
[DEBUG] goroutine 2: get: key - 1, value - 2
[DEBUG] goroutine 0: set: key - 4, value - 0
[DEBUG] goroutine 1: get: key - 3, value - 3
[DEBUG] goroutine 0: get: key - 4, value - 0
[DEBUG] goroutine 1: set: key - 4, value - 4
[DEBUG] goroutine 2: set: key - 2, value - 4
[DEBUG] goroutine 1: get: key - 4, value - 4
[DEBUG] goroutine 2: get: key - 2, value - 4
[DEBUG] goroutine 2: set: key - 3, value - 6
[DEBUG] goroutine 2: get: key - 3, value - 6
[DEBUG] goroutine 2: set: key - 4, value - 8
[DEBUG] goroutine 2: get: key - 4, value - 8
map[0:0 1:2 2:4 3:6 4:8]
```
