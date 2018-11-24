package BLC

import (
	"bytes"
)

type TXOutput struct {
	Value      int64
	PubKeyHash []byte // 公钥
}

//判断当前txOutput消费，和指定的address是否一致
func (txOutput *TXOutput) UnLockWithAddress(address string) bool {
	fullPaylaodHash := Base58Decode([]byte(address))
	pubKeyHash := fullPaylaodHash[1 : len(fullPaylaodHash)-4]
	return bytes.Compare(txOutput.PubKeyHash, pubKeyHash) == 0
}
