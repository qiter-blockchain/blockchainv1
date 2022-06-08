package main

import (
	"crypto/sha256"
	"fmt"
)

// 定义结构
// 前区块哈希
// 当前区块哈希
// 数据

// 创建区块
// 生成哈希
// 引入区块链
// 添加区块
// 重构代码
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Block struct {
	PrevBlockHash []byte //前一个hash
	Hash          []byte //当前hash
	Data          []byte //数据
}

// 创建区块 对Block的每一个字段填充数据
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	block.SetHash()
	return &block
}

//为了生成区块哈希，我们实现一个简单的函数，来计算哈希值，没有随机值，没有难度值
func (block *Block) SetHash() {
	var data []byte
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.Data...)
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

//创建区块链，使用Block 数组模拟
type BlockChain struct {
	Blocks []*Block
}

//实现创建区块链的方法
func NewBlockChain() *BlockChain {
	//创建的时候添加一个区块： 创世块
	genesisBlock := NewBlock(genesisInfo, []byte{0x0000000000000})
	bc := BlockChain{Blocks: []*Block{genesisBlock}}
	return &bc
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 1.创建一个区块
	//bc.Blocks的最后一个区块的Hash值就是当前新区块的PrevBlockHash
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := lastBlock.Hash

	block := NewBlock(data, prevHash)
	// 2.添加到bc.Blocks数组中
	bc.Blocks = append(bc.Blocks, block)
}

func main() {
	// fmt.Printf("helloworld\n")
	// block := NewBlock(genesisInfo, []byte{0x0000000000000})

	bc := NewBlockChain()
	bc.AddBlock("what's?")
	for i, block := range bc.Blocks {
		fmt.Printf("**************%d******************\n", i)
		fmt.Printf("PrevBlockHash = %x\n", block.PrevBlockHash)
		fmt.Printf("Hash = %x\n", block.Hash)
		fmt.Printf("data = %s\n", block.Data)
	}

}
