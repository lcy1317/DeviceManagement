package main

import (
	"code.corp.bcollie.net/whistle/lcyutils/server"
	"fmt"
	log "github.com/sirupsen/logrus"
)

//swagger generate server -f ./swagger.yaml --exclude-main
func main() {
	//server.DBConnect()
	//server.QueryDevice(1)
	fmt.Println("Message Backend Start")
	testConf, _ := server.New()
	log.Info(testConf.ListenAndServer())
}
