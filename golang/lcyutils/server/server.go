package server

import (
	"code.corp.bcollie.net/whistle/lcyutils/restapi"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/qq"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/shebei"
	"github.com/go-openapi/loads"
)

type Server struct {
	host     string
	endPoint int
	baseURL  string
}

func New() (*Server, error) {
	conf := &Server{
		host:     "0.0.0.0",
		endPoint: 8999,
		//baseURL:  "http://106.12.127.2:9190/",
		baseURL: "http://127.0.0.1:9189/",
	}
	return conf, nil
}

func (s *Server) ListenAndServer() error {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}
	api := operations.NewLcyutilsAPI(swaggerSpec)

	api.QqPostSendPrivateMsgHandler = qq.PostSendPrivateMsgHandlerFunc(s.handlePrivateMsg)

	api.ShebeiGetGetDeviceHandler = shebei.GetGetDeviceHandlerFunc(s.handleGetDevice)
	api.ShebeiPostModifyDeviceHandler = shebei.PostModifyDeviceHandlerFunc(s.handleModifyDevice)
	api.ShebeiPostModifyDeviceListHandler = shebei.PostModifyDeviceListHandlerFunc(s.handleModifyDeviceList)
	api.ShebeiGetGetDeviceListHandler = shebei.GetGetDeviceListHandlerFunc(s.handleGetDeviceList)
	api.ShebeiGetGetDeviceListFullHandler = shebei.GetGetDeviceListFullHandlerFunc(s.handleGetDeviceListFull)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = s.host
	server.Port = s.endPoint

	server.ConfigureAPI()
	return server.Serve()
}
