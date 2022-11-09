package titaniumutils

import (
	"context"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/sirupsen/logrus"
)

func parseSDKConfig(configFilePath string) (*conf.Config, error) {
	configs, err := conf.ParseConfigFile(configFilePath)
	if err != nil {
		return nil, err
	}
	config := &configs[0]
	return config, nil
}

//CheckConfigFile check if file empty
func CheckConfigFile(configFile string) {
	if configFile == "" {
		logrus.Fatal("Configuration Failed!")
	}
}

// CheckChainConnected check the connection
func CheckChainConnected(conn *client.Client) error {
	_, err := conn.GetBlockNumber(context.Background())
	return err
}

// FiscoClient client with config desc
type FiscoClient struct {
	Client *client.Client
	Config *conf.Config
}

// GroupFiscoClient clients in group
type GroupFiscoClient struct {
	Clients []*FiscoClient
}

// BuildClients build Group clients
func BuildClients(filePaths string) (map[int]*GroupFiscoClient, error) {
	sdkClients := map[int]*GroupFiscoClient{}
	configFiles := strings.Split(filePaths, ",")

	for _, v := range configFiles {
		c, err := BuildSDKClient(v)
		if err != nil {
			return nil, err
		}

		err = CheckChainConnected(c.Client)
		if err != nil {
			return nil, err
		}

		if groupClient, ok := sdkClients[c.Config.GroupID]; ok {
			groupClient.Clients = append(groupClient.Clients, c)
		} else {
			clients := []*FiscoClient{c}
			sdkClients[c.Config.GroupID] = &GroupFiscoClient{Clients: clients}
		}
	}
	return sdkClients, nil
}

// BuildSDKClient build fiscoClient
func BuildSDKClient(configFilePath string) (*FiscoClient, error) {
	config, err := parseSDKConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	c, err := client.Dial(config)
	if err != nil {
		return nil, err
	}

	return &FiscoClient{
		Client: c,
		Config: config,
	}, nil
}
