package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// Вывод:
// error
//
// Объяснение:
// Интерфейсы включают в себя два поля: информацию о типе и конкретном зачении этого типа (указатель на тип и указатель на значение).
// В данном случае переменная err объявляется как интерфейс error (в данный момент это nil интерфейс, так как нет ни типа, ни значения).
// Затем ей присваивается значение функции test - nil, но тип *customError, а следовательно, это уже не nil интерфейс, так как есть информация о типе.
// Далее уже следует сравнение с nil и мы получаем строку "error" по указанным выше причинам.
