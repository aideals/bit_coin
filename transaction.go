package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

//交易输入
//指明交易发起人可支付资金的来源，包含
type TXInput struct {
	//引用utxo所在交易的id，（知道再哪个房间）
	TXID []byte

	//所消费utxo在output中的索引
	Index int64

	//解锁脚本(签名，公钥)
	ScriptSig string
}

//交易输出(TXOutPut)
//包含资金接收方的相关信息
type TXOutput struct {
	//接收金额
	Value float64
	//锁定脚本（对方公钥的哈希，这个哈希可以通过地址反推出来，所以转账时知道地址就好）
	ScriptPubKey string
}

//定义交易结构体
type Transaction struct {
	//交易ID
	TxID []byte
	//输入数组
	TXIputs []TXInput
	//输出数组
	TXOutputs []TXOutput
	//时间戳
	TimeStamp uint64
}

//一般交易ID
//一般是交易结构的哈希值(参考block的哈希做法)
func (tx *Transaction) SetTXHash() {
	//交易id我们使用sha256来获取
	//获取tx的字节流
	var buffer bytes.Buffer

	//创建编码器
	encoder := gob.NewEncoder(&buffer)

	//编码器Encoder方法，得到字节流
	err := encoder.Encode(tx)

	if err != nil {
		panic(err)
	}

	//调用sha256方法
	hash := sha256.Sum256(buffer.Bytes())

	tx.TxID = hash[:]
}

const reward = 12.5

//挖矿交易
func NewCoinbaseTX(data string,miner string /*矿工地址，使用string代替*/) *Transaction {

	txinput := TXInput{
		TXID:nil,
		Index:-1,
		ScriptSig:data,
	}

	txoutput := TXOutput{
		Value:reward,
		ScriptPubKey:miner,
	}

	tx := Transaction{
		nil,
		[]TXInput{txinput},
		[]TXOutput{txoutput},
		uint64(time.Now().Unix()),
	}

	//调用setTXHash
	tx.SetTXHash()

	return &tx
}
