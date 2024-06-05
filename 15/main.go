package main

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

// Проблема изначального кода состоит в том, что срез представляет собой ссылочный тип данных.
// Получается, что при выполнении justString = v[:100] мы сохраним ссылку на срез строки.
// Это приведет к тому, что значение переменной v не будет очищенно GC, так как на v будет ссылаться justString,
// и мы будем занимать лишнюю память.

// Решением является создание копии среза и сохранение ее в переменную justString
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}

func main() {
	someFunc()
}
