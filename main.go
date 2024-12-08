package main

import (
	"Go-Blockchain/Blockchain"
	"fmt"
)

func main() {
	bc := Blockchain.NewBlockchain()

	bc.AddBlock("Roman Send 50 BTC to Lara")
	bc.AddBlock("Eduardo Send 5 BTC to Ruth")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
