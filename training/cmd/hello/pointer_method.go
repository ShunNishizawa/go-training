// ポインタメソッド
package main

type Pow struct {
	name string
}

func PointerMethod(p *Pow) {
	p.name = p.name + "隊員"
}