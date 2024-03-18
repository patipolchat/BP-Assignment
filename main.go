package main

import (
	"encoding/json"
	"fmt"
	"github.com/patipolchat/BP-Assignment/problem1"
	"github.com/patipolchat/BP-Assignment/problem2"
	"github.com/patipolchat/BP-Assignment/problem3"
)

func main() {
	fmt.Println("Problem 1: Boss Baby's Revenge")
	inputs := []string{"SRSSRRR", "RSSRR", "SSSRRRRS", "SSRR", "SRRSSR"}
	for caseNo, input := range inputs {
		fmt.Printf("Case %d: %s => %s\n", caseNo+1, input, problem1.BossBabyRevenge(input))
	}

	fmt.Println("\nProblem 2: Superman Chicken Rescue")
	chickenPos := [][]int{{2, 5, 10, 12, 15}, {1, 11, 30, 34, 35, 37}}
	lengthOfRoof := []int{5, 10}
	for caseNo, pos := range chickenPos {
		maximumChicken, err := problem2.SupermanChickenRescue(len(pos), lengthOfRoof[caseNo], pos)
		if err != nil {
			fmt.Printf("Case %d: %v\n", caseNo+1, err)
		} else {
			fmt.Printf("Case %d: %v\n", caseNo+1, maximumChicken)
		}
	}

	fmt.Println("\nProblem 3: Ethereum Client")
	conf := problem3.NewConfig()
	ethClient := problem3.NewEthClient(conf)

	fmt.Println("BroadcastTransaction")
	payload := problem3.BroadcastTransactionPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: 1678912345,
	}
	fmt.Printf("WithPayload: %+v\n", payload)
	txResp, err := ethClient.BroadcastTransaction(payload)
	txRespJson, _ := json.Marshal(txResp)
	if err != nil {
		fmt.Println("BroadcastTransaction error:", err)
	} else {
		fmt.Println("BroadcastTransaction response:", string(txRespJson))
	}

	fmt.Println("\nCheckTransactionStatus")
	txHash := "0x1234567890"
	fmt.Printf("WithTxHash: %s\n", txHash)
	txStatusResp, err := ethClient.CheckTransactionStatus(txHash)
	txStatusRespJson, _ := json.Marshal(txStatusResp)
	if err != nil {
		fmt.Println("CheckTransactionStatus error:", err)
	} else {
		fmt.Println("CheckTransactionStatus response:", string(txStatusRespJson))
	}
}
