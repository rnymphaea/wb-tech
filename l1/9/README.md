# L1.9 Конвейер чисел
## Задание
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.
## Пример работы программы
#### Справка
```bash
$ go run pipeline.go --help
```
```
Usage of /path/to/exe/pipeline:
  -debug
    	show logs to check the correctness of execution
  -interval duration
    	interval between messages (seconds) (default 200ms)
  -sizedata int
    	amount of messages to send (default 10)
```

#### Стандартный вывод 
```bash
$ go run pipeline.go
```
```
116
120
90
92
98
110
196
122
64
164
```

#### Режим отображения логов
```bash 
$ go run pipeline.go --debug --interval=300ms --sizedata=7
```
```
Debug mode started
Amount of messages to send: 7
Interval beetween messages: 300ms
[DEBUG] 0ms sender: sending 92 to ch1
[DEBUG] 0ms processor: received 92 from ch1
[DEBUG] 0ms processor: sending 184 to ch2
[DEBUG] 0ms receiver: received 184 from ch2
[DEBUG] 300ms sender: sending 93 to ch1
[DEBUG] 300ms processor: received 93 from ch1
[DEBUG] 300ms processor: sending 186 to ch2
[DEBUG] 300ms receiver: received 186 from ch2
[DEBUG] 601ms sender: sending 84 to ch1
[DEBUG] 601ms processor: received 84 from ch1
[DEBUG] 601ms processor: sending 168 to ch2
[DEBUG] 601ms receiver: received 168 from ch2
[DEBUG] 901ms sender: sending 89 to ch1
[DEBUG] 902ms processor: received 89 from ch1
[DEBUG] 902ms processor: sending 178 to ch2
[DEBUG] 902ms receiver: received 178 from ch2
[DEBUG] 1202ms sender: sending 65 to ch1
[DEBUG] 1202ms processor: received 65 from ch1
[DEBUG] 1202ms processor: sending 130 to ch2
[DEBUG] 1202ms receiver: received 130 from ch2
[DEBUG] 1503ms sender: sending 50 to ch1
[DEBUG] 1503ms processor: received 50 from ch1
[DEBUG] 1503ms processor: sending 100 to ch2
[DEBUG] 1503ms receiver: received 100 from ch2
[DEBUG] 1803ms sender: sending 85 to ch1
[DEBUG] 1803ms processor: received 85 from ch1
[DEBUG] 1803ms processor: sending 170 to ch2
[DEBUG] 1803ms receiver: received 170 from ch2
[DEBUG] 2104ms sender: closing ch1
[DEBUG] 2104ms processor: ch1 is closed. Closing ch2
[DEBUG] 2104ms receiver: ch2 is closed. Quitting...
```
