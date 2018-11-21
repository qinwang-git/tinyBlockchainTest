package BLC

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock(data string) *BlockChain {

	genesisBlock := CreateGenesisBlock(data)
	return &BlockChain{[]*Block{genesisBlock}}
}

func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {

	newBlock := NewBlock(data, prevHash, height)
	//将newBlock添加切片至Block后面
	bc.Blocks = append(bc.Blocks, newBlock)
}
