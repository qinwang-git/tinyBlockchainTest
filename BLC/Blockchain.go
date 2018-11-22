package BLC

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	//存储有序的区块，通过array,但不可持久化
	//Blocks []*Block

	//最后区块的Hash
	Tip []byte
	//数据库对象
	DB *bolt.DB
}

// 创建带有创世区块的链
func CreateBlockChainWithGenesisBlock(data string) {

	/*
	   判断数据库是否存在。如果数据库存在：
	     1.创建BlockChain实例。
	     2.读取数据库中最后一个区块的hash，并设置给BlockChain实例的Tip字段。
	   如果数据库不存在：
	     1.首先我们需要先创建一个创世区块
	     2.打开数据库，并且创建bucket
	     3.将创世区块序列化后存入到数据库中
	     4.将创世区块的hash保存为最后一个块的hash
	     5.创建BlockChain实例，设置Tip为创世区块的hash，并返回该blockchain实例。
	*/

	//先判断数据库是否存在，如果有，从数据库读取
	if dbExists() {
		fmt.Println("数据库已经存在。。")
		return
	}
	fmt.Println("创建创世区块：", data)

	//数据库不存在，说明第一次创建，然后存入到数据库中
	fmt.Println("数据库不存在。。")
	//创建创世区块
	genesisBlock := CreateGenesisBlock(data)
	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//存入数据表
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(BLOCKTABLENAME))
		if err != nil {
			log.Panic(err)
		}
		if b != nil {
			err = b.Put(genesisBlock.Hash, genesisBlock.Serilalize())
			if err != nil {
				log.Panic("创世区块存储有误。。。")
			}
			//存储最新区块的hash
			b.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//返回区块链对象
	//return &BlockChain{genesisBlock.Hash, db}
}

// 添加新区块入链
func (bc *BlockChain) AddBlockToBlockChain(data string) {
	//创建新区块
	//newBlock := NewBlock(data, prevHash, height)
	//将newBlock添加切片至Block后面
	//bc.Blocks = append(bc.Blocks, newBlock)

	//更新数据库
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//打开表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			//根据最新块的hash读取数据，并反序列化最后一个区块
			blockBytes := b.Get(bc.Tip)
			lastBlock := DeserializeBlock(blockBytes)
			//创建新的区块
			newBlock := NewBlock(data, lastBlock.Hash, lastBlock.Height+1)
			//将新的区块序列化并存储
			err := b.Put(newBlock.Hash, newBlock.Serilalize())
			if err != nil {
				log.Panic(err)
			}
			//更新最后一个哈希值，以及blockchain的tip
			b.Put([]byte("l"), newBlock.Hash)
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

//判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(DBNAME); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
//新增方法，遍历数据库，打印输出所有的区块信息
func (bc *BlockChain) PrintChains() {
	var currentHash=bc.Tip
	var count=0
	block:=new(Block)
	for{
		err:=bc.DB.View(func(tx *bolt.Tx)error{
			b:=tx.Bucket([]byte(BLOCKTABLENAME))
			if b:=nil{
				count++
				fmt.Print("第%d个区块的信息：\n", count)
				blockBytes:=b.Get(currentHash)
				block=DeserializeBlock(blockBytes)
				fmt.Printf("\t高度：%d\n", block.Height)
				fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
                fmt.Printf("\t当前的hash：%x\n", block.Hash)
                fmt.Printf("\t数据：%s\n", block.Data)
                fmt.Printf("\t时间：%v\n", block.TimeStamp)
                fmt.Printf("\t时间：%s\n",time.Unix(block.TimeStamp,0).Format("2006-01-02 15:04:05"))
                fmt.Printf("\t次数：%d\n", block.Nonce)
			}
			return nil
		})
		if err !=nil{
			log.Panic(err)
		}
		hashInt:=new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt)==0{
			break
		}
		currentHash=block.PrevBlockHash
	}
}
*/

//获取一个迭代器的方法
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.Tip, bc.DB}
}

//遍历数据库中的Block
func (bc *BlockChain) PrintChains() {
	//获取迭代器对象
	bcIterator := bc.Iterator()
	var count = 0

	//循环迭代
	for {
		block := bcIterator.Next()
		count++
		fmt.Printf("第%d个区块的信息：\n", count)
		//获取当前hash对应的数据，并进行反序列化
		fmt.Printf("\t高度：%d\n", block.Height)
		fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
		fmt.Printf("\t当前的hash：%x\n", block.Hash)
		fmt.Printf("\t数据：%s\n", block.Data)
		//fmt.Printf("\t时间：%v\n", block.TimeStamp)
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("\t次数：%d\n", block.Nonce)

		//直到父hash值为0
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}
	}
}

//获取区块链
func GetBlockchainObject() *BlockChain {
	/*
	   1.如果数据库不存在，直接返回nil
	   2.读取数据库
	*/
	if !dbExists() {
		fmt.Println("数据库不存在，无法获取区块链。。")
		return nil
	}

	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockchain *BlockChain
	//读取数据库
	err = db.View(func(tx *bolt.Tx) error {
		//打开表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			//读取最后一个hash
			hash := b.Get([]byte("l"))
			//创建blockchain
			blockchain = &BlockChain{hash, db}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return blockchain
}
