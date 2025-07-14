# L1.3 Работа нескольких воркеров
## Задание
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
## Пример работы программы
#### Справка
```bash
$ go run workers.go --help
```
```
Usage of /path/to/exe/workers:
  -debug
    	show logs to check the correctness of execution
  -interval duration
    	interval between messages (seconds) (default 1s)
  -sizedata int
    	amount of messages to send to chan (default 10)
  -workers int
    	amount of workers (default 10)
```

#### Стандартный вывод 
```bash
go run workers.go 
```
```
0
1
2
3
4
5
6
7
8
9
```

#### Режим отображения логов
```bash 
$ go run workers.go --debug --interval=500ms --sizedata=15 --workers=5
```
```
Debug mode started
Amount of workers: 5
Amount of messages to send: 15
Interval beetween messages: 500ms
[DEBUG] 0ms main goroutine sended to chan: 0
[DEBUG] 0ms worker 0 received from chan: 0
[DEBUG] 500ms main goroutine sended to chan: 1
[DEBUG] 500ms worker 1 received from chan: 1
[DEBUG] 1001ms main goroutine sended to chan: 2
[DEBUG] 1001ms worker 3 received from chan: 2
[DEBUG] 1502ms main goroutine sended to chan: 3
[DEBUG] 1502ms worker 4 received from chan: 3
[DEBUG] 2002ms main goroutine sended to chan: 4
[DEBUG] 2002ms worker 2 received from chan: 4
[DEBUG] 2503ms main goroutine sended to chan: 5
[DEBUG] 2503ms worker 0 received from chan: 5
[DEBUG] 3004ms main goroutine sended to chan: 6
[DEBUG] 3004ms worker 1 received from chan: 6
[DEBUG] 3505ms main goroutine sended to chan: 7
[DEBUG] 3505ms worker 3 received from chan: 7
[DEBUG] 4006ms main goroutine sended to chan: 8
[DEBUG] 4006ms worker 4 received from chan: 8
[DEBUG] 4507ms main goroutine sended to chan: 9
[DEBUG] 4507ms worker 2 received from chan: 9
[DEBUG] 5008ms main goroutine sended to chan: 10
[DEBUG] 5008ms worker 0 received from chan: 10
[DEBUG] 5508ms main goroutine sended to chan: 11
[DEBUG] 5508ms worker 1 received from chan: 11
[DEBUG] 6009ms main goroutine sended to chan: 12
[DEBUG] 6009ms worker 3 received from chan: 12
[DEBUG] 6510ms main goroutine sended to chan: 13
[DEBUG] 6510ms worker 4 received from chan: 13
[DEBUG] 7011ms main goroutine sended to chan: 14
[DEBUG] 7011ms worker 2 received from chan: 14
```
