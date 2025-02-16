package models

type Transfer struct {
	ReceiverID int
	SenderID   int
	Amount     int
}

type ReceivedTransfer struct {
	SenderName string
	Amount     int
}

type SentTransfer struct {
	ReceiverName string
	Amount       int
}
