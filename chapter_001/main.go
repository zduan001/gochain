package main

import(
	"fmt"
)

func main() {
	block := NewBlock(1, nil, []byte("the first block testing"))
	fmt.Printf("The first block : %v\n", block)
}

