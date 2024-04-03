package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
)

var generatedCount = 0
var numThreads = 16
var startTime = time.Now()
var searchTerm = ""
var shouldStopThreads = false

func generateWallet() {
	for {
		if shouldStopThreads {
			return
		}
		newWallet := solana.NewWallet()
		if strings.HasPrefix(newWallet.PublicKey().String(), searchTerm) && !shouldStopThreads {
			firstCharAfterSearchTerm := strings.Split(newWallet.PublicKey().String(), searchTerm)[1][0:1]
			if firstCharAfterSearchTerm == strings.ToUpper(firstCharAfterSearchTerm) {
				fmt.Printf("Success! Wallet found: %s\n", newWallet.PublicKey())
				fmt.Printf("Secret Key: %v\n", newWallet.PrivateKey)
				fmt.Printf("Attempts required: %d, Time elapsed: %s\n", generatedCount+1, time.Since(startTime))
				shouldStopThreads = true
			}
		}
		generatedCount++
		if generatedCount%1000000 == 0 {
			fmt.Printf("Status: %d wallets generated in %s\n", generatedCount, time.Since(startTime))
		}
	}
}

func main() {
	fmt.Printf("Target prefix: %s\n", searchTerm)
	for i := 0; i < numThreads; i++ {
		go generateWallet()
	}
	fmt.Scanln()
}
