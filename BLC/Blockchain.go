package BLC

type BlockChain struct {
	Blocks []*Block
}

// 创建带有创世区块的链
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock(data)
	//返回区块链对象
	return &BlockChain{[]*Block{genesisBlock}}
}

// 添加新区块入链
func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	//创建新区块
	newBlock := NewBlock(data, prevHash, height)
	//将newBlock添加切片至Block后面
	bc.Blocks = append(bc.Blocks, newBlock)
}
