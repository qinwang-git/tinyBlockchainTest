package BLC

type TXOutput struct {
	Value int64
	//一个锁定脚本(ScriptPubKey)，要花这笔钱，必须要解锁该脚本。
	ScriptPubKey string
}

//判断当前txOutput消费，和指定的address是否一致
func (txOutput *TXOutput) UnLockWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address
}
