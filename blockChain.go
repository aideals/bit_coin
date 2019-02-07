package main

import (
	"awesomeProject/lib/bolt"
	"fmt"
	//"golang.org/x/tools/blog"
	//"net"
	"os"
)

//引入区块链
type BlockChain struct {
	//Blocks []*Block

	//数据库句柄
	db *bolt.DB
	//存储最后一个区块哈希，尾巴
	tail []byte
}

const blockChainFile = "blockChain.db"
const blockBucket = "blockBucket"
const lastHashKey = "lastHashKey"


//创建区块链
func CreateBlockChain() *BlockChain {

	if isFileExist(blockChainFile) {
		fmt.Printf("区块链已经存在，无需重复创建!\n")
		return nil
	}

	//创建区块链结构，一般会在创建的时候，添加一个区块，称之为：创世区块
	var bc BlockChain

	//操作数据库
	//创建blockChain.db文件
	db,err := bolt.Open(blockChainFile,0600,nil)

	if err != nil {
		panic(err)
	}

	//操作数据库
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			//第一次调用这个方法，里面没有bucket时，需要创建，并且添加创世块
			b,err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				panic(err)
			}

			coinbaseTX := NewCoinbaseTX(genesisInfo,"中本聪")

			//将创世块写入bucket中
			genesisBlock := NewBlock([]*Transaction{coinbaseTX},[]byte{})

			//更新区块,更新lastHashKey的值
			b.Put(genesisBlock.CurrentBlockHash,genesisBlock.Serialize())   /*区块的字节流*/
			b.Put([]byte(lastHashKey),genesisBlock.CurrentBlockHash)

			bc.tail = genesisBlock.CurrentBlockHash
		}

		return nil
	})

	bc.db = db

	return &bc
}

//获取已经存在的区块链的实例
func GetBlockChain() *BlockChain {

	if !isFileExist(blockChainFile) {
		fmt.Printf("请先创建区块链文件!\n")
		return nil
	}

	var bc BlockChain

	//操作数据库
	db,err := bolt.Open(blockChainFile,0600,nil)

	if err != nil {
		panic(err)
	}

	//操作数据库
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			fmt.Printf("获取区块链实例，bucket不能为空！\n")
			os.Exit(1)
		}

		lastHash := b.Get([]byte(lastHashKey))
		bc.tail = lastHash

		return nil
	})

	bc.db = db

	return &bc
}

//使用bolt字节的ForEach来打印区块链
func (bc *BlockChain) print1() {

	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%x\n",k)
			return nil
		})
		return nil
	})
}

//定义一个专门的迭代器，用于遍历区块链
type Iterator struct {
	//db
	db *bolt.DB
    //最开始指向lastHash,每次调用Next之后，都会向左移动
    currentHash []byte
}

//创建一个迭代器
func (bc *BlockChain) NextIterator() *Iterator {
	it := Iterator {
		db:bc.db,
        currentHash:bc.tail,
	}

	return &it
}

//实现遍历的方法，Next
func (it *Iterator) Next() *Block {
	var block Block

	it.db.View(func(tx *bolt.Tx) error {
		//获取bucket
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			fmt.Printf("遍历区块链时，bucket不应为空!")
			os.Exit(1)
		}

		//获取区块数据，Get(currentHash)
		blockBytesInfo /*block的字节流*/ := b.Get(it.currentHash)

		//反序列化
		block = *Deserialize(blockBytesInfo)

		return nil
	})

	//currentHash向左移动，一定要记得更新
	it.currentHash = block.PrevBlockHash

	return &block
}

func (bc *BlockChain) GetBalance(address string) float64 {
	//遍历账本
	//遍历交易的output，和指定的地址进行比对，如果比对成功，说明这是我的钱，累计即可

	it := bc.NextIterator()

	var totalMoney float64

	//终止条件，当前区块的前哈希为空(nil),跳出循环
	for {
		block := it.Next()

		//遍历交易
		for _,tx := range block.Transactions {
			//遍历output
			for _,output := range tx.TXOutputs {
				//比较地址
				if output.ScriptPubKey == address {
					totalMoney += output.Value
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return totalMoney
}

//使用我们自己实现的迭代器遍历区块链
func (bc *BlockChain) Print2() {

	//创建一个迭代器
	//it := bc.
}


//创建区块链
//func NewBlockChain() *BlockChain {
//[
//	genesisBlock := NewBlock(genesisInfo,[]byte{})
//
//	block := BlockChain{
//		Blocks:[]*Block{genesisBlock},
//	}
//
//	return &block
//}
//
////添加区块链
//func (bc *BlockChain) AddBlock(data string) {
//
//	 //创建一个区块，前区块的哈希值从bc的最后一个区块元素获取即可
//	 lastBlock := bc.Blocks[len(bc.Blocks) - 1]
//
//     //即将添加的区块的哈希值，就是bc中的最后区块的前哈希值
//     prevHash := lastBlock.CurrentBlockHash
//
//     newBlock := NewBlock(data,prevHash)
//
//     //将新的区块添加到数组中
//     bc.Blocks = append(bc.Blocks,newBlock)
//}






