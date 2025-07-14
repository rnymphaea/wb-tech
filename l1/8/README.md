# L1.8 Установка бита в числе
## Задание
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
## Пример работы программы
#### Справка
```bash
$ go run setting_bit.go --help
```
```
Usage of /path/to/exe/setting_bit:
  -debug
    	show logs to check the correctness of execution
```

#### Стандартный вывод 
```bash
$ go run setting_bit.go
```
```
Enter the number: 50
Enter the number of the bit to be replaced: 3
Replace bit #3 with (1 or 0): 1
Result: 54
```

#### Режим отображения логов
```bash 
$ go run setting_bit.go --debug
```
```
Debug mode started
Enter the number: 50
[DEBUG] bit representation: 110010
Enter the number of the bit to be replaced: 3
Replace bit #3 with (1 or 0): 1
[DEBUG] bit mask: 100
[DEBUG] set zero value to 3 bit: 110010
[DEBUG] replace 3 bit with new value: 110110
Result: 54
[DEBUG] bit representation: 110110
```

## Тестирование
```bash 
$ go test -v
```

