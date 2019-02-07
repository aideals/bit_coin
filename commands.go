package main

import (
	"fmt"
)

func (cli *CLI) createBC() {

	bc := CreateBlockChain()

	if bc == nil {
		return
	}

	//关闭数据库
	defer bc.db.Close()
}

func (cli *CLI) addBlock(data string) {

	bc := GetBlockChain()

	if bc == nil {
		return
	}

	//关闭数据库
	defer bc.db.Close()
}

func (cli *CLI) printChain() {

	bc := GetBlockChain()

	if bc == nil {
		return
	}

	//关闭数据库
	defer bc.db.Close()

	bc.Print2()
}

func (cli *CLI) getBalance(address string) {

	bc := GetBlockChain()

	if bc == nil {
		return
	}

	defer bc.db.Close()

	balance := bc.GetBalance(address)
	fmt.Printf("'%s'的余额为:%f\n",address,balance)
}
