# L1.6 Остановка горутины
## Задание
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
## Пример работы программы


#### Справка
```bash
$ go run stopping_goroutine.go --help
```
```
Usage of /path/to/exe/stopping_goroutine:
  -debug
    	show logs to check the correctness of execution
```

#### Стандартный вывод 
```bash
$ go run stopping_goroutine.go
```
```
func ConditionalExit: goroutine is working
func ConditionalExit: goroutine is working
func ConditionalExit: goroutine is working
func ConditionalExit: goroutine is working
func ConditionalExit: goroutine is working
func ConditionalExit: the condition stops being fulfilled. Quitting...
func ConditionalExit: goroutine stops working
func ChannelExit: goroutine is working
func ChannelExit: goroutine is working
func ChannelExit: goroutine is working
func ChannelExit: goroutine is working
func ChannelExit: goroutine is working
func ChannelExit: sending stop signal to channel. Quitting...
func ChannelExit: goroutine stops working
func ContextWithTimeoutExit: goroutine is working
func ContextWithTimeoutExit: goroutine is working
func ContextWithTimeoutExit: goroutine is working
func ContextWithTimeoutExit: goroutine is working
func ContextWithTimeoutExit: goroutine is working
func ContextWithTimeoutExit: goroutine stops working
func ContextWithCancelExit: goroutine is working
func ContextWithCancelExit: goroutine is working
func ContextWithCancelExit: goroutine is working
func ContextWithCancelExit: goroutine is working
func ContextWithCancelExit: goroutine is working
func ContextWithCancelExit: calling cancel(). Quitting...
func ContextWithCancelExit: goroutine stops working...
func RuntimeGoExit: entering to goroutine
func RuntimeGoExit: quitting goroutine after sleep
func ClosedChannelExit: receiver: channel is not closed
func ClosedChannelExit: receiver: channel is not closed
func ClosedChannelExit: receiver: channel is not closed
func ClosedChannelExit: receiver: channel is not closed
func ClosedChannelExit: receiver: channel is not closed
func ClosedChannelExit: receiver discovered that channel is closed. Quitting...
func RuntimeGoExit: entering to goroutine
func RuntimeGoExit: goroutine recovered from panic: stop goroutine
func TimeAfterExit: goroutine is working
func TimeAfterExit: goroutine is working
func TimeAfterExit: goroutine is working
func TimeAfterExit: goroutine is working
func TimeAfterExit: goroutine is working
func TimeAfterExit: goroutine stops working
func OsExit: entering to goroutine
func OsExit: goroutine is calling os.Exit
```

