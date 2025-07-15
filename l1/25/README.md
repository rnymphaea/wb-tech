# L1.25 Своя функция Sleep
## Задание
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.
## Пример работы программы
#### Справка
```bash
$ go run sleep.go --help
```
```
Usage of /path/to/exe/sleep:
  -debug
    	show logs to check the correctness of execution
```

#### Стандартный вывод 
```bash
$ go run sleep.go
```
```
Enter number of milliseconds to sleep: 1000
Time exceeded
```

#### Режим отображения логов
```bash 
$ go run sleep.go --debug
```
```
Debug mode started
Enter number of milliseconds to sleep: 1000
[DEBUG] 1000ms sleep: time exceeded
Time exceeded
```
