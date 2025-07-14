# L1.5 Таймаут на канал
## Задание
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
## Пример работы программы
```
├── context.go // с использованием контекста
└── time.go    // с использованием таймера
```
Вывод программ не отличается. Поэтому для демонстрации работы будет использован файл context.go

#### Справка
```bash
$ go run context.go --help
```
```
Usage of /path/to/exe/context:
  -debug
    	show logs to check the correctness of execution
  -interval duration
    	interval between messages (seconds) (default 1s)
  -timeout int
    	number of seconds to finish the program (default 5)
```

#### Стандартный вывод 
```bash
$ go run context.go 
```
```
31
29
82
45
87
```

#### Режим отображения логов
```bash 
$ go run context.go --debug --interval=500ms --timeout=3
```
```
Debug mode started
Timeout: 3s
Interval beetween messages: 500ms
[DEBUG] 0ms Sender send: 51
[DEBUG] 0ms Receiver got: 51
[DEBUG] 500ms Sender send: 48
[DEBUG] 500ms Receiver got: 48
[DEBUG] 1001ms Sender send: 10
[DEBUG] 1001ms Receiver got: 10
[DEBUG] 1502ms Sender send: 60
[DEBUG] 1502ms Receiver got: 60
[DEBUG] 2003ms Sender send: 90
[DEBUG] 2003ms Receiver got: 90
[DEBUG] 2504ms Sender send: 81
[DEBUG] 2504ms Receiver got: 81
[DEBUG] 3000ms Receiver: time exceeded
[DEBUG] 3005ms Sender: time exceeded
```
