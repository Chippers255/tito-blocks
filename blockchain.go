package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Transaction struct {
	Sender   string
	Amount   float32
	Receiver string
}

type Message struct {
	Sender string
	Text   string
	Chat   string
}

type Block struct {
	Index        int           `json:"index"`
	Proof        string        `json:"proof"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

type Server interface {
	RegisterNode()
	ValidChain() bool
	ResolveConflicts()
	NewBlock() Block
	NewTransaction(sender string, receiver string, amount float32) int
	LastBlock() Block
	Hash(block Block) string
	WorkProof(last Block) string
	ValidProof(lastProof string, proof string, lastHash string) bool
}

func (b Blockchain) LastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

func (b Blockchain) NewTransaction(sender string, receiver string, amount float32) int {
	newTransaction := Transaction{
		Sender:   sender,
		Amount:   amount,
		Receiv3er: receiver,
	}

	b.CurrentTransactions = append(b.CurrentTransactions, newTransaction)

	return b.LastBlock().Index + 1
}

func Hash(block Block) string {
	s, _ := json.Marshal(block)
	s2 := md5.Sum(s)

	d := fmt.Sprintf("%x", s2)

	return d
}
