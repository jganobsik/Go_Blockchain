package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
type BlockChain struct {
	blocks []*Block
}

func main() {
	//create new blockchain
	chain := InitBlockChain()
	//add blocks via chain.AddBlock("message")
	//example block creation
	chain.AddBlock("First New Block In Chain")
}

//hash block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//create new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	//return value is new block
	return block
}

func (c *BlockChain) AddBlock(data string) {
	//get last block in chain
	prevBlock := c.blocks[len(c.blocks)-1]
	//create new block
	new := CreateBlock(data, prevBlock.Hash)
	//add new block to chain
	c.blocks = append(c.blocks, new)
}

// create the first block
func Genesis() *Block {
	//call function creating first block
	return CreateBlock("genesis", []byte{})
}

// initialize blockchain
func InitBlockChain() *BlockChain {
	//return reference to new blockchain with genesis block
	return &BlockChain{[]*Block{Genesis()}}
}

func (c *BlockChain) PrintBlockData() {
	for _, block := range c.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
