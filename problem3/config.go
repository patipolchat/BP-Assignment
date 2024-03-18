package problem3

import "fmt"

type Config struct {
	BaseUrl         string
	BroadcastTxPath string
	StatusTxPath    string
}

func NewConfig() *Config {
	return &Config{
		BaseUrl:         "https://mock-node-wgqbnxruha-as.a.run.app",
		BroadcastTxPath: "/broadcast",
		StatusTxPath:    "/check/%s",
	}
}

func (c *Config) GetBroadcastTxUrl() string {
	return c.BaseUrl + c.BroadcastTxPath
}

func (c *Config) GetStatusTxUrl(txHash string) string {
	return c.BaseUrl + fmt.Sprintf(c.StatusTxPath, txHash)
}
