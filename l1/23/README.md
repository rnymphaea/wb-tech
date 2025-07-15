# L1.23 Удаление элемента слайса
## Задание
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
## Пример работы программы
#### Справка
```bash
$ go run delete.go --help
```
```
Usage of /path/to/exe/delete:
  -debug
    	show logs to check the correctness of execution
  -random
    	use the random array to check the correctness
```

#### Стандартный вывод 
```bash
$ go run delete.go
```
```
Enter the size of array: 5
Enter 5 elements of array: 2 3 52 1 2
Enter the target index to delete: 1
Initial array:  [2 3 52 1 2]
Index to delete:  1
Result:  [2 52 1 2]
```

#### Режим отображения логов
```bash 
$ go run delete.go --debug --random
```
```
Debug mode started
Random: true
[DEBUG] initial array: [36 45 34 19 24 12]
[DEBUG] index to delete: 0, arr[0] = 36
Result:  [45 34 19 24 12]
```
