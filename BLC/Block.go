package BLC

// The original tutorial is linked to https://github.com/rubyhan1314/PublicChain
// Thanks for the open source code abovementioned.

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

// build Block Struct
type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
	TimeStamp     int64
	Hash          []byte //  32字节，64个16进制数
	Nonce         int64  //随机数

}

// the function of creating new block
func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	//调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// genesis block
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}

//将区块序列化，得到一个字节数组---区块的行为，设计为方法
func (block *Block) Serilalize() []byte {
	//创建一个buffer
	var result byte.Buffer
	//创建一个编码器
	encoder:=gob.NewEncoder(&result)
	//编码--->打包
	err:=encoder.Encode(block)
	if err:=nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化，得到一个区块---设计为函数
func DeserializeBlock(blockBytes []byte) *Block{
	var block Block
	var reader=byte.NewReader(blockBytes)
	//创建一个解码器
	decoder:=gob.NewDecoder(reader)
	//解包
	err:=decoder.Decode(&block)
	if err!=nil{
		log.Panic(err)
	}
	return &block
}
