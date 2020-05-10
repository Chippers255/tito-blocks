package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Foo"

	md5 := md5.Sum([]byte(s))
	sha1 := sha1.Sum([]byte(s))
	sha256 := sha256.Sum256([]byte(s))

	fmt.Printf("%x\n", md5)
	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha256)
	fmt.Println("-----------------------------------------------------")
	b := Block{
		Index:        1,
		Proof:        "sdfd",
		Timestamp:    "asdfasdf",
		Transactions: nil,
		PreviousHash: "adsfasdf",
	}

	a := Hash(b)
	fmt.Println(b)
	fmt.Println(a)
}
