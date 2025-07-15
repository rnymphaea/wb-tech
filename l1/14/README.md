
# L1.14 Определение типа переменной в runtime
## Задание
Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).
## Пример работы программы
```bash
```
```bash
$ go run exhange.go
```
```
Trying int
The type of [123] is int
Trying string
The type of [simple] is string
Trying bool
The type of [true] is bool
Trying chan
The type of [0xc000072070] is chan
Trying other type
The type of [97] is undefined
```

