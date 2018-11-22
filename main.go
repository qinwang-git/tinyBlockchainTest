package main

import (
	"./BLC"
)

func main() {

	//1.测试Block
	//block:=BLC.NewBlock("I am a block",make([]byte,32,32),1)
	//fmt.Println(block)

	//2.测试创世区块
	//genesisBlock :=BLC.CreateGenesisBlock("Genesis Block..")
	//fmt.Println(genesisBlock)

	//3.测试区块链
	//genesisBlockChain := BLC.CreateBlockChainWithGenesisBlock()
	//fmt.Println(genesisBlockChain)
	//fmt.Println(genesisBlockChain.Blocks)
	//fmt.Println(genesisBlockChain.Blocks[0])

	//4.测试添加新区块
	//blockChain := BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")

	//blockChain.AddBlockToBlockChain("Send 100RMB To Wangergou", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 300RMB To lixiaohua", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 500RMB To rose", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	//fmt.Println(blockChain)
	//for _, block := range blockChain.Blocks {
	//	pow := BLC.NewProofOfWork(block)
	//	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.IsValid()))
	//}

	/*
		// 5.检测pow
		//1.创建一个big对象 0000000.....00001
		target := big.NewInt(1)
		fmt.Printf("0x%x\n",target) //0x1

		//2.左移256-bits位
		target = target.Lsh(target, 256-BLC.TargetBit)

		fmt.Printf("0x%x\n",target) //61
		//61位：0x1000000000000000000000000000000000000000000000000000000000000
		//64位：0x0001000000000000000000000000000000000000000000000000000000000000

		s1:="HelloWorld"
		hash:=sha256.Sum256([]byte(s1))
		fmt.Printf("0x%x\n",hash)
	*/

	/*
		//创建区块，存入数据库
		//打开数据库
		block := BLC.NewBlock("helloworld", make([]byte, 32, 32), 0)
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//创建表
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("blocks"))
			if b == nil {
				//创建Bucket
				b, err = tx.CreateBucket([]byte("blocks"))
				if err != nil {
					log.Panic("创建表失败")
				}
			}
			//向表中存储数据
			err = b.Put([]byte("l"), block.Serilalize())
			if err != nil {
				log.Panic(err)
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("blocks"))
			if b != nil {
				////读取最后一个hash
				data := b.Get([]byte("l"))
				//fmt.Printf("%s\n",data)//直接打印会乱码
				//反序列化
				block2 := BLC.DeserializeBlock(data)
				//fmt.Println(block2)
				fmt.Printf("%v\n", block2)
			}
			return nil
		})
	*/

	//测试创世区块存入数据库
	//blockchain := BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	//fmt.Println(blockchain)
	//defer blockchain.DB.Close()

	//测试新添加的区块
	//blockchain.AddBlockToBlockChain("Send 100RMB to wangergou")
	//blockchain.AddBlockToBlockChain("Send 100RMB to lixiaohua")
	//blockchain.AddBlockToBlockChain("Send 100RMB to rose")
	//fmt.Println(blockchain)
	//blockchain.PrintChains()

	//CLI操作
	cli := BLC.CLI{}
	cli.Run()

}
