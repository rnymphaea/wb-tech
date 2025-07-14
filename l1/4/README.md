# L1.4 Завершение по Ctrl+C
## Задание
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
## Пример работы программы
#### Справка
```bash
$ go run graceful_shutdown.go --help
```
```
Usage of /path/to/exe/graceful_shutdown:
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
go run graceful_shutdown.go 
```
```
0
1
2
3
^C
```

#### Режим отображения логов
```bash 
$ go run graceful_shutdown.go --debug --interval=500ms --sizedata=10 --workers=5
```
```
Debug mode started
Amount of workers: 5
Amount of messages to send: 10
Interval beetween messages: 500ms
[DEBUG] Main goroutine sended to chan: 0
[DEBUG] worker 4 received from chan: 0
[DEBUG] Main goroutine sended to chan: 1
[DEBUG] worker 4 received from chan: 1
[DEBUG] Main goroutine sended to chan: 2
[DEBUG] worker 0 received from chan: 2
[DEBUG] Main goroutine sended to chan: 3
[DEBUG] worker 3 received from chan: 3
^C
Received SIGINT. Shutting down...
[DEBUG] worker 1 received shutdown signal
[DEBUG] worker 4 received shutdown signal
[DEBUG] worker 2 received shutdown signal
[DEBUG] worker 3 received shutdown signal
[DEBUG] worker 0 received shutdown signal
[DEBUG] Main goroutine received shutdown signal
```
