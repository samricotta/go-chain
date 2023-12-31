package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	blockchain "github.com/samricotta/go-chain/blockchain/v1"
)

const MAX_BLOCK_SIZE = 1e6

type Blockchain struct {
	Blocks []*blockchain.Block
}

func calculateBlockHash(block *blockchain.Block) string {
	// Concatenate block data into a single string. - shorthand for "integer to ASCII"
	record := strconv.Itoa(int(block.Index)) + block.Timestamp + string(block.Data) + block.PreviousHash
	// Inititalise new SHA256 hashing object.
	h := sha256.New()
	// Write the record string to the hashing object.
	h.Write([]byte(record))
	// get the hashed bytes and return as a string - append the hash to an empty slice of bytes.
	// nil is used when you want to create an empty slice of bytes.
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(data []byte, previousHash string, lastIndex int32) *blockchain.Block {
	block := &blockchain.Block{
		Index:        lastIndex + 1,
		Timestamp:    time.Now().String(),
		Data:         data,     
		PreviousHash: previousHash,
	}
	block.Hash = calculateBlockHash(block)
	return block
}

func (bc *Blockchain) AddBlock(data []byte) {
	//get previous block from the length of block
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	//create new block with data, previous hash and index
	newBlock := NewBlock(data, previousBlock.Hash, previousBlock.Index)
	//append new block to the blockchain
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) Genesis() {
	//create genesis block with data, previous hash and index
	bc.Blocks = append(bc.Blocks, NewBlock([]byte("Genesis Block"), "", 0))
}

func (bc *Blockchain) GetBlockByIndex(index int32) *blockchain.Block {
	//iterate over the blocks
	for _, block := range bc.Blocks {
		if block.Index == index {
			return block
		}
	}
	return nil
}

func(bc *Blockchain) CalculateBlockSize() int32 {
	blockSize := 0
	for _, block := range bc.Blocks {
		 blockSize += len(block.Data)
	}
	
	return int32(blockSize)
}