package main

import (
	"fmt"
	"time"
)

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

		timeFormat := time.Unix(int64(block.TimeStamp),0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳:%s\n",timeFormat)

		fmt.Printf("难度值:%d\n",block.Bits)
		fmt.Printf("随机数:%x\n",block.Nonce)

		pow := NewProofOfWork(*block)
		fmt.Printf("IsVaild:%v\n",pow.IsVaild())

		fmt.Printf("区块数据:%s\n",block.Data)
	}
}