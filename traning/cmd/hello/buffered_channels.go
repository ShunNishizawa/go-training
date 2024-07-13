// バッファの詰まるチャネル

package main

import "fmt"

func BufferedChannels() {
	// チャネルを作成する
	ch := make(chan int, 2)

	// バッファが詰まることはない
	ch <- 1
	ch <- 2

	// バッファが詰まる
	// ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}