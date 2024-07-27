package main

import (
	"fmt"
	"time"
)

func SelectTraining(c, quit chan int) {
	// selectは、複数のチャネルの処理を待ち受ける
	// どのチャネルが準備できるかを待ち受ける
	// どのチャネルも準備できていない場合、ブロックされる
	// どのチャネルも準備できていない場合、default節が実行される
	// どのチャネルも準備できていない場合、ブロックされないようにするために、default節を追加することができる

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
			time.Sleep(50 * time.Millisecond)
		}
		
	}
}