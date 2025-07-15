# L1.16 Быстрая сортировка (quicksort)
## Задание
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
## Пример работы программы
#### Справка
```bash
$ go run qsort.go --help
```
```
Usage of /path/to/exe/qsort:
  -debug
    	show logs to check the correctness of execution
  -random
    	use the random array to check the correctness of algorithm
```

#### Стандартный вывод 
```bash
$ go run qsort.go
```
```
Enter the size of array: 5
Enter 5 elements of array: 7 248 24 1 9
Reverse order? [y/n] n
Result:  [1 7 9 24 248]
```

#### Режим отображения логов
```bash 
$ go run qsort.go --debug --random
```
```
Debug mode started
Random array: true
[DEBUG] initial array: [19 37 32 44 12 1 18 31 40]
Reverse order? [y/n] y

[DEBUG] current array: [19 37 32 44 12 1 18 31 40]
[DEBUG] base = 19
[DEBUG] left array: [37 32 44 31 40]
[DEBUG] right array: [12 1 18]

[DEBUG] current array: [37 32 44 31 40]
[DEBUG] base = 37
[DEBUG] left array: [44 40]
[DEBUG] right array: [32 31]

[DEBUG] current array: [44 40]
[DEBUG] base = 44
[DEBUG] left array: []
[DEBUG] right array: [40]

[DEBUG] current array: []

[DEBUG] current array: [40]

[DEBUG] current array: [32 31]
[DEBUG] base = 32
[DEBUG] left array: []
[DEBUG] right array: [31]

[DEBUG] current array: []

[DEBUG] current array: [31]

[DEBUG] current array: [12 1 18]
[DEBUG] base = 12
[DEBUG] left array: [18]
[DEBUG] right array: [1]

[DEBUG] current array: [18]

[DEBUG] current array: [1]
Result:  [44 40 37 32 31 19 18 12 1]
```

