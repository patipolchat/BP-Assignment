# BP-Assignment

## Problem 1: Boss Baby's Revenge

```go
problem1.BossBabyRevenge("SRSSRRR") //return "Good Boy"
problem1.BossBabyRevenge("RSSRR") //return "Bad Boy"
```

## Problem 2: Superman Chicken Rescue
```go
problem2.SupermanChickenRescue(5, 5, []int{2, 5, 10, 12, 15}) //return 2 
problem2.SupermanChickenRescue(6, 10, []int{1, 11, 30, 34, 35, 37}) //return 4
```

## Problem 3: Transaction Broadcasting and Monitoring Client

### Setup
```go
conf := problem3.NewConfig()
ethClient := problem3.NewEthClient(conf)
```

### Broadcast Transaction
```go
payload := problem3.BroadcastTransactionPayload{
  Symbol:    "ETH",
  Price:     4500,
  Timestamp: 1678912345,
}

txResp, err := ethClient.BroadcastTransaction(payload)
```

### Monitor Transaction
```go
txHash := "0x1234567890"
txStatusResp, err := ethClient.CheckTransactionStatus(txHash)
```