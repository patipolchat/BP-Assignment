package problem3

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusConfirm TransactionStatus = "CONFIRMED"
	TransactionStatusFailed  TransactionStatus = "FAILED"
	TransactionStatusDNE     TransactionStatus = "DNE"
)

type BroadcastTransactionPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type TransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type TransactionStatusResponse struct {
	Status TransactionStatus `json:"tx_status"`
}

type EthClient interface {
	BroadcastTransaction(payload BroadcastTransactionPayload) (*TransactionResponse, error)
	CheckTransactionStatus(txHash string) (*TransactionStatusResponse, error)
}

type ethClient struct {
	config      *Config
	restyClient *resty.Client
}

func NewEthClient(config *Config) EthClient {
	client := resty.New()
	return &ethClient{
		config:      config,
		restyClient: client,
	}
}

func (c *ethClient) BroadcastTransaction(payload BroadcastTransactionPayload) (*TransactionResponse, error) {
	var txResp TransactionResponse
	url := c.config.GetBroadcastTxUrl()
	resp, err := c.restyClient.R().
		EnableTrace().
		SetResult(&txResp).
		SetBody(payload).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to broadcast transaction: %s", resp.String())
	}
	return &txResp, nil
}

func (c *ethClient) CheckTransactionStatus(txHash string) (*TransactionStatusResponse, error) {
	var txStatusResp TransactionStatusResponse
	url := c.config.GetStatusTxUrl(txHash)
	resp, err := c.restyClient.R().
		EnableTrace().
		SetResult(&txStatusResp).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to check transaction status: %s", resp.String())
	}
	return &txStatusResp, nil
}
