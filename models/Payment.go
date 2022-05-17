package models

type Payment struct {
	SenderAccount   int     `json:senderAccount`
	ReceiverAccount int     `json:receiverAccount`
	Amount          float32 `json:amount`
}
