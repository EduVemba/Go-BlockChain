package Blockchain

/*
Using the struct Block to create an array of block so it works as a chain of blocks
*/
type Blockchain struct {
	Blocks []*Block
}

/*
param: data = Data that is going to be added in the block
*/
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

/*
  Genesis block the first block being created because we need a prev Block hash.
*/

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
