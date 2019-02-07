package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"time"
)

const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
const bits = 16

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
	//Data []byte

	//真正交易的数组
	Transactions []*Transaction
}

//创建区块
func NewBlock(txs []*Transaction,prevBlockchain []byte) *Block {

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
		Nonce:0,
		//难度值
		Bits:bits,

		//数据
		//Data:[]byte(data),

		Transactions:txs,
	}

	//设置哈希
	//block.setHash()

	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()

	block.CurrentBlockHash = hash
	block.Nonce = nonce

	//返回区块
	return &block
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

//序列化，将结构转换为字节流
func (b *Block) Serialize() []byte {
	//创建字节流缓存器
	var buffer bytes.Buffer
	//编码
	//创建编码器
	encoder := gob.NewEncoder(&buffer)

	//编码器Encoder方法，得到字符串
	err := encoder.Encode(b)

	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

//反序列化,将字节流转成block
func Deserialize(data []byte) *Block {
	var block Block

	//解码器
	decoder := gob.NewDecoder(bytes.NewReader(data))

	//解码器Decode方法，结构
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}

