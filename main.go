package main

import (
	"log"
	"os"
	_ "GoInActionCh2/matchers"	// 前面下划线的意思是：调用包里面的init函数进行初始化，但并不适用包里的标识符
	"GoInActionCh2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
