package titaniumutils

import (
	"code.corp.bcollie.net/whistle/pkg/titanium-sols/tablesize"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var tableSizeSession map[int]*tablesize.TableSizeSession

func buildClient() (*client.Client, error) {
	configs, err := conf.ParseConfigFile("./cakeyconfig/config0.toml")
	if err != nil {
		return nil, err
	}
	config := &configs[0]
	if err != nil {
		return nil, err
	}
	c, err := client.Dial(config)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// SessionInit init session
func SessionInit(filePaths string) error {
	tableSizeSession = make(map[int]*tablesize.TableSizeSession)
	sdkClients, err := BuildClients(filePaths)
	if err != nil {
		return err
	}

	for groupID, groupClients := range sdkClients {
		session, err := MakeTableSizeSession(groupClients.Clients[0].Client)
		if err != nil {
			return err
		}
		tableSizeSession[groupID] = session
	}

	return nil
}

// MakeTableSizeSession make session
func MakeTableSizeSession(client *client.Client) (*tablesize.TableSizeSession, error) {
	contractAddress := common.HexToAddress(ContractTableSize)
	instance, err := tablesize.NewTableSize(contractAddress, client)
	if err != nil {
		return nil, err
	}

	tableSizeSession := &tablesize.TableSizeSession{
		Contract:     instance,
		CallOpts:     *client.GetCallOpts(),
		TransactOpts: *client.GetTransactOpts(),
	}
	return tableSizeSession, nil
}

// GetTableSizeByNameKey get a table size
func GetTableSizeByNameKey(tablename string, tablekey string) int64 {
	defaultGroupSesson, ok := tableSizeSession[int(ValidGroupID)]
	if !ok {
		return 0
	}
	value, err := defaultGroupSesson.Size(tablename, tablekey) // call Get API
	if err != nil {
		log.Info(err)
	}
	valueint64, err := BigIntToInt64(value)
	if err != nil {
		log.Info(err)
	}
	return valueint64
}
