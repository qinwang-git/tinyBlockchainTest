package BLC

import (
	"bytes"
)

type TXInput struct {
	TxID []byte //交易的ID
	Vout int    //存储Txoutput的vout里面的索引

	Signature []byte //数字签名
	PublicKey []byte //公钥，钱包里面
}

//判断当前txInput消费，和指定的address是否一致
func (txInput *TXInput) UnlockWithAddress(pubKeyHash []byte) bool {
	publicKey := PubKeyHash(txInput.PublicKey)
	return bytes.Compare(pubKeyHash, publicKey) == 0
}
