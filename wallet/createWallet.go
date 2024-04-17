package wallet

import (
	myos "eth_wallet/os"

	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func CreateAds() {
	//私钥生成

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyEnd := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//公钥获取
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	//生成地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	adsEnd := hexutil.Encode(hash.Sum(nil)[12:])

	//保存到本地
	myos.WriteAds(privateKeyEnd, adsEnd)
}
