// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/qq"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/shebei"
)

//go:generate swagger generate server --target ..\..\lcyutils --name Lcyutils --spec ..\swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.LcyutilsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LcyutilsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ShebeiGetGetDeviceHandler == nil {
		api.ShebeiGetGetDeviceHandler = shebei.GetGetDeviceHandlerFunc(func(params shebei.GetGetDeviceParams) middleware.Responder {
			return middleware.NotImplemented("operation shebei.GetGetDevice has not yet been implemented")
		})
	}
	if api.ShebeiGetGetDeviceListHandler == nil {
		api.ShebeiGetGetDeviceListHandler = shebei.GetGetDeviceListHandlerFunc(func(params shebei.GetGetDeviceListParams) middleware.Responder {
			return middleware.NotImplemented("operation shebei.GetGetDeviceList has not yet been implemented")
		})
	}
	if api.ShebeiPostModifyDeviceHandler == nil {
		api.ShebeiPostModifyDeviceHandler = shebei.PostModifyDeviceHandlerFunc(func(params shebei.PostModifyDeviceParams) middleware.Responder {
			return middleware.NotImplemented("operation shebei.PostModifyDevice has not yet been implemented")
		})
	}
	if api.ShebeiPostModifyDeviceListHandler == nil {
		api.ShebeiPostModifyDeviceListHandler = shebei.PostModifyDeviceListHandlerFunc(func(params shebei.PostModifyDeviceListParams) middleware.Responder {
			return middleware.NotImplemented("operation shebei.PostModifyDeviceList has not yet been implemented")
		})
	}
	if api.QqPostSendPrivateMsgHandler == nil {
		api.QqPostSendPrivateMsgHandler = qq.PostSendPrivateMsgHandlerFunc(func(params qq.PostSendPrivateMsgParams) middleware.Responder {
			return middleware.NotImplemented("operation qq.PostSendPrivateMsg has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
