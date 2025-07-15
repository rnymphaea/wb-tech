# L1.26 Уникальные символы в строке
## Задание
Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
## Пример работы программы
#### Справка
```bash
$ go run unique_symbols.go --help
```
```
Usage of /path/to/exe/unique_symbols:
  -debug
    	show logs to check the correctness of execution
```

#### Стандартный вывод 
```bash
$ go run unique_symbols.go
```
```
Enter string: alskdjfA        
String 'alskdjfA' does not contain unique symbols
```

#### Режим отображения логов
```bash 
$ go run unique_symbols.go --debug
```
```
Debug mode started
Enter string: alskdjfA
[DEBUG] excess symbols: [a]
String 'alskdjfA' does not contain unique symbols
```
## Тестирование
```bash 
$ go test -v
```
