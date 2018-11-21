package BLC

// The original tutorial is linked to https://github.com/rubyhan1314/PublicChain
// Thanks for the open source code abovementioned.

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// build Block Struct
type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
	TimeStamp     int64
	Hash          []byte //  32字节，64个16进制数

}

// the function of creating new block
func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	block.SetHash() // 设置hash 接后面func
	return block
}

// set the hash
func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height) // 转换为字节数组
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	timeBytes := []byte(timeString)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes},
		[]byte{})

	//4.生成哈希值
	hash := sha256.Sum256(blockBytes) //数组长度32位
	block.Hash = hash[:]
}

// genesis block
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}
