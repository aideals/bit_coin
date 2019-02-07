package main

import (
	"fmt"
	"os"
)

//定义一个CLI结构
//不需要字段
//需要一个run方法，负责解析，调用相关命令

const Usage = `
     ./block createBC "创建区块链数据库"
     ./block addBlock "DATA" 添加数据到区块链
     ./block printChain "打印区块链数据"
     ./block getBalance address "获取指定地址的余额"
`

type CLI struct {

}

func (cli *CLI) Run() {
	cmds := os.Args

	if len(cmds) < 2 {
		fmt.Println(Usage)
		os.Exit(1)
	}

	switch cmds[1] {
	case "CreateBc":
		fmt.Printf("创建区块链命令被调用!\n")
	    cli.createBC()

	case "addBlock":
		fmt.Printf("添加区块链命令被调用\n")

	    if len(cmds) != 3 {
	    	fmt.Printf("参数不足!\n")
	    	os.Exit(1)
		}

	    data := cmds[2]
	    cli.addBlock(data)

	case "printChain":
		fmt.Printf("打印区块链命令被调用\n")
	    cli.printChain()

	case "getBalance":
		fmt.Printf("查询余额命令被调用!\n")

	    address := cmds[2]
	    cli.getBalance(address)

	default:
		fmt.Printf("无效的命令,%s\n",cmds[1])
	    fmt.Println(Usage)
	    os.Exit(1)
	}
}
