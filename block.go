package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//创建区块结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希值
	PrevBlockHash []byte
	//当前区块哈希值
	CurrentBlockHash []byte
	//时间戳
	TimeStamp uint64
	//梅克耳根
	MerkeRoot []byte
	//难度值
	Bits uint64
	//随机数
	Nonce uint64
	//数据
	Data []byte
}

//创建区块
func NewBlock(data string,prevBlockchain []byte) *Block {

	block := Block{
		//版本号
		Version:00,
		//前区块哈希值
		PrevBlockHash:prevBlockchain,
		//当前区块哈希值
		CurrentBlockHash:nil,
		//时间戳
		TimeStamp:uint64(time.Now().Unix()),
		//梅克尔根
		MerkeRoot:nil,
		//随机数
		Nonce:00,
		//难度值
		Bits:00,

		//数据
		Data:[]byte(data),
	}

	//设置哈希
	block.setHash()

	//返回区块
	return &block
}

//把区块的数据进行sha256运算
func (b *Block) setHash() {

	var blockInfo []byte

	//把新的字段添加进去
	blockInfo = append(blockInfo,uint2Bytes(b.Version)...)
	blockInfo = append(blockInfo,b.PrevBlockHash...)
	blockInfo = append(blockInfo,b.CurrentBlockHash...)
	blockInfo = append(blockInfo,uint2Bytes(b.TimeStamp)...)
	blockInfo = append(blockInfo,b.MerkeRoot...)
	blockInfo = append(blockInfo,uint2Bytes(b.Bits)...)
	blockInfo = append(blockInfo,uint2Bytes(b.Nonce)...)

	//进行sha256运算
	hash := sha256.Sum256(blockInfo)

	b.CurrentBlockHash = hash[:]
}

//把整数转换成字节流
func uint2Bytes(num uint64) []byte {

	//创建字节缓冲器
	var buffer bytes.Buffer

	err := binary.Write(&buffer,binary.BigEndian,num)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

