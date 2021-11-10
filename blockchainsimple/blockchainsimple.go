package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}

	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockCHain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockCHain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PreHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)

	}

}
