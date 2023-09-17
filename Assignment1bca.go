package Assignment1bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	transaction  string
	nonce        int
	hash         string
	previousHash string
}

type BlockList struct {
	blocks []*Block
}

func (b *BlockList) NewBlock(transaction string, nonce int, previousHash string) *Block {
	bl := new(Block)
	bl.transaction = transaction
	bl.nonce = nonce
	bl.previousHash = previousHash

	Str := strconv.Itoa(bl.nonce) + bl.previousHash + bl.transaction
	bl.hash = CalculateHash(Str)

	b.blocks = append(b.blocks, bl)

	return bl
}

func (b *BlockList) ListBlocks() {
	counter := 1
	for _, i := range b.blocks {
		fmt.Printf("%s Transaction %d %s \n", strings.Repeat("=", 35), counter, strings.Repeat("=", 35))
		fmt.Println("Transaction:   ", i.transaction)
		fmt.Println("Nonce:         ", i.nonce)
		fmt.Println("Previous Hash: ", i.previousHash)
		fmt.Println()
		counter++
	}
}

func (b *BlockList) ChangeBlock(count int, transaction string) *Block {
	count = count - 1
	b.blocks[count].transaction = transaction
	Stri := strconv.Itoa(b.blocks[count].nonce) + b.blocks[count].previousHash + b.blocks[count].transaction
	b.blocks[count].hash = CalculateHash(Stri)
	modBlock := b.blocks[count]

	return modBlock
}

func (b *BlockList) VerifyChain() int {
	count := 0
	flag := 0

	for count = 0; count < (len(b.blocks) - 1); count++ {
		if b.blocks[count].hash != b.blocks[count+1].previousHash {
			flag = 1
			fmt.Println("Changes detected in Block", count+1)
		}
	}

	return flag
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	Hash := fmt.Sprintf("%x", hash)
	return Hash
}
