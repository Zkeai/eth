package module_check

import (
	Wallet "eth_wallet/wallet"
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

func Create() {

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
