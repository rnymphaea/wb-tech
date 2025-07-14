# L1.10 Группировка температур
## Задание
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить эти значения в группы с шагом 10 градусов.

Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
## Пример работы программы
#### Справка
```bash
$ go run pipeline.go --help
```
```
Usage of /path/to/exe/temperatures:
  -debug
    	show logs to check the correctness of execution
  -default
    	use default array of measurements (-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)
  -sizedata int
    	amount of measurements (default 10)
```

#### Стандартный вывод 
```bash
$ go run temperatures.go --default
```
```
map[-20:[-25.4 -27 -21] 10:[13 19 15.5] 20:[24.5] 30:[32.5]]
```

#### Режим отображения логов
```bash 
$ go run temperatures.go --debug --sizedata=15
```
```
Debug mode started
Amount of measurements: 15
[DEBUG] array of measurements: [-48.8 -79.4 -82.3 55.9 -17.9 -54.3 34.4 -75 -0.7 11.3 -67.4 18.9 -14.2 28.5 15.6]
map[-80:[-82.3] -70:[-79.4 -75] -60:[-67.4] -50:[-54.3] -40:[-48.8] -10:[-17.9 -14.2] 0:[-0.7] 10:[11.3 18.9 15.6] 20:[28.5] 30:[34.4] 50:[55.9]]
```

