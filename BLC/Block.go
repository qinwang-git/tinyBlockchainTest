package BLC

// The original tutorial is linked to https://github.com/rubyhan1314/PublicChain
// Thanks for the open source code abovementioned.

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// build Block Struct
type Block struct {
	Height        int64
	PrevBlockHash []byte
	Txs           []*Transaction
	TimeStamp     int64
	Hash          []byte //  32字节，64个16进制数
	Nonce         int64  //随机数

}

// the function of creating new block
func NewBlock(txs []*Transaction, prevBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil, 0}

	//调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// genesis block
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, make([]byte, 32, 32), 0)
}

//将区块序列化，得到一个字节数组---区块的行为，设计为方法
func (block *Block) Serilalize() []byte {
	//创建一个buffer
	var result bytes.Buffer
	//创建一个编码器
	encoder := gob.NewEncoder(&result)
	//编码--->打包
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化，得到一个区块---设计为函数
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	//创建一个解码器
	decoder := gob.NewDecoder(reader)
	//解包
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//将Txs转为[]byte
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
