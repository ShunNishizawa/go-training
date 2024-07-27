package main

import (
	"fmt"
	"time"
)

// フィールドの集まりをstructとして定義する
type Vertex struct {
	X, Y float64
}

// structのリテラル

// フィールドの集まりをstructとして定義する
type Dogs struct {
	NAME, COLOR string
}

var (
	v1 = Dogs{"ロッキー", "white"}  // has type Dogs
	v2 = Dogs{NAME: "マーシャル"}  // 犬の名前だけを初期化
	v3 = Dogs{}      // X:0 and Y:0
)

// 関数の戻り値に名前をつけることができる
// bool型の変数は初期値がfalse
var c, python, java bool

	// map
	var m map[string]Dogs

func main() {

	// p, q := 20, 30

	// fmt.Println(Greeting("Gopher", "Hello, World!"))

	// // pointer
	// r := &p // &pはpのアドレスを表す
	// fmt.Printf("アドレスは %v\n", r)// pのアドレスを表示
	// fmt.Println(*r) // *でポインタの指す先の変数を示す
	// *r = 21         // set i through the pointer
	// fmt.Println(p)  // see the new value of i

	// // pointer
	// r = &q
	// fmt.Println(*r) // read i through the pointer


	// var i int
	// c, python, java = true, false, true
	// fmt.Println(i, c, python, java)

	// fmt.Println("struct...\n", Vertex{1, 2})

	// // structのフィールドにアクセスする
	// v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v.X)

	// // structのポインタを使ってアクセスする
	// pv := &v
	// pv.X = 1e9
	// fmt.Println(v)

	// // StructリテラルDogs
	// fmt.Println(v1, p, v2, v3)

	// // 配列
	// arr := []string{"マーシャル", "チェイス", "スカイ", "ロッキー", "ズーマ"}

	// arr = arr[1:4]
	// fmt.Println(arr)

	// // 配列の要素数を取得する
	// fmt.Println(len(arr))

	// arr = arr[:4]
	// fmt.Println(arr)

	// // Nil slices
	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }

	// // make関数でスライスを作成する
	// a := make([]string, 5)
	// PrintSlice("a", a)

	// // スライスへの要素の追加
	// var pow []string
	// PrintSlicePow(pow)

	// pow = append(pow, "マーシャル")
	// PrintSlicePow(pow)

	// // Range
	// var pow2 = []string{"マーシャル", "チェイス", "スカイ", "ロッキー", "ズーマ", "エベレスト"}
	// for i, name := range pow2 {
	// 	fmt.Printf("隊員%d = %s\n", i, name)
	// }
	// // インデックスだけを取得する
	// for i := range pow2 {
	// 	fmt.Println(i)
	// }


	// m = make(map[string]Dogs)
	// m["マーシャル"] = Dogs{"マーシャル", "white"}
	// fmt.Println(m["マーシャル"])
	
	// // Mutating Maps
	// m2 := make(map[string]int)
	// m2["Answer"] = 42
	// fmt.Println("The value:", m2["Answer"])
	// delete(m2, "Answer")
	// fmt.Println("The value:", m2["Answer"])

	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())

	// // Interface
	// var inter I = T{"マーシャル"}
	// inter.M()

	// // empty interface
	// var inter2 interface{}
	// emptyInterface(inter2)

	// inter2 = 42
	// emptyInterface(inter2)

	// inter2 = "hello"
	// emptyInterface(inter2)

	// go Say("こんにちは！マーシャル")
	// Say("こんにちは！チェイス")

	// ChannelTraining()

	// c := make(chan int)
	// quit := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// SelectTraining()

	// DefaultSelect()

	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("hogekey"))
}
