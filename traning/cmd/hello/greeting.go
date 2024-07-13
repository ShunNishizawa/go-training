package main

// 複数の戻り値を返すことができる
func Greeting(name, message string) (string, string) {
	return name + " says: " + message, name
}