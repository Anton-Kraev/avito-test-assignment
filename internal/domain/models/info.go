package models

type Info struct {
	Balance             int
	Inventory           []Item
	CoinsReceiveHistory []ReceivedTransfer
	CoinsSentHistory    []SentTransfer
}
