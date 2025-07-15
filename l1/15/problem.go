// var justString string
//
//	func someFunc() {
//	 v := createHugeString(1 &lt;&lt; 10)
//	 justString = v[:100]
//	}
//
//	func main() {
//	 someFunc()
//	}
//
// Проблемы:
// 1. Утечка памяти. createHugeString(1 << 10) скорее всего создаёт большую строку
// При присваивании justString = v[:100] сохраняется ссылка на исходный массив
// Вся большая строка остаётся в памяти, хотя используется только 100 символов
// 2. Потенциальная паника. Если createHugeString вернёт строку короче 100 символов, то срез v[:100] вызовет панику
// 3. Глобальная переменная. Это позволяет изменять её из любого участка пакета.

func someFunc() (string, error) {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		return "", fmt.Errorf("string is too short")
	}
	buf := make([]byte, 0, 100)
	buf = append(buf, v[:100]...)

	v = ""

	return string(buf)
}

func main() {
	justString, err := someFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
}
