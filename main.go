package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

type resp struct {
	Message      string        `json:"message"`
	Index        int           `json:"index"`
	Transactions []Transaction `json:"transactions"`
	Proof        string        `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

type resp2 struct {
	Chain  []Block `json:"chain"`
	Length int     `json:"length"`
}

func (b *Blockchain) mine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lastBlock := b.LastBlock()
	lastProof := lastBlock.Proof
	proof := b.ProofOfWork(lastProof)

	nt := Transaction{
		Sender:    "0",
		Amount:    1,
		Recipient: viper.GetString("NODE_ID"),
	}

	b.NewTransaction(nt)
	previousHash := Hash(lastBlock)
	block := b.NewBlock(proof, previousHash)

	response := resp{
		Message:      "New Block Forged",
		Index:        block.Index,
		Transactions: block.Transactions,
		Proof:        block.Proof,
		PreviousHash: block.PreviousHash,
	}

	_ = json.NewEncoder(w).Encode(&response)
}

func (b *Blockchain) chain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := resp2{
		Chain:  b.Chain,
		Length: len(b.Chain),
	}

	_ = json.NewEncoder(w).Encode(&response)
}

func main() {
	nodeID := strings.ReplaceAll(uuid.New().String(), "-", "")
	viper.Set("NODE_ID", nodeID)

	router := mux.NewRouter()
	b := NewBlockchain()

	// Dispatch map for CRUD operations.
	router.HandleFunc("/mine", b.mine).Methods("GET")
	router.HandleFunc("/chain", b.chain).Methods("GET")
	router.HandleFunc("/transactions/new", b.NewTransactionEndpoint).Methods("POST")

	// Start the server.
	port := ":5000"
	fmt.Println("\nListening on port " + port)
	_ = http.ListenAndServe(port, router)
}
