package main

import "os"

func isFileExist(fileName string) bool {

	_,err := os.Stat(fileName)

	//一定不要使用os.IsExist
	if os.IsNotExist(err) {
		return false
	}

	return true
}
