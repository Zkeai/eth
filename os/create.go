package os

import (
	"fmt"
	"os"
)

func WriteAds(privateKeyEnd string, adsEnd string) {
	file, err := os.OpenFile("address.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("file close err: %v \n", err)
			return
		}
	}(file)
	content := privateKeyEnd + "----" + adsEnd + "\n"
	_, err = file.WriteString(content)

	if err != nil {
		fmt.Printf("WriteString err: %v \n", err)
		return
	}

}