#### Режим отображения логов
```bash 
$ go run stopping_goroutine.go --debug
```
```
Debug mode started

Demonstrating of exiting a goroutine on condition
Interval between messages of goroutine: 200ms
Timeout: 1s
[DEBUG] func ConditionalExit: [0ms] goroutine is working
[DEBUG] func ConditionalExit: [200ms] goroutine is working
[DEBUG] func ConditionalExit: [401ms] goroutine is working
[DEBUG] func ConditionalExit: [601ms] goroutine is working
[DEBUG] func ConditionalExit: [802ms] goroutine is working
[DEBUG] func ConditionalExit: [1000ms] the condition stops being fulfilled
[DEBUG] func ConditionalExit: [1002ms] goroutine: condition is no longer fulfilled

Demonstrating of exiting a goroutine on channel notification
Interval between messages of goroutine: 200ms
Timeout: 1s
[DEBUG] func ChannelExit: [0ms] goroutine is working
[DEBUG] func ChannelExit: [200ms] goroutine is working
[DEBUG] func ChannelExit: [400ms] goroutine is working
[DEBUG] func ChannelExit: [601ms] goroutine is working
[DEBUG] func ChannelExit: [801ms] goroutine is working
[DEBUG] func ChannelExit: [1000ms] sending stop signal to channel
[DEBUG] func ChannelExit: [1002ms] goroutine: received stop signal from channel

Demonstrating of exiting a goroutine with context.WithTimeout
Interval between messages of goroutine: 200ms
Timeout: 1s
[DEBUG] func ContextWithTimeoutExit: [0ms] sleeping for 1000...
[DEBUG] func ContextWithTimeoutExit: [0ms] goroutine is working
[DEBUG] func ContextWithTimeoutExit: [200ms] goroutine is working
[DEBUG] func ContextWithTimeoutExit: [401ms] goroutine is working
[DEBUG] func ContextWithTimeoutExit: [601ms] goroutine is working
[DEBUG] func ContextWithTimeoutExit: [802ms] goroutine is working
[DEBUG] func ContextWithTimeoutExit: [1008ms] goroutine: time exceeded

Demonstrating of exiting a goroutine with context.WithCancel
Interval between messages of goroutine: 200ms
Timeout: 1s
[DEBUG] func ContextWithCancelExit: [0ms] goroutine is working
[DEBUG] func ContextWithCancelExit: [200ms] goroutine is working
[DEBUG] func ContextWithCancelExit: [400ms] goroutine is working
[DEBUG] func ContextWithCancelExit: [601ms] goroutine is working
[DEBUG] func ContextWithCancelExit: [801ms] goroutine is working
[DEBUG] func ContextWithCancelExit: [1001ms] calling cancel()
[DEBUG] func ContextWithCancelExit: [1002ms] goroutine: received cancel signal

Demonstrating of exiting a goroutine with runtime.Goexit
Timeout: 1s
func RuntimeGoExit: [0ms] waiting for goroutine
[DEBUG] func RuntimeGoExit: [0ms] entering to goroutine and sleeping for 1s...
[DEBUG] func RuntimeGoExit: [1000ms] goroutine: quitting after sleep

Demonstrating of exiting a goroutine by closing a channel
Interval between messages sending to channel: 200ms
[DEBUG] func ClosedChannelExit: [0ms] sender: sending 0
[DEBUG] func ClosedChannelExit: [0ms] receiver got: 0
[DEBUG] func ClosedChannelExit: [200ms] sender: sending 1
[DEBUG] func ClosedChannelExit: [200ms] receiver got: 1
[DEBUG] func ClosedChannelExit: [400ms] sender: sending 2
[DEBUG] func ClosedChannelExit: [400ms] receiver got: 2
[DEBUG] func ClosedChannelExit: [601ms] sender: sending 3
[DEBUG] func ClosedChannelExit: [601ms] receiver got: 3
[DEBUG] func ClosedChannelExit: [801ms] sender: sending 4
[DEBUG] func ClosedChannelExit: [801ms] receiver got: 4
[DEBUG] func ClosedChannelExit: [1002ms] sender: closing the channel
[DEBUG] func ClosedChannelExit: [1002ms] receiver: channel is closed

Demonstrating of exiting a goroutine by panic
Timeout: 1s
[DEBUG] func RuntimeGoExit: [0ms] entering to goroutine
[DEBUG] func RuntimeGoExit: [1000ms] goroutine: calling panic...
[DEBUG] func RuntimeGoExit: [1000ms] goroutine recovered from panic: stop goroutine

Demonstrating of exiting a goroutine with time.After
Interval between messages of goroutine: 200ms
Timeout: 1s
[DEBUG] func TimeAfterExit: [0ms] sleeping for 1000...
[DEBUG] func TimeAfterExit: [0ms] goroutine is working
[DEBUG] func TimeAfterExit: [200ms] goroutine is working
[DEBUG] func TimeAfterExit: [400ms] goroutine is working
[DEBUG] func TimeAfterExit: [601ms] goroutine is working
[DEBUG] func TimeAfterExit: [801ms] goroutine is working
[DEBUG] func TimeAfterExit: [1002ms] goroutine: time exceeded

Demonstrating of exiting a goroutine with os.Exit
Timeout: 1s
[DEBUG] func OsExit: [0ms] entering to goroutine and sleeping for 1s
[DEBUG] func OsExit: [1000ms] goroutine: calling os.Exit...
```
