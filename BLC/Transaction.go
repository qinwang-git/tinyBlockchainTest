package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"time"
)

//创建Transaction结构体
type Transaction struct {
	TxID  []byte      //交易ID
	Vins  []*TXInput  //输入
	Vouts []*TXOutput //输出
}

/*
Transaction 创建分两种情况
1.创世区块创建时的Transaction
2.转账时产生的Transaction
*/
//Coinbase交易
func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TXInput{[]byte{}, -1, "Genesis Data"}
	txOutput := &TXOutput{10, address}
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinbase.SetTxID()
	return txCoinbase
}

//普通交易
func NewSimpleTransaction(from, to string, amount int64, bc *BlockChain, txs []*Transaction) *Transaction {
	var txInputs []*TXInput
	var txOutputs []*TXOutput
	balance, spendableUTXO := bc.FindSpendableUTXOs(from, amount, txs)

	//代表消费
	for txID, indexArray := range spendableUTXO {
		txIDBytes, _ := hex.DecodeString(txID)
		for _, index := range indexArray {
			txInput := &TXInput{txIDBytes, index, from}
			txInputs = append(txInputs, txInput)

		}
	}

	//转账
	txOutput1 := &TXOutput{amount, to}
	txOutputs = append(txOutputs, txOutput1)

	//找零
	txOutput2 := &TXOutput{balance - amount, from}
	txOutputs = append(txOutputs, txOutput2)

	tx := &Transaction{[]byte{}, txInputs, txOutputs}

	//设置hash值
	tx.SetTxID()
	return tx
}

//判断当前交易是否是Coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxID) == 0 && tx.Vins[0].Vout == -1

}

//设置交易ID，其实就是hash
func (tx *Transaction) SetTxID() {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	buffBytes := bytes.Join([][]byte{IntToHex(time.Now().Unix()), buff.Bytes()}, []byte{})

	hash := sha256.Sum256(buffBytes)
	tx.TxID = hash[:]
}
