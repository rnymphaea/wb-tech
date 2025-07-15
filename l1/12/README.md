# L1.12 Собственное множество строк
## Задание
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
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
$ go run unique_words.go --default
```
```
Result:  [cat dog tree]
```

#### Режим отображения логов
```bash 
$ go run unique_words.go --debug
```
```
Debug mode started
Enter number of words: 5
Enter 5 words: Cat horse DOG cat dog
Ignore case? [y/n] y
[DEBUG] initial array: [Cat horse DOG cat dog]
[DEBUG] ingore case: true
[DEBUG] map after iterating over array: map[cat:{} dog:{} horse:{}]
Result:  [cat horse dog]
```
## Тестирование
```bash 
$ go test -v
```
