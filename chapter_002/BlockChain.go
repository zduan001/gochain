package main

type BlockChain struct {
	Blocks []*Block 
}

// create the first block
func CreateGenesisBlockWithGeneisBlock() *BlockChain {
	block := CreateGenesisBlock([]byte("init blockchain"))
	return &BlockChain{Blocks: []*Block{block}}
}

func(bc *BlockChain) AddBlock(height int64, prevBlockHash []byte, data []byte){
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
} 
