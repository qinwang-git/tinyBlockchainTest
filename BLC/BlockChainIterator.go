package BLC

import (
	"log"

	"github.com/boltdb/bolt/"
)

//新增结构体
type BlockChainIterator struct {
	//当前区块hash
	CurrentHash []byte
	//数据库
	DB *bolt.DB
}

//获取区块
func (bcIterator *BlockChainIterator) Next() *Block {
	block := new(Block)
	//打开数据库并读取
	err := bcIterator.DB.View(func(tx *bolt.Tx) error {
		//打开数据表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			//根据当前hash获取数据并反序列化
			blockBytes := b.Get(bcIterator.CurrentHash)
			block = DeserializeBlock(blockBytes)
			//更新当前hash
			bcIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
