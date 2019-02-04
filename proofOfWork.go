package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//定义一个ProofOfWork的结构
//区块
//目标值

type ProofOfWork struct {
	//区块
	block Block

	//目标值
	target big.Int
}

//创建一个ProofOfWork的方法
func NewProofOfWork(block Block) (*ProofOfWork) {
	//定义一个数值
	bigIntTmp := big.NewInt(1)

	//向左移动256 - 16，16就是难度值
	bigIntTmp.Lsh(bigIntTmp,256 - bits)

    pow := ProofOfWork{
    	block:block,
    	target:*bigIntTmp,
	}

    return &pow
}

//
func (pow *ProofOfWork) prepareData(nonce uint64) []byte {

	b := pow.block

	tmp := [][]byte{
		uint2Bytes(b.Version),
		b.PrevBlockHash,
		b.MerkeRoot,
		uint2Bytes(b.TimeStamp),
		uint2Bytes(b.Nonce),
		uint2Bytes(b.Bits),
	    b.Data,
	}

	blockInfo := bytes.Join(tmp,[]byte{})

	return blockInfo
}

//给ProofOfWork提供一个计算的方法,用于找到Nonce
func (pow *ProofOfWork) Run() (uint64,[]byte) {

	//定义一个nonce变量，用于不断变化
	var nonce uint64
	var hash [32]byte

	for nonce <= math.MaxInt64 {
		fmt.Printf("%x\r",hash)

		//对拼接好的数据进行sha256运算，得到一个哈希值，需要转换为big.Int
		hash = sha256.Sum256(pow.prepareData(nonce))

		//转换为big.Int类型的数据
		var bigIntTmp big.Int

		bigIntTmp.SetBytes(hash[:])

		//如果生成哈希小于目标值，满足条件，返回哈希值，nonce，直接退出
		if bigIntTmp.Cmp(&pow.target) == -1 {
			fmt.Printf("挖矿成功,hash:%x,nonce:%d\n",hash,nonce)
			break
		} else {
			//如果生成哈希大于目标值，不满足条件，nonce++,继续遍历
			nonce++
		}
	}
	return nonce,hash[:]
}

//提供一个校验方法，用于检测挖矿得到的随机数是否满足系统的条件
func (pow *ProofOfWork) IsVaild() bool {
	//得到哈希
	//做哈希值
	//与系统的哈希值比较

	block := pow.block

	//旷工校验时，会拿到区块数据，然后自己校验哈希
	data := pow.prepareData(block.Nonce)
	fmt.Printf("-----isVaild,Nonce:%d\n",block.Nonce)

	hash := sha256.Sum256(data)
	fmt.Printf("-----isVaild,block hash:%x\n",block.CurrentBlockHash)
	fmt.Printf("-----isVaild,Hash:%x\n",hash)

	//哈希值与bit.Int比较
	var bigIntTmp big.Int

	bigIntTmp.SetBytes(hash[:])
	fmt.Printf("pow.target:%x\n",pow.target.Bytes())

	res := bigIntTmp.Cmp(&pow.target)
	fmt.Printf("bigIntTmp:%x\n",bigIntTmp.Bytes())

	fmt.Printf("res:%x\n",res)

	if bigIntTmp.Cmp(&pow.target) == -1 {
		fmt.Printf("111111\n")
		return true
	}

	return false
}

