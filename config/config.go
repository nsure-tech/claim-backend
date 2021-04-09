package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type GbeConfig struct {
	DataSource              DataSourceConfig `json:"dataSource"`
	RestServer              RestServerConfig `json:"restServer"`
	Log                     LogConfig        `json:"log"`
	JwtSecret               string           `json:"jwtSecret"`
	RawUrl                  string           `json:"rawUrl"`
	ChainId                 int64            `json:"chainId"`
	ContractTreasuryAddress string           `json:"contractTreasuryAddress"`
	ContractBuyAddress      string           `json:"contractBuyAddress"`
	DepositAddress          string           `json:"depositAddress"`
	KeySecret               string           `json:"keySecret"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}
type LogConfig struct {
	Level      int    `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

var config GbeConfig
var configOnce sync.Once

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("./vote.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
