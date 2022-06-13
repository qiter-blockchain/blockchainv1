package main

import (
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

func main() {
	// fmt.Printf("helloworld\n")
	// block := NewBlock(genesisInfo, []byte{0x0000000000000})

	bc := NewBlockChain()
	bc.AddBlock("班主任来了，大家欢迎~")
	bc.AddBlock("班主任走了")

	for i, block := range bc.Blocks {
		fmt.Printf("**************%d******************\n", i)
		fmt.Printf("PrevBlockHash = %x\n", block.PrevBlockHash)
		fmt.Printf("Hash = %x\n", block.Hash)
		fmt.Printf("data = %s\n", block.Data)
	}

}
