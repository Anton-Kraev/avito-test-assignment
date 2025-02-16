package models

type Transfer struct {
	ReceiverName string
	SenderID     int
	Amount       int
}

type ReceivedTransfer struct {
	SenderName string
	Amount     int
}

type SentTransfer struct {
	ReceiverName string
	Amount       int
}
