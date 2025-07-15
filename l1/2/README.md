# L1.2 Конкурентное возведение в квадрат
## Задание
Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
## Пример работы программы
#### Стандартный вывод 
```bash
$ go run concurrent_squaring.go
```
```
4
100
64
36
16
```
```bash
$ go run concurrent_squaring.go --debug
```
```
Debug mode started
Initial array: [2 4 6 8 10]
[DEBUG] Goroutine 4 received value 10. After processing got: 100
[DEBUG] Goroutine 3 received value 8. After processing got: 64
[DEBUG] Goroutine 0 received value 2. After processing got: 4
[DEBUG] Goroutine 1 received value 4. After processing got: 16
[DEBUG] Goroutine 2 received value 6. After processing got: 36
```
