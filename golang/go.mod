module code.corp.bcollie.net/whistle

go 1.16

replace github.com/FISCO-BCOS/go-sdk v0.11.0 => ./pkg/fisco-go-sdk

require (
	github.com/FISCO-BCOS/go-sdk v0.11.0
	github.com/ethereum/go-ethereum v1.10.15
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/loads v0.21.1
	github.com/go-openapi/runtime v0.23.0
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/go-openapi/validate v0.20.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/context v1.1.1
	github.com/influxdata/toml v0.0.0-20180607005434-2a2e3012f7cf
	github.com/jessevdk/go-flags v1.5.0
	github.com/sirupsen/logrus v1.8.1
	github.com/xuri/excelize/v2 v2.6.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3
	xorm.io/xorm v1.2.5
)
