package main

import "fmt"

func PrintSlicePow(s []string) {
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}