package main

import (
	"./BLC"
	"fmt"
)

func main() {

}
    //1.测试Block
    block:=BLC.NewBlock("I am a block",make([]byte,32,32),1)
    fmt.Println(block)

    //2.测试创世区块
    genesisBlock :=BLC.CreateGenesisBlock("Genesis Block..")
    fmt.Println(genesisBlock)

    //3.测试区块链
    genesisBlockChain := BLC.CreateBlockChainWithGenesisBlock()
    fmt.Println(genesisBlockChain)
    fmt.Println(genesisBlockChain.Blocks)
    fmt.Println(genesisBlockChain.Blocks[0])

    //4.测试添加新区块
    blockChain:=BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")
    blockChain.AddBlockToBlockChain("Send 100RMB To Wangergou",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
    blockChain.AddBlockToBlockChain("Send 300RMB To lixiaohua",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
    blockChain.AddBlockToBlockChain("Send 500RMB To rose",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

    fmt.Println(blockChain)
}