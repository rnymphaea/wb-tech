# L1.24 Расстояние между точками
## Задание
Разработать программу нахождения расстояния между двумя точками на плоскости. Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.
## Пример работы программы
#### Справка
```bash
$ go run distance.go --help
```
```
Usage of /path/to/exe/distance:
  -debug
    	show logs to check the correctness of execution
  -random
    	use the random array to check the correctness
```

#### Стандартный вывод 
```bash
$ go run distance.go
```
```
Enter (x y) of point 1: 0 0
Enter (x y) of point 2: 1 1
Distance between points:  1.4142135623730951
```

#### Режим отображения логов
```bash 
$ go run distance.go --debug --random
```
```
Debug mode started
Random: true
[DEBUG] point 1: (70.943764, 10.495848)
[DEBUG] point 2: (34.276538, -51.646016)
Distance between points:  72.15328740143389
```
