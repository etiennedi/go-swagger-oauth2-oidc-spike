// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/davecgh/go-spew/spew"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/etiennedi/go-swagger-oauth2-oidc-spike/restapi/operations"
	"github.com/etiennedi/go-swagger-oauth2-oidc-spike/restapi/operations/foo"
)

//go:generate swagger generate server --target ../../go-swagger-oauth2-oidc-spike --name Spikeyfoo --spec ../swagger.yml

func configureFlags(api *operations.SpikeyfooAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func setupOIDC() (*oidc.IDTokenVerifier, error) {
	provider, err := oidc.NewProvider(context.Background(), "http://localhost:8090/auth/realms/master")
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: "spikeyfoo",
	})

	return verifier, nil
}

func configureAPI(api *operations.SpikeyfooAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	verifier, err := setupOIDC()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.OidcAuth = func(rawToken string, scopes []string) (interface{}, error) {
		token, err := verifier.Verify(context.Background(), rawToken)
		if err != nil {
			return nil, fmt.Errorf("validation failed: %s", err)
		}

		spew.Dump(token)
		return token, nil
	}

	api.FooGetFooHandler = foo.GetFooHandlerFunc(func(params foo.GetFooParams, principal interface{}) middleware.Responder {
		spew.Dump(principal)
		return foo.NewGetFooOK().WithPayload("Hello world")
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
