package main

import(
	"fmt"
)

func main() {
	//~ block := NewBlock(1, nil, []byte("the first block testing"))
	//~ fmt.Printf("The first block : %v\n", block)
	bc := CreateGenesisBlockWithGeneisBlock()
	//~ fmt.Printf("blockchain:%v\n", bc.Blocks[0])
	
	// ddBlock(height int63, prevBlockHash []byte, data []byte){
	height := bc.Blocks[len(bc.Blocks)-1].Height+1
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	data := []byte("alice send 100 btc to bob")
	
	bc.AddBlock(height, prevHash, data)
	
	height = bc.Blocks[len(bc.Blocks)-1].Height+1
	prevHash = bc.Blocks[len(bc.Blocks)-1].Hash
	data = []byte("bob send 10 btc to alice")
	
	bc.AddBlock(height, prevHash, data)
	
	for _, block := range bc.Blocks {
		fmt.Printf("prevHash: %x , Hash : %x\n", block.PrevHash, block.Hash)
	}
}

