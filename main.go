package main

import (
	Wallet "eth_wallet/wallet"
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var num int

func main() {
	fmt.Print("请输入需要创建的钱包数量：")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Wallet.CreateAds()
		}()
	}
	wg.Wait()

}
