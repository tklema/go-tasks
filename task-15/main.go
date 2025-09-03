package main

import (
	"fmt"
	"strings"
)

/*
	Проблема изначального кода в том, что срез строки не создает новую строку,
	а ссылается на ее память. Соответственно, мы больше не используем v, но
	при этом ее память используется в программе (не вся), и получается
	утечка памяти.
	В итоге justString ссылается на память v, и из за этого память v не может
	очиститься, хотя и не используется в программе

	Решение: используем специальный метод, чтобы склонировать часть строки
*/

var justString string

func createHugeString(length int) string {
	return strings.Repeat("a", length)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:10])
}

func main() {
	fmt.Printf("\"%s\"\n", justString)
	someFunc()
	fmt.Printf("\"%s\"", justString)
}
