package main

//引入区块链
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockChain() *BlockChain {

	genesisBlock := NewBlock(genesisInfo,[]byte{})

	block := BlockChain{
		Blocks:[]*Block{genesisBlock},
	}

	return &block
}

//添加区块链
func (bc *BlockChain) AddBlock(data string) {

	 //创建一个区块，前区块的哈希值从bc的最后一个区块元素获取即可
	 lastBlock := bc.Blocks[len(bc.Blocks) - 1]

     //即将添加的区块的哈希值，就是bc中的最后区块的前哈希值
     prevHash := lastBlock.CurrentBlockHash

     newBlock := NewBlock(data,prevHash)

     //将新的区块添加到数组中
     bc.Blocks = append(bc.Blocks,newBlock)
}






