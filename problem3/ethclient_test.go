package problem3

import (
	"github.com/go-resty/resty/v2"
	"slices"
	"testing"
)

func Test_ethClient_BroadcastTransaction(t *testing.T) {
	type fields struct {
		config      *Config
		restyClient *resty.Client
	}
	type args struct {
		payload BroadcastTransactionPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				config:      NewConfig(),
				restyClient: resty.New(),
			},
			args: args{
				payload: BroadcastTransactionPayload{
					Symbol:    "ETH",
					Price:     4500,
					Timestamp: 1678912345,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ethClient{
				config:      tt.fields.config,
				restyClient: tt.fields.restyClient,
			}
			got, err := c.BroadcastTransaction(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.TxHash == "" {
				t.Errorf("BroadcastTransaction() got = %v, want non-empty txHash", got)
			}
		})
	}
}

func Test_ethClient_CheckTransactionStatus(t *testing.T) {
	type fields struct {
		config      *Config
		restyClient *resty.Client
	}
	type args struct {
		txHash string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		transactionStatuses []TransactionStatus
		wantErr             bool
	}{
		{
			name: "success",
			fields: fields{
				config:      NewConfig(),
				restyClient: resty.New(),
			},
			args: args{
				txHash: "e97ae379d666eed7576bf285c451baf9f0edba93ce718bf15b06c8a85d07b8d1",
			},
			transactionStatuses: []TransactionStatus{TransactionStatusPending, TransactionStatusConfirm, TransactionStatusFailed, TransactionStatusDNE},
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ethClient{
				config:      tt.fields.config,
				restyClient: tt.fields.restyClient,
			}
			got, err := c.CheckTransactionStatus(tt.args.txHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckTransactionStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && slices.Contains(tt.transactionStatuses, got.Status) == false {
				t.Errorf("CheckTransactionStatus() got = %v, want one of %v", got.Status, tt.transactionStatuses)
			}
		})
	}
}
