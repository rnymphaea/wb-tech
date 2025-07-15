# L1.17 Бинарный поиск
## Задание
Реализовать алгоритм бинарного поиска встроенными методами языка. Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.
## Пример работы программы
#### Справка
```bash
$ go run binsearch.go --help
```
```
Usage of /path/to/exe/binsearch:
  -debug
    	show logs to check the correctness of execution
  -random
    	use the random array to check the correctness of algorithm
```

#### Стандартный вывод 
```bash
$ go run binsearch.go
```
```
Enter the size of array: 5
Enter 5 elements of array: 23 53 61 96 34
Enter the target to find: 61
Sorted arr:  [23 34 53 61 96]
Index of 61 in array is: 3
```

#### Режим отображения логов
```bash 
$ go run binsearch.go --debug --random
```
```
Debug mode started
Random: true
[DEBUG] initial array: [1 4 10 14 20 21 21 28 29 30 30 32 34 35 41 43 45]
[DEBUG] target: 21
[DEBUG] iteration #1
[DEBUG] left = 0, right = 17, mid = 8
[DEBUG] arr[mid] = 29
[DEBUG] target < arr[mid]. Finding in [1 4 10 14 20 21 21 28]
[DEBUG] iteration #2
[DEBUG] left = 0, right = 8, mid = 4
[DEBUG] arr[mid] = 20
[DEBUG] target > arr[mid]. Finding in [21 21 28]
[DEBUG] iteration #3
[DEBUG] left = 5, right = 8, mid = 6
[DEBUG] arr[mid] = 21
Index of 21 in array is: 6
```
## Тестирование
```bash
$ go test -v
```
