# L1.13 Обмен значениями без третьей переменной
## Задание
Поменять местами два числа без использования временной переменной.
## Пример работы программы
#### Справка
```bash
$ go run unique_words.go --help
```
```
Usage of /path/to/exe/unique_words:
  -debug
    	show logs to check the correctness of execution
  -default
    	use default array of strings ('cat', 'cat', 'dog', 'cat', 'tree')

```

#### Стандартный вывод 
```bash
$ go run exhange.go
```
```
Enter first number: 5
Enter second number: 10
num1 = 5, num2 = 10
First way (arithmetic operations):
num1 = 10, num = 5
Second way (logical operations):
num1 = 5, num = 10
```

#### Режим отображения логов
```bash 
$ go run exhange.go --debug
```
```
Debug mode started
Enter first number: 5
Enter second number: 10
num1 = 5, num2 = 10
First way (arithmetic operations):
[DEBUG] changing values by using addition and subtraction
[DEBUG] step 1: assign to num2 sum of num1 and num2: num1 = 5, num2 = 15
[DEBUG] step 2: assign to num1 sub of num1 and num2 (num1 + num2): num1 = 10, num2 = 15
[DEBUG] step 3: assign to num2 sub of num1 (num2) and num2 (num1 + num2): num1 = 10, num2 = 5
num1 = 10, num = 5
Second way (logical operations):
[DEBUG] changing values by using XOR
[DEBUG] step 1: assign to num2 XOR of num1 and num2: num1 = 1010 (10), num2 = 1111 (15)
[DEBUG] step 2: assign to num1 XOR of num1 and num2 (num1 XOR num2): num1 = 101 (5), num2 = 1111 (15)
[DEBUG] step 3: assign to num2 XOR of num1 (num2) and num2 (num1 XOR num2): num1 = 101 (5), num2 = 1010 (10)
num1 = 5, num = 10
```

