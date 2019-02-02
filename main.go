package main

import "fmt"

func main() {

	bc := NewBlockChain()

	bc.AddBlock("Hello,航头!")
	bc.AddBlock("再见，航头!")

	for i,block := range bc.Blocks {
		fmt.Printf("======区块链高度:%d======\n",i)

		fmt.Printf("前区块哈希值:%x\n",block.PrevBlockHash)
		fmt.Printf("当前区块哈希值:%x\n",block.CurrentBlockHash)
		fmt.Printf("区块链数据:%s\n",block.Data)
		fmt.Printf("区块链数据(十六):%x\n",block.Data)
	}
}