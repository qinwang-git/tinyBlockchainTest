package BLC

import "fmt"

func (cli *CLI) addressLists() {
	fmt.Println()

	Wallets := NewWallets()
	for address, _ := range Wallets.WalletsMap {
		fmt.Println("address:", address)
	}
}
