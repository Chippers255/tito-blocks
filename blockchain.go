package main

type Transaction struct {
	Sender    string  `json:"sender"`
	Amount    float32 `json:"amount"`
	Recipient string  `json:"recipient"`
}

type Block struct {
	Index        int           `json:"index"`
	Proof        string        `json:"proof"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

type Blockchain struct {
	Chain               []Block       `json:"chain"`
	CurrentTransactions []Transaction `json:"current_transactions"`
}

func (b *Blockchain) NewTransaction(sender string, recipient string, amount float32) int {
	tx := Transaction{Sender: sender}

	b.CurrentTransactions = append(b.CurrentTransactions, tx)

	return 1
}
