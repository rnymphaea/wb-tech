# L1.11 Пересечение множеств
## Задание
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:\
A = {1,2,3}\
B = {2,3,4}\
Пересечение = {2,3}
## Пример работы программы
#### Справка
```bash
$ go run intersection.go --help
```
```
Usage of /path/to/exe/intersection:
  -debug
    	show logs to check the correctness of execution
```

#### Стандартный вывод 
```bash
$ go run intersection.go
```
```
Set 1: [16 5 4 17 12]
Set 2: [17 15 2]
Intersection of these sets:  [17]
```

#### Режим отображения логов
```bash 
$ go run intersection.go --debug
```
```
Debug mode started
[DEBUG] map after initializing: map[3:{} 4:{} 8:{} 13:{} 15:{} 16:{} 17:{} 18:{}]
[DEBUG] map after initializing: map[1:{} 2:{} 4:{} 11:{} 13:{} 15:{}]
Set 1: [3 17 18 13 4 16 8 15]
Set 2: [2 1 4 13 15 11]
[DEBUG] map after writing elements of set 1: map[3:{} 4:{} 8:{} 13:{} 15:{} 16:{} 17:{} 18:{}]
[DEBUG] checking set 2
[DEBUG] check: element 2 is in set 1: false
[DEBUG] check: element 1 is in set 1: false
[DEBUG] check: element 4 is in set 1: true
[DEBUG] check: element 13 is in set 1: true
[DEBUG] check: element 15 is in set 1: true
[DEBUG] check: element 11 is in set 1: false
Intersection of these sets:  [4 13 15]
```

