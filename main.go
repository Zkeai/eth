package main

import (
	MuyuModule "eth_wallet/module_check"
	"fmt"
	"time"
)

var checked int

func main() {
	fmt.Println("Evm工具 v1.0.0 -by 木鱼.")
	fmt.Println("请选择模块⬇️")
	fmt.Printf("1.钱包模块\n2.暂未开发\n")
	_, _ = fmt.Scanf("%d", &checked)

	switch checked {
	case 1:
		MuyuModule.Create()
	case 2:
		fmt.Println("暂未开发,即将退出")
		time.Sleep(2)
		return
	}
}
